package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// ItemSharedPayload is the webhook body for item.shared.
type ItemSharedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewItemSharedPayload(orgID, actorID, resourceID string, meta map[string]string) (*ItemSharedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &ItemSharedPayload{
		Event: "item.shared", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *ItemSharedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ItemSharedPayload) Validate() error {
	if p.Event != "item.shared" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
