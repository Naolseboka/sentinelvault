package apimodels

import "time"

// DelegateModel is the API model for delegation-service.
type DelegateModel struct {
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

// CreateDelegateRequest is the Create request for delegation-service.
type CreateDelegateRequest struct {
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

// CreateDelegateResponse is the Create response for delegation-service.
type CreateDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// GetDelegateRequest is the Get request for delegation-service.
type GetDelegateRequest struct {
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

// GetDelegateResponse is the Get response for delegation-service.
type GetDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ListDelegateRequest is the List request for delegation-service.
type ListDelegateRequest struct {
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

// ListDelegateResponse is the List response for delegation-service.
type ListDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// UpdateDelegateRequest is the Update request for delegation-service.
type UpdateDelegateRequest struct {
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

// UpdateDelegateResponse is the Update response for delegation-service.
type UpdateDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DeleteDelegateRequest is the Delete request for delegation-service.
type DeleteDelegateRequest struct {
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

// DeleteDelegateResponse is the Delete response for delegation-service.
type DeleteDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ArchiveDelegateRequest is the Archive request for delegation-service.
type ArchiveDelegateRequest struct {
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

// ArchiveDelegateResponse is the Archive response for delegation-service.
type ArchiveDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RestoreDelegateRequest is the Restore request for delegation-service.
type RestoreDelegateRequest struct {
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

// RestoreDelegateResponse is the Restore response for delegation-service.
type RestoreDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ExportDelegateRequest is the Export request for delegation-service.
type ExportDelegateRequest struct {
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

// ExportDelegateResponse is the Export response for delegation-service.
type ExportDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ImportDelegateRequest is the Import request for delegation-service.
type ImportDelegateRequest struct {
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

// ImportDelegateResponse is the Import response for delegation-service.
type ImportDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ShareDelegateRequest is the Share request for delegation-service.
type ShareDelegateRequest struct {
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

// ShareDelegateResponse is the Share response for delegation-service.
type ShareDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RotateDelegateRequest is the Rotate request for delegation-service.
type RotateDelegateRequest struct {
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

// RotateDelegateResponse is the Rotate response for delegation-service.
type RotateDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ApproveDelegateRequest is the Approve request for delegation-service.
type ApproveDelegateRequest struct {
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

// ApproveDelegateResponse is the Approve response for delegation-service.
type ApproveDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DenyDelegateRequest is the Deny request for delegation-service.
type DenyDelegateRequest struct {
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

// DenyDelegateResponse is the Deny response for delegation-service.
type DenyDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// SearchDelegateRequest is the Search request for delegation-service.
type SearchDelegateRequest struct {
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

// SearchDelegateResponse is the Search response for delegation-service.
type SearchDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// StatsDelegateRequest is the Stats request for delegation-service.
type StatsDelegateRequest struct {
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

// StatsDelegateResponse is the Stats response for delegation-service.
type StatsDelegateResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *DelegateModel `json:"item,omitempty"`
	Items []DelegateModel `json:"items,omitempty"`
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
