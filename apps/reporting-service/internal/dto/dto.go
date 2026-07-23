package dto

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// ReportDTO is the transport object for reporting APIs.
type ReportDTO struct {
	ID             string            `json:"id"`
	OrgID          string            `json:"org_id"`
	Name           string            `json:"name"`
	Status         string            `json:"status"`
	ExternalID     string            `json:"externalid"`
	Description    string            `json:"description"`
	OwnerID        string            `json:"ownerid"`
	CollectionID   string            `json:"collectionid"`
	FolderID       string            `json:"folderid"`
	RiskLevel      string            `json:"risklevel"`
	Classification string            `json:"classification"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

// CreateReportRequest is the create payload.
type CreateReportRequest struct {
	OrgID       string            `json:"org_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	OwnerID     string            `json:"owner_id"`
	Tags        []string          `json:"tags"`
	Metadata    map[string]string `json:"metadata"`
	Payload     json.RawMessage   `json:"payload"`
}

// UpdateReportRequest is the update payload.
type UpdateReportRequest struct {
	Name        *string           `json:"name"`
	Description *string           `json:"description"`
	Status      *string           `json:"status"`
	Tags        []string          `json:"tags"`
	Metadata    map[string]string `json:"metadata"`
	Payload     json.RawMessage   `json:"payload"`
}

// ValidateCreate validates a create request.
func ValidateCreateReport(r CreateReportRequest) error {
	if strings.TrimSpace(r.OrgID) == "" {
		return common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return common.New(common.ErrCodeInvalidArgument, "name required")
	}
	if len(r.Name) > 256 {
		return common.New(common.ErrCodeInvalidArgument, "name too long")
	}
	for _, t := range r.Tags {
		if len(t) > 64 {
			return common.New(common.ErrCodeInvalidArgument, "tag too long")
		}
	}
	return nil
}

// ApplyUpdate applies patch fields onto dto.
func ApplyUpdateReport(dst *ReportDTO, r UpdateReportRequest) {
	if r.Name != nil {
		dst.Name = strings.TrimSpace(*r.Name)
	}
	if r.Description != nil {
		dst.Description = strings.TrimSpace(*r.Description)
	}
	if r.Status != nil {
		dst.Status = strings.TrimSpace(*r.Status)
	}
	if r.Tags != nil {
		dst.Tags = append([]string(nil), r.Tags...)
	}
	if r.Metadata != nil {
		if dst.Metadata == nil {
			dst.Metadata = map[string]string{}
		}
		for k, v := range r.Metadata {
			dst.Metadata[k] = v
		}
	}
	dst.UpdatedAt = common.NowUTC()
}
