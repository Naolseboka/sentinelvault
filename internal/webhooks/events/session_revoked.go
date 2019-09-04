package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SessionRevokedPayload is the webhook body for session.revoked.
type SessionRevokedPayload struct {
	Event     string            `json:"event"`
	OrgID     string            `json:"org_id"`
	ActorID   string            `json:"actor_id"`
	ResourceID string           `json:"resource_id"`
	OccurredAt time.Time        `json:"occurred_at"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func NewSessionRevokedPayload(orgID, actorID, resourceID string, meta map[string]string) (*SessionRevokedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &SessionRevokedPayload{
		Event: "session.revoked", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *SessionRevokedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *SessionRevokedPayload) Validate() error {
	if p.Event != "session.revoked" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
