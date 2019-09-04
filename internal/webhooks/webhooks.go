package webhooks

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

type Endpoint struct {
	ID        string
	OrgID     string
	URL       string
	Secret    string
	Events    []string
	Active    bool
	CreatedAt time.Time
}

type Attempt struct {
	ID         string
	EndpointID string
	EventType  string
	StatusCode int
	Success    bool
	Error      string
	Attempt    int
	CreatedAt  time.Time
}

type Registry struct {
	mu   sync.Mutex
	eps  map[string]*Endpoint
	logs []Attempt
	client *http.Client
}

func NewRegistry() *Registry {
	return &Registry{eps: map[string]*Endpoint{}, logs: []Attempt{}, client: &http.Client{Timeout: 5 * time.Second}}
}

func (r *Registry) Register(orgID, rawURL string, events []string) (*Endpoint, error) {
	if orgID == "" || rawURL == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and url required")
	}
	sec, err := crypto.RandomBytes(16)
	if err != nil {
		return nil, err
	}
	ep := &Endpoint{ID: common.NewID(), OrgID: orgID, URL: rawURL, Secret: hex.EncodeToString(sec), Events: append([]string(nil), events...), Active: true, CreatedAt: common.NowUTC()}
	r.mu.Lock()
	r.eps[ep.ID] = ep
	r.mu.Unlock()
	cp := *ep
	return &cp, nil
}

func Sign(secret string, body []byte) string {
	m := hmac.New(sha256.New, []byte(secret))
	_, _ = m.Write(body)
	return hex.EncodeToString(m.Sum(nil))
}

func Backoff(attempt int) time.Duration {
	if attempt <= 0 {
		return time.Second
	}
	d := time.Second << (attempt - 1)
	if d > 30*time.Second {
		return 30 * time.Second
	}
	return d
}

func (r *Registry) Deliver(ctx context.Context, orgID, eventType string, body []byte) []Attempt {
	r.mu.Lock()
	targets := make([]*Endpoint, 0)
	for _, ep := range r.eps {
		if ep.OrgID == orgID && ep.Active && contains(ep.Events, eventType) {
			cp := *ep
			targets = append(targets, &cp)
		}
	}
	r.mu.Unlock()
	out := make([]Attempt, 0)
	for _, ep := range targets {
		att := r.deliverOne(ctx, ep, eventType, body, 1)
		out = append(out, att)
	}
	return out
}

func (r *Registry) deliverOne(ctx context.Context, ep *Endpoint, eventType string, body []byte, attempt int) Attempt {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.URL, bytes.NewReader(body))
	att := Attempt{ID: common.NewID(), EndpointID: ep.ID, EventType: eventType, Attempt: attempt, CreatedAt: common.NowUTC()}
	if err != nil {
		att.Error = err.Error()
		r.log(att)
		return att
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-SV-Signature", Sign(ep.Secret, body))
	req.Header.Set("X-SV-Event", eventType)
	res, err := r.client.Do(req)
	if err != nil {
		att.Error = err.Error()
		r.log(att)
		return att
	}
	defer res.Body.Close()
	_, _ = io.Copy(io.Discard, res.Body)
	att.StatusCode = res.StatusCode
	att.Success = res.StatusCode >= 200 && res.StatusCode < 300
	if !att.Success {
		att.Error = res.Status
	}
	r.log(att)
	return att
}

func (r *Registry) log(a Attempt) {
	r.mu.Lock()
	r.logs = append(r.logs, a)
	r.mu.Unlock()
}

func (r *Registry) Attempts() []Attempt {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]Attempt, len(r.logs))
	copy(out, r.logs)
	return out
}

func contains(ss []string, x string) bool {
	for _, s := range ss {
		if s == x || s == "*" {
			return true
		}
	}
	return false
}
