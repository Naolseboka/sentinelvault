package kubernetes

import (
	"encoding/base64"
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	svcrypto "github.com/sentinelvault/sentinelvault/internal/crypto"
)

// Role defines a Kubernetes service account token role.
type Role struct {
	ID         string        `json:"id"`
	OrgID      string        `json:"org_id"`
	Name       string        `json:"name"`
	Namespace  string        `json:"namespace"`
	SAName     string        `json:"service_account"`
	Audiences  []string      `json:"audiences"`
	DefaultTTL time.Duration `json:"default_ttl"`
}

// ServiceAccountToken is a projected token.
type ServiceAccountToken struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Token     string    `json:"token"`
	Namespace string    `json:"namespace"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Engine issues Kubernetes SA tokens (simulated JWT).
type Engine struct {
	mu       sync.Mutex
	policies map[string]*PolicyDocument
	roles    map[string]*Role
	tokens   map[string]*ServiceAccountToken
	signer   []byte
}

// NewEngine creates a Kubernetes secrets engine.
func NewEngine() (*Engine, error) {
	signer, err := svcrypto.RandomBytes(32)
	if err != nil {
		return nil, err
	}
	return &Engine{
		policies: map[string]*PolicyDocument{},
		roles:    map[string]*Role{},
		tokens:   map[string]*ServiceAccountToken{},
		signer:   signer,
	}, nil
}

// UpsertRole stores a role.
func (e *Engine) UpsertRole(r Role) (*Role, error) {
	if r.OrgID == "" || r.Name == "" || r.Namespace == "" || r.SAName == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, namespace, service_account required")
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

// IssueToken creates a service account token.
func (e *Engine) IssueToken(roleID string, ttl time.Duration) (*ServiceAccountToken, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	role, ok := e.roles[roleID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	if ttl <= 0 {
		ttl = role.DefaultTTL
	}
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	claims, _ := json.Marshal(map[string]any{
		"iss": "sentinelvault/k8s",
		"sub": "system:serviceaccount:" + role.Namespace + ":" + role.SAName,
		"aud": role.Audiences,
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"kubernetes.io": map[string]any{
			"namespace":      role.Namespace,
			"serviceaccount": map[string]string{"name": role.SAName},
		},
	})
	payload := base64.RawURLEncoding.EncodeToString(claims)
	mac, err := svcrypto.HMACSHA256(e.signer, []byte(header+"."+payload))
	if err != nil {
		// fallback: hash
		sum := svcrypto.MustSHA256(append([]byte(header+"."+payload), e.signer...))
		mac = sum
	}
	sig := base64.RawURLEncoding.EncodeToString(mac)
	tok := &ServiceAccountToken{
		ID: common.NewID(), RoleID: role.ID, Namespace: role.Namespace,
		Token: header + "." + payload + "." + sig, ExpiresAt: common.NowUTC().Add(ttl),
	}
	e.tokens[tok.ID] = tok
	cp := *tok
	return &cp, nil
}
