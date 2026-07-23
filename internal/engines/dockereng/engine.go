package dockereng

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Role struct {
	ID         string            `json:"id"`
	OrgID      string            `json:"org_id"`
	Name       string            `json:"name"`
	Config     map[string]string `json:"config"`
	DefaultTTL time.Duration     `json:"default_ttl"`
	MaxTTL     time.Duration     `json:"max_ttl"`
}

type Lease struct {
	ID        string            `json:"id"`
	RoleID    string            `json:"role_id"`
	OrgID     string            `json:"org_id"`
	Token     string            `json:"token"`
	ExpiresAt time.Time         `json:"expires_at"`
	Revoked   bool              `json:"revoked"`
	Meta      map[string]string `json:"meta,omitempty"`
}

type Engine struct {
	policies map[string]*PolicyDocument
	mu       sync.Mutex
	roles    map[string]*Role
	leases   map[string]*Lease
}

func NewEngine() *Engine {
	return &Engine{policies: map[string]*PolicyDocument{}, roles: map[string]*Role{}, leases: map[string]*Lease{}}
}

func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
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
	if r.Config == nil {
		r.Config = map[string]string{}
	}
	cp := r
	e.mu.Lock()
	e.roles[r.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

func (e *Engine) Issue(roleID string, ttl time.Duration, meta map[string]string) (*Lease, error) {
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
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%d", role.ID, role.Name, time.Now().UnixNano())))
	l := &Lease{
		ID: common.NewID(), RoleID: role.ID, OrgID: role.OrgID,
		Token:     "dockereng_" + hex.EncodeToString(sum[:16]),
		ExpiresAt: common.NowUTC().Add(ttl), Meta: meta,
	}
	e.leases[l.ID] = l
	cp := *l
	return &cp, nil
}

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

func (e *Engine) ListLeases(orgID string) []Lease {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]Lease, 0)
	for _, l := range e.leases {
		if orgID == "" || l.OrgID == orgID {
			out = append(out, *l)
		}
	}
	return out
}

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
