package aws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role maps to IAM-like temporary credentials.
type Role struct {
	ID          string        `json:"id"`
	OrgID       string        `json:"org_id"`
	Name        string        `json:"name"`
	AccountID   string        `json:"account_id"`
	PolicyARNs  []string      `json:"policy_arns"`
	DefaultTTL  time.Duration `json:"default_ttl"`
	MaxTTL      time.Duration `json:"max_ttl"`
}

// Credentials are STS-like temporary keys.
type Credentials struct {
	ID              string    `json:"id"`
	RoleID          string    `json:"role_id"`
	AccessKeyID     string    `json:"access_key_id"`
	SecretAccessKey string    `json:"secret_access_key"`
	SessionToken    string    `json:"session_token"`
	ExpiresAt       time.Time `json:"expires_at"`
}

// Engine issues temporary AWS-style credentials (software simulation).
type Engine struct {
	mu       sync.Mutex
	policies map[string]*PolicyDocument
	roles    map[string]*Role
	creds    map[string]*Credentials
	root     []byte
}

// NewEngine creates an AWS secrets engine.
func NewEngine(rootSecret []byte) *Engine {
	if len(rootSecret) < 32 {
		rootSecret = sha256.New().Sum([]byte("sentinelvault-aws-engine"))
	}
	return &Engine{
		policies: map[string]*PolicyDocument{},
		roles:    map[string]*Role{},
		creds:    map[string]*Credentials{},
		root:     rootSecret,
	}
}

// UpsertRole stores a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" || r.AccountID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, account_id required")
	}
	if r.ID == "" {
		r.ID = common.NewID()
	}
	if r.DefaultTTL <= 0 {
		r.DefaultTTL = time.Hour
	}
	if r.MaxTTL <= 0 {
		r.MaxTTL = 12 * time.Hour
	}
	cp := r
	e.mu.Lock()
	e.roles[r.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

// AssumeRole issues temporary credentials.
func (e *Engine) AssumeRole(roleID string, ttl time.Duration) (*Credentials, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	role, ok := e.roles[roleID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	if ttl <= 0 {
		ttl = role.DefaultTTL
	}
	if ttl > role.MaxTTL {
		ttl = role.MaxTTL
	}
	id := common.NewID()
	akid := "AKIA" + hex.EncodeToString(hmacSHA(e.root, []byte(id)))[:16]
	secret := hex.EncodeToString(hmacSHA(e.root, []byte(akid+role.AccountID)))
	token := hex.EncodeToString(hmacSHA(e.root, []byte(fmt.Sprintf("%s:%s:%d", id, role.ID, time.Now().UnixNano()))))
	c := &Credentials{
		ID: id, RoleID: role.ID, AccessKeyID: akid, SecretAccessKey: secret,
		SessionToken: token, ExpiresAt: common.NowUTC().Add(ttl),
	}
	e.creds[c.ID] = c
	cp := *c
	return &cp, nil
}

func hmacSHA(key, msg []byte) []byte {
	m := hmac.New(sha256.New, key)
	_, _ = m.Write(msg)
	return m.Sum(nil)
}
