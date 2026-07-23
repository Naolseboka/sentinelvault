package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// EmergencyGrantedPayload is the webhook body for emergency.granted.
type EmergencyGrantedPayload struct {
	Event      string            `json:"event"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewEmergencyGrantedPayload(orgID, actorID, resourceID string, meta map[string]string) (*EmergencyGrantedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &EmergencyGrantedPayload{
		Event: "emergency.granted", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *EmergencyGrantedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *EmergencyGrantedPayload) Validate() error {
	if p.Event != "emergency.granted" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
