package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// OrgProvisionedPayload is the webhook body for org.provisioned.
type OrgProvisionedPayload struct {
	Event     string            `json:"event"`
	OrgID     string            `json:"org_id"`
	ActorID   string            `json:"actor_id"`
	ResourceID string           `json:"resource_id"`
	OccurredAt time.Time        `json:"occurred_at"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func NewOrgProvisionedPayload(orgID, actorID, resourceID string, meta map[string]string) (*OrgProvisionedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &OrgProvisionedPayload{
		Event: "org.provisioned", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *OrgProvisionedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *OrgProvisionedPayload) Validate() error {
	if p.Event != "org.provisioned" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
// maintenance note 53
