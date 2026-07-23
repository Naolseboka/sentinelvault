package domain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/events"
	"github.com/sentinelvault/sentinelvault/internal/queue"
)

const (
	defaultAccessTTL  = 15 * time.Minute
	defaultRefreshTTL = 7 * 24 * time.Hour
	issuerName        = "SentinelVault"
)

var slugRe = regexp.MustCompile(`[^a-z0-9]+`)

// RegisterInput bootstraps a user and organization.
type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"org_name"`
}

// LoginInput authenticates a user.
type LoginInput struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	DeviceFingerprint string `json:"device_fingerprint"`
	TOTPCode          string `json:"totp_code,omitempty"`
}

// LoginResult holds tokens or MFA challenge state.
type LoginResult struct {
	Tokens      *auth.TokenPair `json:"tokens,omitempty"`
	MFARequired bool            `json:"mfa_required"`
	UserID      string          `json:"user_id,omitempty"`
	OrgID       string          `json:"org_id,omitempty"`
	SessionID   string          `json:"session_id,omitempty"`
}

// UserRecord is the bootstrapped user aggregate stored locally.
type UserRecord struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Status      string    `json:"status"`
	MFAEnabled  bool      `json:"mfa_enabled"`
	CreatedAt   time.Time `json:"created_at"`
}

// OrgRecord is the bootstrapped organization aggregate.
type OrgRecord struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Credential stores authentication material for a user.
type Credential struct {
	UserID       string    `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	TOTPSecret   string    `json:"-"`
	MFAEnabled   bool      `json:"mfa_enabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SessionRef tracks an active authenticated session.
type SessionRef struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	OrgID       string     `json:"org_id"`
	DeviceID    string     `json:"device_id"`
	RefreshHash string     `json:"-"`
	ExpiresAt   time.Time  `json:"expires_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	MFAVerified bool       `json:"mfa_verified"`
	CreatedAt   time.Time  `json:"created_at"`
}

// CredentialRepository persists credentials.
type CredentialRepository interface {
	Create(ctx context.Context, cred *Credential) error
	GetByEmail(ctx context.Context, email string) (*Credential, error)
	GetByUserID(ctx context.Context, userID string) (*Credential, error)
	Update(ctx context.Context, cred *Credential) error
}

// SessionRepository persists session references.
type SessionRepository interface {
	Create(ctx context.Context, sess *SessionRef) error
	Get(ctx context.Context, id string) (*SessionRef, error)
	GetByRefreshHash(ctx context.Context, hash string) (*SessionRef, error)
	Update(ctx context.Context, sess *SessionRef) error
	ListByUser(ctx context.Context, userID string) ([]SessionRef, error)
}

// Service implements auth-service business logic.
type Service struct {
	creds     CredentialRepository
	sessions  SessionRepository
	users     *UserStore
	orgs      *OrgStore
	tokens    *auth.TokenManager
	events    queue.Producer
	jwtSecret []byte
}

// ServiceConfig configures the auth domain service.
type ServiceConfig struct {
	JWTSecret  []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Events     queue.Producer
}

// NewService creates an auth domain service.
func NewService(creds CredentialRepository, sessions SessionRepository, cfg ServiceConfig) *Service {
	accessTTL := cfg.AccessTTL
	if accessTTL <= 0 {
		accessTTL = defaultAccessTTL
	}
	secret := cfg.JWTSecret
	if len(secret) == 0 {
		secret = []byte("dev-jwt-secret-change-in-production")
	}
	return &Service{
		creds:     creds,
		sessions:  sessions,
		users:     NewUserStore(),
		orgs:      NewOrgStore(),
		tokens:    auth.NewTokenManager(secret, accessTTL),
		events:    cfg.Events,
		jwtSecret: secret,
	}
}

// Register creates a user, organization, and credential.
func (s *Service) Register(ctx context.Context, in RegisterInput) (*UserRecord, *OrgRecord, error) {
	if err := common.ValidateEmail(in.Email); err != nil {
		return nil, nil, err
	}
	if err := common.ValidateRequired("password", in.Password); err != nil {
		return nil, nil, err
	}
	if err := common.ValidateRequired("org_name", in.OrgName); err != nil {
		return nil, nil, err
	}
	email := strings.ToLower(strings.TrimSpace(in.Email))
	if existing, _ := s.creds.GetByEmail(ctx, email); existing != nil {
		return nil, nil, common.New(common.ErrCodeConflict, "email already registered")
	}

	now := common.NowUTC()
	orgID := common.NewID()
	org := &OrgRecord{
		ID:        orgID,
		Name:      strings.TrimSpace(in.OrgName),
		Slug:      slugify(in.OrgName),
		Status:    "active",
		CreatedAt: now,
	}
	userID := common.NewID()
	user := &UserRecord{
		ID:          userID,
		OrgID:       orgID,
		Email:       email,
		DisplayName: strings.Split(email, "@")[0],
		Status:      "active",
		CreatedAt:   now,
	}
	hash, err := auth.HashPassword(in.Password)
	if err != nil {
		return nil, nil, err
	}
	cred := &Credential{
		UserID:       userID,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.orgs.Create(org); err != nil {
		return nil, nil, err
	}
	if err := s.users.Create(user); err != nil {
		return nil, nil, err
	}
	if err := s.creds.Create(ctx, cred); err != nil {
		return nil, nil, err
	}
	s.publishEvent(ctx, auth.NewAuthEvent(auth.EventLoginSucceeded).WithUser(userID).WithOrg(orgID).WithEmail(email))
	return user, org, nil
}

// Login authenticates credentials and returns tokens or MFA requirement.
func (s *Service) Login(ctx context.Context, in LoginInput) (*LoginResult, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))
	cred, err := s.creds.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if cred == nil {
		s.publishEvent(ctx, auth.NewAuthEvent(auth.EventLoginFailed).WithEmail(email).WithReason("unknown email"))
		return nil, common.New(common.ErrCodeUnauthorized, "invalid credentials")
	}
	if err := auth.VerifyPassword(in.Password, cred.PasswordHash); err != nil {
		s.publishEvent(ctx, auth.NewAuthEvent(auth.EventLoginFailed).WithUser(cred.UserID).WithEmail(email).WithReason("bad password"))
		return nil, common.New(common.ErrCodeUnauthorized, "invalid credentials")
	}
	user, err := s.users.Get(cred.UserID)
	if err != nil || user == nil {
		return nil, common.New(common.ErrCodeUnauthorized, "invalid credentials")
	}
	if user.Status != "active" {
		return nil, common.New(common.ErrCodeForbidden, "user disabled")
	}

	if cred.MFAEnabled {
		if strings.TrimSpace(in.TOTPCode) == "" {
			return &LoginResult{MFARequired: true, UserID: user.ID, OrgID: user.OrgID}, nil
		}
		if !auth.ValidateTOTPCode(cred.TOTPSecret, in.TOTPCode) {
			s.publishEvent(ctx, auth.NewAuthEvent(auth.EventLoginFailed).WithUser(user.ID).WithReason("invalid totp"))
			return nil, common.New(common.ErrCodeUnauthorized, "invalid mfa code")
		}
	}

	tokens, sess, err := s.issueSession(ctx, user, cred.MFAEnabled, in.DeviceFingerprint, cred.MFAEnabled)
	if err != nil {
		return nil, err
	}
	s.publishEvent(ctx, auth.NewAuthEvent(auth.EventLoginSucceeded).WithUser(user.ID).WithOrg(user.OrgID).WithSession(sess.ID).WithEmail(email))
	return &LoginResult{
		Tokens:      tokens,
		MFARequired: false,
		UserID:      user.ID,
		OrgID:       user.OrgID,
		SessionID:   sess.ID,
	}, nil
}

// Refresh rotates refresh token and issues a new access token.
func (s *Service) Refresh(ctx context.Context, refreshToken string) (*auth.TokenPair, error) {
	if strings.TrimSpace(refreshToken) == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "refresh_token is required")
	}
	hash := hashRefreshToken(refreshToken)
	sess, err := s.sessions.GetByRefreshHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if sess == nil || sess.RevokedAt != nil || time.Now().UTC().After(sess.ExpiresAt) {
		return nil, common.New(common.ErrCodeUnauthorized, "invalid refresh token")
	}
	user, err := s.users.Get(sess.UserID)
	if err != nil || user == nil {
		return nil, common.New(common.ErrCodeUnauthorized, "invalid refresh token")
	}

	newRefresh, err := auth.NewRefreshTokenID()
	if err != nil {
		return nil, err
	}
	sess.RefreshHash = hashRefreshToken(newRefresh)
	sess.ExpiresAt = time.Now().UTC().Add(defaultRefreshTTL)
	if err := s.sessions.Update(ctx, sess); err != nil {
		return nil, err
	}

	access, err := s.tokens.IssueAccessToken(auth.Claims{
		Subject:     auth.UserID(user.ID),
		OrgID:       user.OrgID,
		SessionID:   sess.ID,
		Roles:       []string{"admin"},
		MFAVerified: sess.MFAVerified,
	})
	if err != nil {
		return nil, err
	}
	return &auth.TokenPair{
		AccessToken:  access,
		RefreshToken: newRefresh,
		TokenType:    "Bearer",
		ExpiresIn:    int(defaultAccessTTL.Seconds()),
		ExpiresAt:    time.Now().UTC().Add(defaultAccessTTL),
	}, nil
}

// Logout revokes a session.
func (s *Service) Logout(ctx context.Context, sessionID string) error {
	sess, err := s.sessions.Get(ctx, sessionID)
	if err != nil {
		return err
	}
	if sess == nil {
		return common.New(common.ErrCodeNotFound, "session not found")
	}
	now := common.NowUTC()
	sess.RevokedAt = &now
	if err := s.sessions.Update(ctx, sess); err != nil {
		return err
	}
	s.publishEvent(ctx, auth.NewAuthEvent(auth.EventSessionRevoked).WithUser(sess.UserID).WithOrg(sess.OrgID).WithSession(sess.ID))
	return nil
}

// EnrollTOTP starts MFA enrollment for a user.
func (s *Service) EnrollTOTP(ctx context.Context, userID string) (*auth.TOTPSecret, error) {
	cred, err := s.creds.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if cred == nil {
		return nil, common.New(common.ErrCodeNotFound, "credential not found")
	}
	secret, err := auth.GenerateTOTPSecret(issuerName, cred.Email)
	if err != nil {
		return nil, err
	}
	cred.TOTPSecret = secret.Secret
	cred.UpdatedAt = common.NowUTC()
	if err := s.creds.Update(ctx, cred); err != nil {
		return nil, err
	}
	return secret, nil
}

// VerifyTOTP confirms enrollment and enables MFA.
func (s *Service) VerifyTOTP(ctx context.Context, userID, code string) error {
	cred, err := s.creds.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if cred == nil || cred.TOTPSecret == "" {
		return common.New(common.ErrCodeInvalidArgument, "totp enrollment not started")
	}
	if !auth.ValidateTOTPCode(cred.TOTPSecret, code) {
		return common.New(common.ErrCodeUnauthorized, "invalid totp code")
	}
	cred.MFAEnabled = true
	cred.UpdatedAt = common.NowUTC()
	if err := s.creds.Update(ctx, cred); err != nil {
		return err
	}
	if err := s.users.SetMFAEnabled(userID, true); err != nil {
		return err
	}
	s.publishEvent(ctx, auth.NewAuthEvent(auth.EventMFAEnrolled).WithUser(userID))
	return nil
}

func (s *Service) issueSession(ctx context.Context, user *UserRecord, mfaVerified bool, deviceFingerprint string, mfaEnabled bool) (*auth.TokenPair, *SessionRef, error) {
	refreshRaw, err := auth.NewRefreshTokenID()
	if err != nil {
		return nil, nil, err
	}
	now := common.NowUTC()
	sess := &SessionRef{
		ID:          common.NewID(),
		UserID:      user.ID,
		OrgID:       user.OrgID,
		DeviceID:    deviceFingerprint,
		RefreshHash: hashRefreshToken(refreshRaw),
		ExpiresAt:   now.Add(defaultRefreshTTL),
		MFAVerified: mfaVerified || !mfaEnabled,
		CreatedAt:   now,
	}
	if err := s.sessions.Create(ctx, sess); err != nil {
		return nil, nil, err
	}
	access, err := s.tokens.IssueAccessToken(auth.Claims{
		Subject:     auth.UserID(user.ID),
		OrgID:       user.OrgID,
		SessionID:   sess.ID,
		Roles:       []string{"admin"},
		MFAVerified: sess.MFAVerified,
	})
	if err != nil {
		return nil, nil, err
	}
	return &auth.TokenPair{
		AccessToken:  access,
		RefreshToken: refreshRaw,
		TokenType:    "Bearer",
		ExpiresIn:    int(defaultAccessTTL.Seconds()),
		ExpiresAt:    now.Add(defaultAccessTTL),
		MFARequired:  mfaEnabled && !mfaVerified,
	}, sess, nil
}

func (s *Service) publishEvent(ctx context.Context, ev auth.AuthEvent) {
	if s.events == nil {
		return
	}
	payload, _ := json.Marshal(ev)
	_ = s.events.Publish(ctx, events.TopicUserCreated, ev.UserID, payload)
}

func hashRefreshToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func slugify(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = slugRe.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		s = "org"
	}
	return s
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "auth" }

// TableName returns the primary table name.
func TableName() string { return "credentials" }

// UserStore holds bootstrapped users in memory.
type UserStore struct {
	mu      sync.RWMutex
	byID    map[string]*UserRecord
	byEmail map[string]string
}

func NewUserStore() *UserStore {
	return &UserStore{byID: map[string]*UserRecord{}, byEmail: map[string]string{}}
}

func (s *UserStore) Create(u *UserRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := *u
	s.byID[u.ID] = &cp
	s.byEmail[u.Email] = u.ID
	return nil
}

func (s *UserStore) Get(id string) (*UserRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.byID[id]
	if !ok {
		return nil, nil
	}
	cp := *u
	return &cp, nil
}

func (s *UserStore) SetMFAEnabled(id string, enabled bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.byID[id]
	if !ok {
		return common.New(common.ErrCodeNotFound, "user not found")
	}
	u.MFAEnabled = enabled
	return nil
}

// OrgStore holds bootstrapped orgs in memory.
type OrgStore struct {
	mu   sync.RWMutex
	byID map[string]*OrgRecord
}

func NewOrgStore() *OrgStore {
	return &OrgStore{byID: map[string]*OrgRecord{}}
}

func (s *OrgStore) Create(o *OrgRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := *o
	s.byID[o.ID] = &cp
	return nil
}

// MemoryCredentialRepository is an in-memory credential store.
type MemoryCredentialRepository struct {
	mu      sync.RWMutex
	byEmail map[string]*Credential
	byUser  map[string]*Credential
}

func NewMemoryCredentialRepository() *MemoryCredentialRepository {
	return &MemoryCredentialRepository{
		byEmail: map[string]*Credential{},
		byUser:  map[string]*Credential{},
	}
}

func (r *MemoryCredentialRepository) Create(_ context.Context, cred *Credential) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *cred
	r.byEmail[cred.Email] = &cp
	r.byUser[cred.UserID] = &cp
	return nil
}

func (r *MemoryCredentialRepository) GetByEmail(_ context.Context, email string) (*Credential, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byEmail[email]
	if !ok {
		return nil, nil
	}
	cp := *c
	return &cp, nil
}

func (r *MemoryCredentialRepository) GetByUserID(_ context.Context, userID string) (*Credential, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byUser[userID]
	if !ok {
		return nil, nil
	}
	cp := *c
	return &cp, nil
}

func (r *MemoryCredentialRepository) Update(_ context.Context, cred *Credential) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *cred
	r.byEmail[cred.Email] = &cp
	r.byUser[cred.UserID] = &cp
	return nil
}

// MemorySessionRepository is an in-memory session store.
type MemorySessionRepository struct {
	mu            sync.RWMutex
	byID          map[string]*SessionRef
	byRefreshHash map[string]string
}

func NewMemorySessionRepository() *MemorySessionRepository {
	return &MemorySessionRepository{
		byID:          map[string]*SessionRef{},
		byRefreshHash: map[string]string{},
	}
}

func (r *MemorySessionRepository) Create(_ context.Context, sess *SessionRef) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *sess
	r.byID[sess.ID] = &cp
	r.byRefreshHash[sess.RefreshHash] = sess.ID
	return nil
}

func (r *MemorySessionRepository) Get(_ context.Context, id string) (*SessionRef, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.byID[id]
	if !ok {
		return nil, nil
	}
	cp := *s
	return &cp, nil
}

func (r *MemorySessionRepository) GetByRefreshHash(_ context.Context, hash string) (*SessionRef, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byRefreshHash[hash]
	if !ok {
		return nil, nil
	}
	s := r.byID[id]
	cp := *s
	return &cp, nil
}

func (r *MemorySessionRepository) Update(_ context.Context, sess *SessionRef) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for h, id := range r.byRefreshHash {
		if id == sess.ID && h != sess.RefreshHash {
			delete(r.byRefreshHash, h)
		}
	}
	cp := *sess
	r.byID[sess.ID] = &cp
	r.byRefreshHash[sess.RefreshHash] = sess.ID
	return nil
}

func (r *MemorySessionRepository) ListByUser(_ context.Context, userID string) ([]SessionRef, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]SessionRef, 0)
	for _, s := range r.byID {
		if s.UserID == userID {
			out = append(out, *s)
		}
	}
	return out, nil
}
