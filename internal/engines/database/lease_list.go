package database

import "time"

// LeaseSummary is a list view of active leases.
type LeaseSummary struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `json:"revoked"`
}

func (e *Engine) ListLeases(orgID string, includeRevoked bool) []LeaseSummary {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]LeaseSummary, 0)
	for _, l := range e.leases {
		role := e.roles[l.RoleID]
		if orgID != "" && (role == nil || role.OrgID != orgID) {
			continue
		}
		if !includeRevoked && l.Revoked {
			continue
		}
		out = append(out, LeaseSummary{ID: l.ID, RoleID: l.RoleID, ExpiresAt: l.ExpiresAt, Revoked: l.Revoked})
	}
	return out
}

func (e *Engine) ExpireLeases(now time.Time) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	n := 0
	for _, l := range e.leases {
		if !l.Revoked && !now.Before(l.ExpiresAt) {
			l.Revoked = true
			n++
		}
	}
	return n
}
