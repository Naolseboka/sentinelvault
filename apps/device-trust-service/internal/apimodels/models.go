package apimodels

import "time"

// TrustModel is the API model for device-trust-service.
type TrustModel struct {
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

// CreateTrustRequest is the Create request for device-trust-service.
type CreateTrustRequest struct {
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

// CreateTrustResponse is the Create response for device-trust-service.
type CreateTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// GetTrustRequest is the Get request for device-trust-service.
type GetTrustRequest struct {
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

// GetTrustResponse is the Get response for device-trust-service.
type GetTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ListTrustRequest is the List request for device-trust-service.
type ListTrustRequest struct {
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

// ListTrustResponse is the List response for device-trust-service.
type ListTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// UpdateTrustRequest is the Update request for device-trust-service.
type UpdateTrustRequest struct {
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

// UpdateTrustResponse is the Update response for device-trust-service.
type UpdateTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DeleteTrustRequest is the Delete request for device-trust-service.
type DeleteTrustRequest struct {
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

// DeleteTrustResponse is the Delete response for device-trust-service.
type DeleteTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ArchiveTrustRequest is the Archive request for device-trust-service.
type ArchiveTrustRequest struct {
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

// ArchiveTrustResponse is the Archive response for device-trust-service.
type ArchiveTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RestoreTrustRequest is the Restore request for device-trust-service.
type RestoreTrustRequest struct {
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

// RestoreTrustResponse is the Restore response for device-trust-service.
type RestoreTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ExportTrustRequest is the Export request for device-trust-service.
type ExportTrustRequest struct {
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

// ExportTrustResponse is the Export response for device-trust-service.
type ExportTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ImportTrustRequest is the Import request for device-trust-service.
type ImportTrustRequest struct {
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

// ImportTrustResponse is the Import response for device-trust-service.
type ImportTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ShareTrustRequest is the Share request for device-trust-service.
type ShareTrustRequest struct {
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

// ShareTrustResponse is the Share response for device-trust-service.
type ShareTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// RotateTrustRequest is the Rotate request for device-trust-service.
type RotateTrustRequest struct {
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

// RotateTrustResponse is the Rotate response for device-trust-service.
type RotateTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// ApproveTrustRequest is the Approve request for device-trust-service.
type ApproveTrustRequest struct {
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

// ApproveTrustResponse is the Approve response for device-trust-service.
type ApproveTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// DenyTrustRequest is the Deny request for device-trust-service.
type DenyTrustRequest struct {
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

// DenyTrustResponse is the Deny response for device-trust-service.
type DenyTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// SearchTrustRequest is the Search request for device-trust-service.
type SearchTrustRequest struct {
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

// SearchTrustResponse is the Search response for device-trust-service.
type SearchTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
	Total     int64        `json:"total,omitempty"`
	Warnings  []string     `json:"warnings,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

// StatsTrustRequest is the Stats request for device-trust-service.
type StatsTrustRequest struct {
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

// StatsTrustResponse is the Stats response for device-trust-service.
type StatsTrustResponse struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error,omitempty"`
	Item      *TrustModel  `json:"item,omitempty"`
	Items     []TrustModel `json:"items,omitempty"`
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
