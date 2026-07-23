package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PolicyUpdatedPayload is the webhook body for policy.updated.
type PolicyUpdatedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewPolicyUpdatedPayload(orgID, actorID, resourceID string, meta map[string]string) (*PolicyUpdatedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &PolicyUpdatedPayload{
		Event: "policy.updated", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *PolicyUpdatedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *PolicyUpdatedPayload) Validate() error {
	if p.Event != "policy.updated" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
