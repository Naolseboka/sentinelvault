package apimodels

import "time"

// LifecycleModel is the API model for lifecycle-service.
type LifecycleModel struct {
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

// CreateLifecycleRequest is the Create request for lifecycle-service.
type CreateLifecycleRequest struct {
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

// CreateLifecycleResponse is the Create response for lifecycle-service.
type CreateLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// GetLifecycleRequest is the Get request for lifecycle-service.
type GetLifecycleRequest struct {
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

// GetLifecycleResponse is the Get response for lifecycle-service.
type GetLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ListLifecycleRequest is the List request for lifecycle-service.
type ListLifecycleRequest struct {
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

// ListLifecycleResponse is the List response for lifecycle-service.
type ListLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// UpdateLifecycleRequest is the Update request for lifecycle-service.
type UpdateLifecycleRequest struct {
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

// UpdateLifecycleResponse is the Update response for lifecycle-service.
type UpdateLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DeleteLifecycleRequest is the Delete request for lifecycle-service.
type DeleteLifecycleRequest struct {
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

// DeleteLifecycleResponse is the Delete response for lifecycle-service.
type DeleteLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ArchiveLifecycleRequest is the Archive request for lifecycle-service.
type ArchiveLifecycleRequest struct {
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

// ArchiveLifecycleResponse is the Archive response for lifecycle-service.
type ArchiveLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RestoreLifecycleRequest is the Restore request for lifecycle-service.
type RestoreLifecycleRequest struct {
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

// RestoreLifecycleResponse is the Restore response for lifecycle-service.
type RestoreLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ExportLifecycleRequest is the Export request for lifecycle-service.
type ExportLifecycleRequest struct {
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

// ExportLifecycleResponse is the Export response for lifecycle-service.
type ExportLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ImportLifecycleRequest is the Import request for lifecycle-service.
type ImportLifecycleRequest struct {
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

// ImportLifecycleResponse is the Import response for lifecycle-service.
type ImportLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ShareLifecycleRequest is the Share request for lifecycle-service.
type ShareLifecycleRequest struct {
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

// ShareLifecycleResponse is the Share response for lifecycle-service.
type ShareLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RotateLifecycleRequest is the Rotate request for lifecycle-service.
type RotateLifecycleRequest struct {
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

// RotateLifecycleResponse is the Rotate response for lifecycle-service.
type RotateLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ApproveLifecycleRequest is the Approve request for lifecycle-service.
type ApproveLifecycleRequest struct {
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

// ApproveLifecycleResponse is the Approve response for lifecycle-service.
type ApproveLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DenyLifecycleRequest is the Deny request for lifecycle-service.
type DenyLifecycleRequest struct {
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

// DenyLifecycleResponse is the Deny response for lifecycle-service.
type DenyLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// SearchLifecycleRequest is the Search request for lifecycle-service.
type SearchLifecycleRequest struct {
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

// SearchLifecycleResponse is the Search response for lifecycle-service.
type SearchLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// StatsLifecycleRequest is the Stats request for lifecycle-service.
type StatsLifecycleRequest struct {
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

// StatsLifecycleResponse is the Stats response for lifecycle-service.
type StatsLifecycleResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *LifecycleModel  `json:"item,omitempty"`
	Items     []LifecycleModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
