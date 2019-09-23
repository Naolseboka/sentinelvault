package audit

import "time"

// LegacyEvent is a simple audit DTO used by gateway event publishers.
// Prefer Chain Event for immutable signed logs.
type LegacyEvent struct {
	ID         string            `json:"id"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	Action     string            `json:"action"`
	Resource   string            `json:"resource"`
	Outcome    string            `json:"outcome"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	IPAddress  string            `json:"ip_address,omitempty"`
	OccurredAt time.Time         `json:"occurred_at"`
}
