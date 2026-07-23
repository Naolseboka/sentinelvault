package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// UserCreatedPayload is the webhook body for user.created.
type UserCreatedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewUserCreatedPayload(orgID, actorID, resourceID string, meta map[string]string) (*UserCreatedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &UserCreatedPayload{
		Event: "user.created", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *UserCreatedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *UserCreatedPayload) Validate() error {
	if p.Event != "user.created" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
