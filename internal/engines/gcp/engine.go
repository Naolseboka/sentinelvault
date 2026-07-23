package gcp

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role maps to a GCP service account impersonation role.
type Role struct {
	ID             string        `json:"id"`
	OrgID          string        `json:"org_id"`
	Name           string        `json:"name"`
	ProjectID      string        `json:"project_id"`
	ServiceAccount string        `json:"service_account"`
	Scopes         []string      `json:"scopes"`
	DefaultTTL     time.Duration `json:"default_ttl"`
}

// AccessToken is a temporary GCP access token.
type AccessToken struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Engine issues GCP-style access tokens.
type Engine struct {
	mu       sync.Mutex
	policies map[string]*PolicyDocument
	roles    map[string]*Role
	tokens   map[string]*AccessToken
}

// NewEngine creates a GCP secrets engine.
func NewEngine() *Engine {
	return &Engine{
		policies: map[string]*PolicyDocument{},
		roles:    map[string]*Role{},
		tokens:   map[string]*AccessToken{},
	}
}

// UpsertRole stores a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" || r.ProjectID == "" || r.ServiceAccount == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, project_id, service_account required")
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

// IssueToken creates an access token for a role.
func (e *Engine) IssueToken(roleID string, ttl time.Duration) (*AccessToken, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	role, ok := e.roles[roleID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	if ttl <= 0 {
		ttl = role.DefaultTTL
	}
	claims := map[string]any{
		"iss": role.ServiceAccount,
		"aud": "https://oauth2.googleapis.com/token",
		"sub": role.ServiceAccount,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(ttl).Unix(),
		"scp": role.Scopes,
	}
	raw, _ := json.Marshal(claims)
	sum := sha256.Sum256(append(raw, []byte(role.ProjectID)...))
	tok := &AccessToken{
		ID: common.NewID(), RoleID: role.ID,
		Token:     "ya29." + base64.RawURLEncoding.EncodeToString(sum[:]),
		ExpiresAt: common.NowUTC().Add(ttl),
	}
	e.tokens[tok.ID] = tok
	cp := *tok
	return &cp, nil
}
// maintenance note 39
