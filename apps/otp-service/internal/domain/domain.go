package domain

import (
	"context"
	"encoding/base64"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
	"github.com/sentinelvault/sentinelvault/internal/common"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
)

// TOTPRecord stores an encrypted TOTP seed for a user.
type TOTPRecord struct {
	ID         string    `json:"id"`
	OrgID      string    `json:"org_id"`
	UserID     string    `json:"user_id"`
	Label      string    `json:"label"`
	SecretBlob []byte    `json:"-"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
}

// TOTPRecordView is the safe API representation without secret material.
type TOTPRecordView struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	UserID    string    `json:"user_id"`
	Label     string    `json:"label"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
}

// EnrollInput holds TOTP enrollment fields.
type EnrollInput struct {
	OrgID  string `json:"org_id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	KeyID  string `json:"key_id"`
}

// EnrollResult returns provisioning details (secret shown once at enroll).
type EnrollResult struct {
	Record *TOTPRecordView `json:"record"`
	Secret *auth.TOTPSecret `json:"secret"`
}

// VerifyInput holds TOTP verification fields.
type VerifyInput struct {
	UserID string `json:"user_id"`
	Code   string `json:"code"`
	KeyID  string `json:"key_id"`
	AAD    string `json:"aad"`
}

// Service manages encrypted TOTP seeds.
type Service struct {
	mu         sync.RWMutex
	records    map[string]*TOTPRecord // by userID
	encryptor  *encdomain.InProcessEncryptor
	defaultKey string
}

// NewService creates an OTP domain service.
func NewService(enc *encdomain.Service, defaultKeyID string) *Service {
	return &Service{
		records:    map[string]*TOTPRecord{},
		encryptor:  &encdomain.InProcessEncryptor{Svc: enc},
		defaultKey: defaultKeyID,
	}
}

// Enroll generates and stores an encrypted TOTP seed.
func (s *Service) Enroll(ctx context.Context, in EnrollInput) (*EnrollResult, error) {
	if in.UserID == "" || in.OrgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and user_id required")
	}
	keyID := in.KeyID
	if keyID == "" {
		keyID = s.defaultKey
	}
	secret, err := auth.GenerateTOTPSecret("SentinelVault", auth.TOTPAccountFromEmail(in.Email))
	if err != nil {
		return nil, err
	}
	enc, err := s.encryptor.EnvelopeEncrypt(ctx, encdomain.EncryptRequest{
		KeyID: keyID, Plaintext: []byte(secret.Secret), AAD: []byte(in.OrgID),
	})
	if err != nil {
		return nil, err
	}
	rec := &TOTPRecord{
		ID: common.NewID(), OrgID: in.OrgID, UserID: in.UserID,
		Label: in.Email, SecretBlob: enc.Ciphertext, Enabled: false,
		CreatedAt: common.NowUTC(),
	}
	s.mu.Lock()
	s.records[in.UserID] = rec
	s.mu.Unlock()
	return &EnrollResult{Record: toRecordView(rec), Secret: secret}, nil
}

// Verify validates a TOTP code and enables the factor on success.
func (s *Service) Verify(ctx context.Context, in VerifyInput) error {
	s.mu.RLock()
	rec, ok := s.records[in.UserID]
	s.mu.RUnlock()
	if !ok {
		return common.New(common.ErrCodeNotFound, "totp not enrolled")
	}
	aad := in.AAD
	if aad == "" {
		aad = rec.OrgID
	}
	dec, err := s.encryptor.EnvelopeDecrypt(ctx, encdomain.DecryptRequest{
		Ciphertext: rec.SecretBlob, AAD: []byte(aad),
	})
	if err != nil {
		return err
	}
	if !auth.ValidateTOTPCode(string(dec.Plaintext), in.Code) {
		return common.New(common.ErrCodeInvalidArgument, "invalid totp code")
	}
	s.mu.Lock()
	rec.Enabled = true
	s.mu.Unlock()
	return nil
}

// Get returns enrollment metadata without the encrypted secret.
func (s *Service) Get(_ context.Context, userID string) (*TOTPRecordView, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rec, ok := s.records[userID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "totp not enrolled")
	}
	return toRecordView(rec), nil
}

func toRecordView(rec *TOTPRecord) *TOTPRecordView {
	return &TOTPRecordView{
		ID: rec.ID, OrgID: rec.OrgID, UserID: rec.UserID,
		Label: rec.Label, Enabled: rec.Enabled, CreatedAt: rec.CreatedAt,
	}
}

// SecretBlobBase64 returns encrypted blob encoding for persistence layers.
func SecretBlobBase64(blob []byte) string {
	return base64.StdEncoding.EncodeToString(blob)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "otp" }

// TableName returns the primary table name.
func TableName() string { return "totp_seeds" }
