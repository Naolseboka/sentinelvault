package apimodels

import "time"

// RotateModel is the API model for rotation-service.
type RotateModel struct {
	ID string `json:"id"`
	OrgID string `json:"org_id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Description string `json:"description,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Version int `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// CreateRotateRequest is the Create request for rotation-service.
type CreateRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// CreateRotateResponse is the Create response for rotation-service.
type CreateRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// GetRotateRequest is the Get request for rotation-service.
type GetRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// GetRotateResponse is the Get response for rotation-service.
type GetRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ListRotateRequest is the List request for rotation-service.
type ListRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ListRotateResponse is the List response for rotation-service.
type ListRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// UpdateRotateRequest is the Update request for rotation-service.
type UpdateRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// UpdateRotateResponse is the Update response for rotation-service.
type UpdateRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DeleteRotateRequest is the Delete request for rotation-service.
type DeleteRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// DeleteRotateResponse is the Delete response for rotation-service.
type DeleteRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ArchiveRotateRequest is the Archive request for rotation-service.
type ArchiveRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ArchiveRotateResponse is the Archive response for rotation-service.
type ArchiveRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RestoreRotateRequest is the Restore request for rotation-service.
type RestoreRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// RestoreRotateResponse is the Restore response for rotation-service.
type RestoreRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ExportRotateRequest is the Export request for rotation-service.
type ExportRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ExportRotateResponse is the Export response for rotation-service.
type ExportRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ImportRotateRequest is the Import request for rotation-service.
type ImportRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ImportRotateResponse is the Import response for rotation-service.
type ImportRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ShareRotateRequest is the Share request for rotation-service.
type ShareRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ShareRotateResponse is the Share response for rotation-service.
type ShareRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RotateRotateRequest is the Rotate request for rotation-service.
type RotateRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// RotateRotateResponse is the Rotate response for rotation-service.
type RotateRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ApproveRotateRequest is the Approve request for rotation-service.
type ApproveRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// ApproveRotateResponse is the Approve response for rotation-service.
type ApproveRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DenyRotateRequest is the Deny request for rotation-service.
type DenyRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// DenyRotateResponse is the Deny response for rotation-service.
type DenyRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// SearchRotateRequest is the Search request for rotation-service.
type SearchRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// SearchRotateResponse is the Search response for rotation-service.
type SearchRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// StatsRotateRequest is the Stats request for rotation-service.
type StatsRotateRequest struct {
	OrgID string `json:"org_id"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	ActorID string `json:"actor_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// StatsRotateResponse is the Stats response for rotation-service.
type StatsRotateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *RotateModel `json:"item,omitempty"`
	Items []RotateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}


// ErrorBody is a standard API error.
type ErrorBody struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
