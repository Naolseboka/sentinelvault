package apimodels

import "time"

// SearchDocModel is the API model for search-service.
type SearchDocModel struct {
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

// CreateSearchDocRequest is the Create request for search-service.
type CreateSearchDocRequest struct {
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

// CreateSearchDocResponse is the Create response for search-service.
type CreateSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// GetSearchDocRequest is the Get request for search-service.
type GetSearchDocRequest struct {
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

// GetSearchDocResponse is the Get response for search-service.
type GetSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ListSearchDocRequest is the List request for search-service.
type ListSearchDocRequest struct {
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

// ListSearchDocResponse is the List response for search-service.
type ListSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// UpdateSearchDocRequest is the Update request for search-service.
type UpdateSearchDocRequest struct {
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

// UpdateSearchDocResponse is the Update response for search-service.
type UpdateSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DeleteSearchDocRequest is the Delete request for search-service.
type DeleteSearchDocRequest struct {
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

// DeleteSearchDocResponse is the Delete response for search-service.
type DeleteSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ArchiveSearchDocRequest is the Archive request for search-service.
type ArchiveSearchDocRequest struct {
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

// ArchiveSearchDocResponse is the Archive response for search-service.
type ArchiveSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RestoreSearchDocRequest is the Restore request for search-service.
type RestoreSearchDocRequest struct {
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

// RestoreSearchDocResponse is the Restore response for search-service.
type RestoreSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ExportSearchDocRequest is the Export request for search-service.
type ExportSearchDocRequest struct {
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

// ExportSearchDocResponse is the Export response for search-service.
type ExportSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ImportSearchDocRequest is the Import request for search-service.
type ImportSearchDocRequest struct {
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

// ImportSearchDocResponse is the Import response for search-service.
type ImportSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ShareSearchDocRequest is the Share request for search-service.
type ShareSearchDocRequest struct {
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

// ShareSearchDocResponse is the Share response for search-service.
type ShareSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RotateSearchDocRequest is the Rotate request for search-service.
type RotateSearchDocRequest struct {
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

// RotateSearchDocResponse is the Rotate response for search-service.
type RotateSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ApproveSearchDocRequest is the Approve request for search-service.
type ApproveSearchDocRequest struct {
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

// ApproveSearchDocResponse is the Approve response for search-service.
type ApproveSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DenySearchDocRequest is the Deny request for search-service.
type DenySearchDocRequest struct {
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

// DenySearchDocResponse is the Deny response for search-service.
type DenySearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// SearchSearchDocRequest is the Search request for search-service.
type SearchSearchDocRequest struct {
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

// SearchSearchDocResponse is the Search response for search-service.
type SearchSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// StatsSearchDocRequest is the Stats request for search-service.
type StatsSearchDocRequest struct {
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

// StatsSearchDocResponse is the Stats response for search-service.
type StatsSearchDocResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SearchDocModel `json:"item,omitempty"`
	Items []SearchDocModel `json:"items,omitempty"`
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
