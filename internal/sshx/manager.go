package sshx

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	svcrypto "github.com/sentinelvault/sentinelvault/internal/crypto"
	"golang.org/x/crypto/ssh"
)

// KeyType distinguishes user vs host keys.
const (
	KeyTypeUser = "user"
	KeyTypeHost = "host"
)

// KeyRecord stores SSH key material metadata.
type KeyRecord struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	PublicKey   string    `json:"public_key"`
	PrivatePEM  string    `json:"private_pem,omitempty"`
	Fingerprint string    `json:"fingerprint"`
	CreatedAt   time.Time `json:"created_at"`
	RotatedAt   *time.Time `json:"rotated_at,omitempty"`
}

// SignedCert is an SSH certificate issued by a CA key.
type SignedCert struct {
	ID         string    `json:"id"`
	KeyID      string    `json:"key_id"`
	OrgID      string    `json:"org_id"`
	Principals []string  `json:"principals"`
	ValidAfter time.Time `json:"valid_after"`
	ValidBefore time.Time `json:"valid_before"`
	CertString string    `json:"cert"`
}

// Manager handles SSH keys and certificates.
type Manager struct {
	mu      sync.Mutex
	keys    map[string]*KeyRecord
	certs   map[string]*SignedCert
	caPriv  ed25519.PrivateKey
	caPub   ed25519.PublicKey
	signer  ssh.Signer
}

// NewManager creates a manager with an embedded SSH CA.
func NewManager() (*Manager, error) {
	pub, priv, err := svcrypto.GenerateEd25519KeyPair()
	if err != nil {
		return nil, err
	}
	signer, err := ssh.NewSignerFromKey(priv)
	if err != nil {
		return nil, err
	}
	return &Manager{
		keys:   map[string]*KeyRecord{},
		certs:  map[string]*SignedCert{},
		caPriv: priv,
		caPub:  pub,
		signer: signer,
	}, nil
}

// GenerateKey creates a new Ed25519 SSH keypair.
func (m *Manager) GenerateKey(orgID, name, keyType string) (*KeyRecord, error) {
	if orgID == "" || name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
	}
	if keyType != KeyTypeUser && keyType != KeyTypeHost {
		return nil, common.New(common.ErrCodeInvalidArgument, "type must be user or host")
	}
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		return nil, err
	}
	block, err := ssh.MarshalPrivateKey(priv, "")
	if err != nil {
		return nil, err
	}
	rec := &KeyRecord{
		ID: common.NewID(), OrgID: orgID, Name: name, Type: keyType,
		PublicKey:   string(ssh.MarshalAuthorizedKey(sshPub)),
		PrivatePEM:  string(pem.EncodeToMemory(block)),
		Fingerprint: ssh.FingerprintSHA256(sshPub),
		CreatedAt:   common.NowUTC(),
	}
	m.mu.Lock()
	m.keys[rec.ID] = rec
	m.mu.Unlock()
	cp := *rec
	return &cp, nil
}

// RotateKey replaces key material keeping the same ID metadata name/type.
func (m *Manager) RotateKey(id string) (*KeyRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	old, ok := m.keys[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		return nil, err
	}
	block, err := ssh.MarshalPrivateKey(priv, "")
	if err != nil {
		return nil, err
	}
	now := common.NowUTC()
	old.PublicKey = string(ssh.MarshalAuthorizedKey(sshPub))
	old.PrivatePEM = string(pem.EncodeToMemory(block))
	old.Fingerprint = ssh.FingerprintSHA256(sshPub)
	old.RotatedAt = &now
	cp := *old
	return &cp, nil
}

// SignUserCert issues an SSH user certificate.
func (m *Manager) SignUserCert(orgID, keyID string, principals []string, ttl time.Duration) (*SignedCert, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	rec, ok := m.keys[keyID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	if rec.OrgID != orgID {
		return nil, common.New(common.ErrCodeForbidden, "org mismatch")
	}
	pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(rec.PublicKey))
	if err != nil {
		return nil, err
	}
	if ttl <= 0 {
		ttl = 8 * time.Hour
	}
	now := time.Now().UTC()
	cert := &ssh.Certificate{
		Key:             pk,
		CertType:        ssh.UserCert,
		KeyId:           fmt.Sprintf("%s/%s", orgID, keyID),
		ValidPrincipals: principals,
		ValidAfter:      uint64(now.Unix()),
		ValidBefore:     uint64(now.Add(ttl).Unix()),
		Permissions: ssh.Permissions{
			Extensions: map[string]string{"permit-pty": ""},
		},
	}
	if err := cert.SignCert(rand.Reader, m.signer); err != nil {
		return nil, err
	}
	out := &SignedCert{
		ID: common.NewID(), KeyID: keyID, OrgID: orgID, Principals: principals,
		ValidAfter: now, ValidBefore: now.Add(ttl),
		CertString: string(ssh.MarshalAuthorizedKey(cert)),
	}
	m.certs[out.ID] = out
	cp := *out
	return &cp, nil
}

// GetKey returns key without private material unless includePrivate.
func (m *Manager) GetKey(id string, includePrivate bool) (*KeyRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	rec, ok := m.keys[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	cp := *rec
	if !includePrivate {
		cp.PrivatePEM = ""
	}
	return &cp, nil
}
