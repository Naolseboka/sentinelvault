package apimodels

import "time"

// ApprovalModel is the API model for approval-service.
type ApprovalModel struct {
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

// CreateApprovalRequest is the Create request for approval-service.
type CreateApprovalRequest struct {
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

// CreateApprovalResponse is the Create response for approval-service.
type CreateApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// GetApprovalRequest is the Get request for approval-service.
type GetApprovalRequest struct {
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

// GetApprovalResponse is the Get response for approval-service.
type GetApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ListApprovalRequest is the List request for approval-service.
type ListApprovalRequest struct {
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

// ListApprovalResponse is the List response for approval-service.
type ListApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// UpdateApprovalRequest is the Update request for approval-service.
type UpdateApprovalRequest struct {
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

// UpdateApprovalResponse is the Update response for approval-service.
type UpdateApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// DeleteApprovalRequest is the Delete request for approval-service.
type DeleteApprovalRequest struct {
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

// DeleteApprovalResponse is the Delete response for approval-service.
type DeleteApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ArchiveApprovalRequest is the Archive request for approval-service.
type ArchiveApprovalRequest struct {
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

// ArchiveApprovalResponse is the Archive response for approval-service.
type ArchiveApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// RestoreApprovalRequest is the Restore request for approval-service.
type RestoreApprovalRequest struct {
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

// RestoreApprovalResponse is the Restore response for approval-service.
type RestoreApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ExportApprovalRequest is the Export request for approval-service.
type ExportApprovalRequest struct {
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

// ExportApprovalResponse is the Export response for approval-service.
type ExportApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ImportApprovalRequest is the Import request for approval-service.
type ImportApprovalRequest struct {
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

// ImportApprovalResponse is the Import response for approval-service.
type ImportApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ShareApprovalRequest is the Share request for approval-service.
type ShareApprovalRequest struct {
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

// ShareApprovalResponse is the Share response for approval-service.
type ShareApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// RotateApprovalRequest is the Rotate request for approval-service.
type RotateApprovalRequest struct {
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

// RotateApprovalResponse is the Rotate response for approval-service.
type RotateApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ApproveApprovalRequest is the Approve request for approval-service.
type ApproveApprovalRequest struct {
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

// ApproveApprovalResponse is the Approve response for approval-service.
type ApproveApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// DenyApprovalRequest is the Deny request for approval-service.
type DenyApprovalRequest struct {
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

// DenyApprovalResponse is the Deny response for approval-service.
type DenyApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// SearchApprovalRequest is the Search request for approval-service.
type SearchApprovalRequest struct {
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

// SearchApprovalResponse is the Search response for approval-service.
type SearchApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// StatsApprovalRequest is the Stats request for approval-service.
type StatsApprovalRequest struct {
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

// StatsApprovalResponse is the Stats response for approval-service.
type StatsApprovalResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	Item      *ApprovalModel  `json:"item,omitempty"`
	Items     []ApprovalModel `json:"items,omitempty"`
	Total     int64           `json:"total,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
