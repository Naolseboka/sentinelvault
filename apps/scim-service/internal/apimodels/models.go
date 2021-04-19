package apimodels

import "time"

// SCIMModel is the API model for scim-service.
type SCIMModel struct {
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

// CreateSCIMRequest is the Create request for scim-service.
type CreateSCIMRequest struct {
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

// CreateSCIMResponse is the Create response for scim-service.
type CreateSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// GetSCIMRequest is the Get request for scim-service.
type GetSCIMRequest struct {
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

// GetSCIMResponse is the Get response for scim-service.
type GetSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ListSCIMRequest is the List request for scim-service.
type ListSCIMRequest struct {
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

// ListSCIMResponse is the List response for scim-service.
type ListSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// UpdateSCIMRequest is the Update request for scim-service.
type UpdateSCIMRequest struct {
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

// UpdateSCIMResponse is the Update response for scim-service.
type UpdateSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DeleteSCIMRequest is the Delete request for scim-service.
type DeleteSCIMRequest struct {
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

// DeleteSCIMResponse is the Delete response for scim-service.
type DeleteSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ArchiveSCIMRequest is the Archive request for scim-service.
type ArchiveSCIMRequest struct {
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

// ArchiveSCIMResponse is the Archive response for scim-service.
type ArchiveSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RestoreSCIMRequest is the Restore request for scim-service.
type RestoreSCIMRequest struct {
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

// RestoreSCIMResponse is the Restore response for scim-service.
type RestoreSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ExportSCIMRequest is the Export request for scim-service.
type ExportSCIMRequest struct {
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

// ExportSCIMResponse is the Export response for scim-service.
type ExportSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ImportSCIMRequest is the Import request for scim-service.
type ImportSCIMRequest struct {
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

// ImportSCIMResponse is the Import response for scim-service.
type ImportSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ShareSCIMRequest is the Share request for scim-service.
type ShareSCIMRequest struct {
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

// ShareSCIMResponse is the Share response for scim-service.
type ShareSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RotateSCIMRequest is the Rotate request for scim-service.
type RotateSCIMRequest struct {
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

// RotateSCIMResponse is the Rotate response for scim-service.
type RotateSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ApproveSCIMRequest is the Approve request for scim-service.
type ApproveSCIMRequest struct {
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

// ApproveSCIMResponse is the Approve response for scim-service.
type ApproveSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DenySCIMRequest is the Deny request for scim-service.
type DenySCIMRequest struct {
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

// DenySCIMResponse is the Deny response for scim-service.
type DenySCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// SearchSCIMRequest is the Search request for scim-service.
type SearchSCIMRequest struct {
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

// SearchSCIMResponse is the Search response for scim-service.
type SearchSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// StatsSCIMRequest is the Stats request for scim-service.
type StatsSCIMRequest struct {
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

// StatsSCIMResponse is the Stats response for scim-service.
type StatsSCIMResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *SCIMModel `json:"item,omitempty"`
	Items []SCIMModel `json:"items,omitempty"`
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
