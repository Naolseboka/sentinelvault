package apimodels

import "time"

// WebAuthnCredModel is the API model for passkey-service.
type WebAuthnCredModel struct {
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

// CreateWebAuthnCredRequest is the Create request for passkey-service.
type CreateWebAuthnCredRequest struct {
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

// CreateWebAuthnCredResponse is the Create response for passkey-service.
type CreateWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// GetWebAuthnCredRequest is the Get request for passkey-service.
type GetWebAuthnCredRequest struct {
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

// GetWebAuthnCredResponse is the Get response for passkey-service.
type GetWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ListWebAuthnCredRequest is the List request for passkey-service.
type ListWebAuthnCredRequest struct {
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

// ListWebAuthnCredResponse is the List response for passkey-service.
type ListWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// UpdateWebAuthnCredRequest is the Update request for passkey-service.
type UpdateWebAuthnCredRequest struct {
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

// UpdateWebAuthnCredResponse is the Update response for passkey-service.
type UpdateWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// DeleteWebAuthnCredRequest is the Delete request for passkey-service.
type DeleteWebAuthnCredRequest struct {
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

// DeleteWebAuthnCredResponse is the Delete response for passkey-service.
type DeleteWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ArchiveWebAuthnCredRequest is the Archive request for passkey-service.
type ArchiveWebAuthnCredRequest struct {
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

// ArchiveWebAuthnCredResponse is the Archive response for passkey-service.
type ArchiveWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// RestoreWebAuthnCredRequest is the Restore request for passkey-service.
type RestoreWebAuthnCredRequest struct {
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

// RestoreWebAuthnCredResponse is the Restore response for passkey-service.
type RestoreWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ExportWebAuthnCredRequest is the Export request for passkey-service.
type ExportWebAuthnCredRequest struct {
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

// ExportWebAuthnCredResponse is the Export response for passkey-service.
type ExportWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ImportWebAuthnCredRequest is the Import request for passkey-service.
type ImportWebAuthnCredRequest struct {
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

// ImportWebAuthnCredResponse is the Import response for passkey-service.
type ImportWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ShareWebAuthnCredRequest is the Share request for passkey-service.
type ShareWebAuthnCredRequest struct {
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

// ShareWebAuthnCredResponse is the Share response for passkey-service.
type ShareWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// RotateWebAuthnCredRequest is the Rotate request for passkey-service.
type RotateWebAuthnCredRequest struct {
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

// RotateWebAuthnCredResponse is the Rotate response for passkey-service.
type RotateWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ApproveWebAuthnCredRequest is the Approve request for passkey-service.
type ApproveWebAuthnCredRequest struct {
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

// ApproveWebAuthnCredResponse is the Approve response for passkey-service.
type ApproveWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// DenyWebAuthnCredRequest is the Deny request for passkey-service.
type DenyWebAuthnCredRequest struct {
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

// DenyWebAuthnCredResponse is the Deny response for passkey-service.
type DenyWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// SearchWebAuthnCredRequest is the Search request for passkey-service.
type SearchWebAuthnCredRequest struct {
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

// SearchWebAuthnCredResponse is the Search response for passkey-service.
type SearchWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// StatsWebAuthnCredRequest is the Stats request for passkey-service.
type StatsWebAuthnCredRequest struct {
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

// StatsWebAuthnCredResponse is the Stats response for passkey-service.
type StatsWebAuthnCredResponse struct {
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Item      *WebAuthnCredModel  `json:"item,omitempty"`
	Items     []WebAuthnCredModel `json:"items,omitempty"`
	Total     int64               `json:"total,omitempty"`
	Warnings  []string            `json:"warnings,omitempty"`
	RequestID string              `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
