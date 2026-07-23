package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// ItemRotatedPayload is the webhook body for item.rotated.
type ItemRotatedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewItemRotatedPayload(orgID, actorID, resourceID string, meta map[string]string) (*ItemRotatedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &ItemRotatedPayload{
		Event: "item.rotated", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *ItemRotatedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ItemRotatedPayload) Validate() error {
	if p.Event != "item.rotated" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
