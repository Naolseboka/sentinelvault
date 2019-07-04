package database

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role defines a dynamic DB credential role.
type Role struct {
	ID               string        `json:"id"`
	OrgID            string        `json:"org_id"`
	Name             string        `json:"name"`
	DBType           string        `json:"db_type"` // postgres|mysql|mssql
	DefaultTTL       time.Duration `json:"default_ttl"`
	MaxTTL           time.Duration `json:"max_ttl"`
	CreationSQL      string        `json:"creation_sql"`
	RevocationSQL    string        `json:"revocation_sql"`
	UsernameTemplate string        `json:"username_template"`
}

// Lease is a short-lived credential.
type Lease struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	OrgID     string    `json:"org_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `json:"revoked"`
}

// Engine issues dynamic database credentials.
type Engine struct {
	mu     sync.Mutex
	roles  map[string]*Role
	leases map[string]*Lease
}

// NewEngine creates a database secrets engine.
func NewEngine() *Engine {
	return &Engine{roles: map[string]*Role{}, leases: map[string]*Lease{}}
}

// UpsertRole registers a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
	}
	switch r.DBType {
	case "postgres", "mysql", "mssql":
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, "unsupported db_type")
	}
	if r.ID == "" {
		r.ID = common.NewID()
	}
	if r.DefaultTTL <= 0 {
		r.DefaultTTL = time.Hour
	}
	if r.MaxTTL <= 0 {
		r.MaxTTL = 24 * time.Hour
	}
	if r.UsernameTemplate == "" {
		r.UsernameTemplate = "sv-{{.Role}}-{{.Random}}"
	}
	cp := r
	e.mu.Lock()
	e.roles[r.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

// Issue creates a lease for a role.
func (e *Engine) Issue(roleID string, ttl time.Duration) (*Lease, error) {
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
	pw, err := randomPassword(24)
	if err != nil {
		return nil, err
	}
	randPart, err := randomPassword(6)
	if err != nil {
		return nil, err
	}
	username := fmt.Sprintf("sv-%s-%s", sanitize(role.Name), randPart)
	lease := &Lease{
		ID: common.NewID(), RoleID: role.ID, OrgID: role.OrgID,
		Username: username, Password: pw, ExpiresAt: common.NowUTC().Add(ttl),
	}
	e.leases[lease.ID] = lease
	cp := *lease
	return &cp, nil
}

// Revoke invalidates a lease.
func (e *Engine) Revoke(leaseID string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	l, ok := e.leases[leaseID]
	if !ok {
		return common.New(common.ErrCodeNotFound, "lease not found")
	}
	l.Revoked = true
	return nil
}

// CleanupExpired revokes expired leases; returns count.
func (e *Engine) CleanupExpired(now time.Time) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	n := 0
	for _, l := range e.leases {
		if !l.Revoked && now.After(l.ExpiresAt) {
			l.Revoked = true
			n++
		}
	}
	return n
}

func randomPassword(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b)[:n], nil
}

func sanitize(s string) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			out = append(out, r)
		}
	}
	if len(out) == 0 {
		return "role"
	}
	return string(out)
}
