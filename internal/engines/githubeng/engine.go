package githubeng

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role for GitHub App installation tokens.
type Role struct {
	ID             string            `json:"id"`
	OrgID          string            `json:"org_id"`
	Name           string            `json:"name"`
	InstallationID string            `json:"installation_id"`
	Repositories   []string          `json:"repositories"`
	Permissions    map[string]string `json:"permissions"`
	DefaultTTL     time.Duration     `json:"default_ttl"`
}

// Token is a GitHub installation access token.
type Token struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Engine issues GitHub installation tokens (simulated).
type Engine struct {
	policies map[string]*PolicyDocument
	mu       sync.Mutex
	roles    map[string]*Role
	tokens   map[string]*Token
}

// NewEngine creates a GitHub secrets engine.
func NewEngine() *Engine {
	return &Engine{policies: map[string]*PolicyDocument{}, roles: map[string]*Role{}, tokens: map[string]*Token{}}
}

// UpsertRole stores a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" || r.InstallationID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, installation_id required")
	}
	if r.ID == "" {
		r.ID = common.NewID()
	}
	if r.DefaultTTL <= 0 {
		r.DefaultTTL = time.Hour
	}
	cp := r
	e.mu.Lock()
	e.roles[r.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

// IssueToken creates an installation token.
func (e *Engine) IssueToken(roleID string, ttl time.Duration) (*Token, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	role, ok := e.roles[roleID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	if ttl <= 0 {
		ttl = role.DefaultTTL
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%d", role.InstallationID, role.ID, time.Now().UnixNano())))
	tok := &Token{
		ID: common.NewID(), RoleID: role.ID,
		Token: "ghs_" + hex.EncodeToString(sum[:20]), ExpiresAt: common.NowUTC().Add(ttl),
	}
	e.tokens[tok.ID] = tok
	cp := *tok
	return &cp, nil
}
