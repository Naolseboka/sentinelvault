package apimodels

import "time"

// AuditModel is the API model for audit-service.
type AuditModel struct {
	ID          string            `json:"id"`
	OrgID       string            `json:"org_id"`
	Name        string            `json:"name"`
	Status      string            `json:"status"`
	Description string            `json:"description,omitempty"`
	OwnerID     string            `json:"owner_id,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Version     int               `json:"version"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   *time.Time        `json:"deleted_at,omitempty"`
}

// CreateAuditRequest is the Create request for audit-service.
type CreateAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// CreateAuditResponse is the Create response for audit-service.
type CreateAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// GetAuditRequest is the Get request for audit-service.
type GetAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// GetAuditResponse is the Get response for audit-service.
type GetAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ListAuditRequest is the List request for audit-service.
type ListAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ListAuditResponse is the List response for audit-service.
type ListAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// UpdateAuditRequest is the Update request for audit-service.
type UpdateAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// UpdateAuditResponse is the Update response for audit-service.
type UpdateAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DeleteAuditRequest is the Delete request for audit-service.
type DeleteAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// DeleteAuditResponse is the Delete response for audit-service.
type DeleteAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ArchiveAuditRequest is the Archive request for audit-service.
type ArchiveAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ArchiveAuditResponse is the Archive response for audit-service.
type ArchiveAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RestoreAuditRequest is the Restore request for audit-service.
type RestoreAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// RestoreAuditResponse is the Restore response for audit-service.
type RestoreAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ExportAuditRequest is the Export request for audit-service.
type ExportAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ExportAuditResponse is the Export response for audit-service.
type ExportAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ImportAuditRequest is the Import request for audit-service.
type ImportAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ImportAuditResponse is the Import response for audit-service.
type ImportAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ShareAuditRequest is the Share request for audit-service.
type ShareAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ShareAuditResponse is the Share response for audit-service.
type ShareAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RotateAuditRequest is the Rotate request for audit-service.
type RotateAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// RotateAuditResponse is the Rotate response for audit-service.
type RotateAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ApproveAuditRequest is the Approve request for audit-service.
type ApproveAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// ApproveAuditResponse is the Approve response for audit-service.
type ApproveAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DenyAuditRequest is the Deny request for audit-service.
type DenyAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// DenyAuditResponse is the Deny response for audit-service.
type DenyAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// SearchAuditRequest is the Search request for audit-service.
type SearchAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// SearchAuditResponse is the Search response for audit-service.
type SearchAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// StatsAuditRequest is the Stats request for audit-service.
type StatsAuditRequest struct {
	OrgID          string            `json:"org_id"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Status         string            `json:"status,omitempty"`
	ActorID        string            `json:"actor_id,omitempty"`
	ResourceID     string            `json:"resource_id,omitempty"`
	Reason         string            `json:"reason,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	Payload        map[string]any    `json:"payload,omitempty"`
	Offset         int               `json:"offset,omitempty"`
	Limit          int               `json:"limit,omitempty"`
	IncludeDeleted bool              `json:"include_deleted,omitempty"`
}

// StatsAuditResponse is the Stats response for audit-service.
type StatsAuditResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *AuditModel  `json:"item,omitempty"`
	Items     []AuditModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
