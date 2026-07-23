package apimodels

import "time"

// EncOpModel is the API model for encryption-service.
type EncOpModel struct {
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

// CreateEncOpRequest is the Create request for encryption-service.
type CreateEncOpRequest struct {
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

// CreateEncOpResponse is the Create response for encryption-service.
type CreateEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// GetEncOpRequest is the Get request for encryption-service.
type GetEncOpRequest struct {
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

// GetEncOpResponse is the Get response for encryption-service.
type GetEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ListEncOpRequest is the List request for encryption-service.
type ListEncOpRequest struct {
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

// ListEncOpResponse is the List response for encryption-service.
type ListEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// UpdateEncOpRequest is the Update request for encryption-service.
type UpdateEncOpRequest struct {
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

// UpdateEncOpResponse is the Update response for encryption-service.
type UpdateEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DeleteEncOpRequest is the Delete request for encryption-service.
type DeleteEncOpRequest struct {
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

// DeleteEncOpResponse is the Delete response for encryption-service.
type DeleteEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ArchiveEncOpRequest is the Archive request for encryption-service.
type ArchiveEncOpRequest struct {
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

// ArchiveEncOpResponse is the Archive response for encryption-service.
type ArchiveEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RestoreEncOpRequest is the Restore request for encryption-service.
type RestoreEncOpRequest struct {
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

// RestoreEncOpResponse is the Restore response for encryption-service.
type RestoreEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ExportEncOpRequest is the Export request for encryption-service.
type ExportEncOpRequest struct {
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

// ExportEncOpResponse is the Export response for encryption-service.
type ExportEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ImportEncOpRequest is the Import request for encryption-service.
type ImportEncOpRequest struct {
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

// ImportEncOpResponse is the Import response for encryption-service.
type ImportEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ShareEncOpRequest is the Share request for encryption-service.
type ShareEncOpRequest struct {
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

// ShareEncOpResponse is the Share response for encryption-service.
type ShareEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RotateEncOpRequest is the Rotate request for encryption-service.
type RotateEncOpRequest struct {
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

// RotateEncOpResponse is the Rotate response for encryption-service.
type RotateEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ApproveEncOpRequest is the Approve request for encryption-service.
type ApproveEncOpRequest struct {
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

// ApproveEncOpResponse is the Approve response for encryption-service.
type ApproveEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DenyEncOpRequest is the Deny request for encryption-service.
type DenyEncOpRequest struct {
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

// DenyEncOpResponse is the Deny response for encryption-service.
type DenyEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// SearchEncOpRequest is the Search request for encryption-service.
type SearchEncOpRequest struct {
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

// SearchEncOpResponse is the Search response for encryption-service.
type SearchEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// StatsEncOpRequest is the Stats request for encryption-service.
type StatsEncOpRequest struct {
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

// StatsEncOpResponse is the Stats response for encryption-service.
type StatsEncOpResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *EncOpModel  `json:"item,omitempty"`
	Items     []EncOpModel `json:"items,omitempty"`
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
