package webhooks

import (
	"encoding/json"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// BackupCompletedPayload is the webhook body for backup.completed.
type BackupCompletedPayload struct {
	Event     string            `json:"event"`
	OrgID     string            `json:"org_id"`
	ActorID   string            `json:"actor_id"`
	ResourceID string           `json:"resource_id"`
	OccurredAt time.Time        `json:"occurred_at"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func NewBackupCompletedPayload(orgID, actorID, resourceID string, meta map[string]string) (*BackupCompletedPayload, error) {
	if orgID == "" || actorID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
	}
	return &BackupCompletedPayload{
		Event: "backup.completed", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		OccurredAt: common.NowUTC(), Metadata: meta,
	}, nil
}

func (p *BackupCompletedPayload) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *BackupCompletedPayload) Validate() error {
	if p.Event != "backup.completed" {
		return common.New(common.ErrCodeInvalidArgument, "event type mismatch")
	}
	return nil
}
