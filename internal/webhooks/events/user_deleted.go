package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// UserDeletedPayload is the webhook body for user.deleted.
type UserDeletedPayload struct {
	Event     string            `json:"event"`
	OrgID     string            `json:"org_id"`
	ActorID   string            `json:"actor_id"`
	ResourceID string           `json:"resource_id"`
	OccurredAt time.Time        `json:"occurred_at"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func NewUserDeletedPayload(orgID, actorID, resourceID string, meta map[string]string) (*UserDeletedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &UserDeletedPayload{
		Event: "user.deleted", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *UserDeletedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *UserDeletedPayload) Validate() error {
	if p.Event != "user.deleted" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
