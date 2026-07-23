package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// FindingDetectedPayload is the webhook body for finding.detected.
type FindingDetectedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewFindingDetectedPayload(orgID, actorID, resourceID string, meta map[string]string) (*FindingDetectedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &FindingDetectedPayload{
		Event: "finding.detected", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *FindingDetectedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *FindingDetectedPayload) Validate() error {
	if p.Event != "finding.detected" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
