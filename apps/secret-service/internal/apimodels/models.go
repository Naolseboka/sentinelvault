package apimodels

import "time"

// SecretModel is the API model for secret-service.
type SecretModel struct {
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

// CreateSecretRequest is the Create request for secret-service.
type CreateSecretRequest struct {
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

// CreateSecretResponse is the Create response for secret-service.
type CreateSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// GetSecretRequest is the Get request for secret-service.
type GetSecretRequest struct {
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

// GetSecretResponse is the Get response for secret-service.
type GetSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ListSecretRequest is the List request for secret-service.
type ListSecretRequest struct {
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

// ListSecretResponse is the List response for secret-service.
type ListSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// UpdateSecretRequest is the Update request for secret-service.
type UpdateSecretRequest struct {
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

// UpdateSecretResponse is the Update response for secret-service.
type UpdateSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DeleteSecretRequest is the Delete request for secret-service.
type DeleteSecretRequest struct {
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

// DeleteSecretResponse is the Delete response for secret-service.
type DeleteSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ArchiveSecretRequest is the Archive request for secret-service.
type ArchiveSecretRequest struct {
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

// ArchiveSecretResponse is the Archive response for secret-service.
type ArchiveSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RestoreSecretRequest is the Restore request for secret-service.
type RestoreSecretRequest struct {
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

// RestoreSecretResponse is the Restore response for secret-service.
type RestoreSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ExportSecretRequest is the Export request for secret-service.
type ExportSecretRequest struct {
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

// ExportSecretResponse is the Export response for secret-service.
type ExportSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ImportSecretRequest is the Import request for secret-service.
type ImportSecretRequest struct {
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

// ImportSecretResponse is the Import response for secret-service.
type ImportSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ShareSecretRequest is the Share request for secret-service.
type ShareSecretRequest struct {
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

// ShareSecretResponse is the Share response for secret-service.
type ShareSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RotateSecretRequest is the Rotate request for secret-service.
type RotateSecretRequest struct {
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

// RotateSecretResponse is the Rotate response for secret-service.
type RotateSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ApproveSecretRequest is the Approve request for secret-service.
type ApproveSecretRequest struct {
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

// ApproveSecretResponse is the Approve response for secret-service.
type ApproveSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DenySecretRequest is the Deny request for secret-service.
type DenySecretRequest struct {
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

// DenySecretResponse is the Deny response for secret-service.
type DenySecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// SearchSecretRequest is the Search request for secret-service.
type SearchSecretRequest struct {
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

// SearchSecretResponse is the Search response for secret-service.
type SearchSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// StatsSecretRequest is the Stats request for secret-service.
type StatsSecretRequest struct {
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

// StatsSecretResponse is the Stats response for secret-service.
type StatsSecretResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *SecretModel  `json:"item,omitempty"`
	Items     []SecretModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
