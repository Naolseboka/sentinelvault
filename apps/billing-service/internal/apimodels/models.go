package apimodels

import "time"

// BillingModel is the API model for billing-service.
type BillingModel struct {
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

// CreateBillingRequest is the Create request for billing-service.
type CreateBillingRequest struct {
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

// CreateBillingResponse is the Create response for billing-service.
type CreateBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// GetBillingRequest is the Get request for billing-service.
type GetBillingRequest struct {
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

// GetBillingResponse is the Get response for billing-service.
type GetBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ListBillingRequest is the List request for billing-service.
type ListBillingRequest struct {
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

// ListBillingResponse is the List response for billing-service.
type ListBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// UpdateBillingRequest is the Update request for billing-service.
type UpdateBillingRequest struct {
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

// UpdateBillingResponse is the Update response for billing-service.
type UpdateBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// DeleteBillingRequest is the Delete request for billing-service.
type DeleteBillingRequest struct {
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

// DeleteBillingResponse is the Delete response for billing-service.
type DeleteBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ArchiveBillingRequest is the Archive request for billing-service.
type ArchiveBillingRequest struct {
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

// ArchiveBillingResponse is the Archive response for billing-service.
type ArchiveBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// RestoreBillingRequest is the Restore request for billing-service.
type RestoreBillingRequest struct {
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

// RestoreBillingResponse is the Restore response for billing-service.
type RestoreBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ExportBillingRequest is the Export request for billing-service.
type ExportBillingRequest struct {
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

// ExportBillingResponse is the Export response for billing-service.
type ExportBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ImportBillingRequest is the Import request for billing-service.
type ImportBillingRequest struct {
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

// ImportBillingResponse is the Import response for billing-service.
type ImportBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ShareBillingRequest is the Share request for billing-service.
type ShareBillingRequest struct {
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

// ShareBillingResponse is the Share response for billing-service.
type ShareBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// RotateBillingRequest is the Rotate request for billing-service.
type RotateBillingRequest struct {
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

// RotateBillingResponse is the Rotate response for billing-service.
type RotateBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ApproveBillingRequest is the Approve request for billing-service.
type ApproveBillingRequest struct {
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

// ApproveBillingResponse is the Approve response for billing-service.
type ApproveBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// DenyBillingRequest is the Deny request for billing-service.
type DenyBillingRequest struct {
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

// DenyBillingResponse is the Deny response for billing-service.
type DenyBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// SearchBillingRequest is the Search request for billing-service.
type SearchBillingRequest struct {
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

// SearchBillingResponse is the Search response for billing-service.
type SearchBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// StatsBillingRequest is the Stats request for billing-service.
type StatsBillingRequest struct {
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

// StatsBillingResponse is the Stats response for billing-service.
type StatsBillingResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *BillingModel  `json:"item,omitempty"`
	Items     []BillingModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
