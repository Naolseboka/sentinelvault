package azure

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role represents an Azure service principal role binding.
type Role struct {
	ID           string        `json:"id"`
	OrgID        string        `json:"org_id"`
	Name         string        `json:"name"`
	TenantID     string        `json:"tenant_id"`
	Subscription string        `json:"subscription"`
	Scopes       []string      `json:"scopes"`
	DefaultTTL   time.Duration `json:"default_ttl"`
}

// Credentials are temporary Azure-like credentials.
type Credentials struct {
	ID           string    `json:"id"`
	RoleID       string    `json:"role_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	TenantID     string    `json:"tenant_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// Engine issues Azure-style temporary credentials.
type Engine struct {
	mu       sync.Mutex
	policies map[string]*PolicyDocument
	roles    map[string]*Role
	creds    map[string]*Credentials
}

// NewEngine creates an Azure secrets engine.
func NewEngine() *Engine {
	return &Engine{
		policies: map[string]*PolicyDocument{},
		roles:    map[string]*Role{},
		creds:    map[string]*Credentials{},
	}
}

// UpsertRole stores a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" || r.TenantID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, tenant_id required")
	}
	if r.ID == "" {
		r.ID = common.NewID()
	}
	if r.DefaultTTL <= 0 {
		r.DefaultTTL = 2 * time.Hour
	}
	cp := r
	e.mu.Lock()
	e.roles[r.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

// Issue creates temporary credentials for a role.
func (e *Engine) Issue(roleID string, ttl time.Duration) (*Credentials, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	role, ok := e.roles[roleID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	if ttl <= 0 {
		ttl = role.DefaultTTL
	}
	id := common.NewID()
	sum := sha256.Sum256([]byte(id + role.TenantID + role.Name))
	c := &Credentials{
		ID: id, RoleID: role.ID, TenantID: role.TenantID,
		ClientID:     hex.EncodeToString(sum[:8]),
		ClientSecret: hex.EncodeToString(sum[:]),
		ExpiresAt:    common.NowUTC().Add(ttl),
	}
	e.creds[c.ID] = c
	cp := *c
	return &cp, nil
}
