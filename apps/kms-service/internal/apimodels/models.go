package apimodels

import "time"

// CryptoKeyModel is the API model for kms-service.
type CryptoKeyModel struct {
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

// CreateCryptoKeyRequest is the Create request for kms-service.
type CreateCryptoKeyRequest struct {
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

// CreateCryptoKeyResponse is the Create response for kms-service.
type CreateCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// GetCryptoKeyRequest is the Get request for kms-service.
type GetCryptoKeyRequest struct {
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

// GetCryptoKeyResponse is the Get response for kms-service.
type GetCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ListCryptoKeyRequest is the List request for kms-service.
type ListCryptoKeyRequest struct {
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

// ListCryptoKeyResponse is the List response for kms-service.
type ListCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// UpdateCryptoKeyRequest is the Update request for kms-service.
type UpdateCryptoKeyRequest struct {
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

// UpdateCryptoKeyResponse is the Update response for kms-service.
type UpdateCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DeleteCryptoKeyRequest is the Delete request for kms-service.
type DeleteCryptoKeyRequest struct {
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

// DeleteCryptoKeyResponse is the Delete response for kms-service.
type DeleteCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ArchiveCryptoKeyRequest is the Archive request for kms-service.
type ArchiveCryptoKeyRequest struct {
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

// ArchiveCryptoKeyResponse is the Archive response for kms-service.
type ArchiveCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RestoreCryptoKeyRequest is the Restore request for kms-service.
type RestoreCryptoKeyRequest struct {
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

// RestoreCryptoKeyResponse is the Restore response for kms-service.
type RestoreCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ExportCryptoKeyRequest is the Export request for kms-service.
type ExportCryptoKeyRequest struct {
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

// ExportCryptoKeyResponse is the Export response for kms-service.
type ExportCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ImportCryptoKeyRequest is the Import request for kms-service.
type ImportCryptoKeyRequest struct {
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

// ImportCryptoKeyResponse is the Import response for kms-service.
type ImportCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ShareCryptoKeyRequest is the Share request for kms-service.
type ShareCryptoKeyRequest struct {
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

// ShareCryptoKeyResponse is the Share response for kms-service.
type ShareCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RotateCryptoKeyRequest is the Rotate request for kms-service.
type RotateCryptoKeyRequest struct {
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

// RotateCryptoKeyResponse is the Rotate response for kms-service.
type RotateCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ApproveCryptoKeyRequest is the Approve request for kms-service.
type ApproveCryptoKeyRequest struct {
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

// ApproveCryptoKeyResponse is the Approve response for kms-service.
type ApproveCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DenyCryptoKeyRequest is the Deny request for kms-service.
type DenyCryptoKeyRequest struct {
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

// DenyCryptoKeyResponse is the Deny response for kms-service.
type DenyCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// SearchCryptoKeyRequest is the Search request for kms-service.
type SearchCryptoKeyRequest struct {
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

// SearchCryptoKeyResponse is the Search response for kms-service.
type SearchCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// StatsCryptoKeyRequest is the Stats request for kms-service.
type StatsCryptoKeyRequest struct {
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

// StatsCryptoKeyResponse is the Stats response for kms-service.
type StatsCryptoKeyResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *CryptoKeyModel  `json:"item,omitempty"`
	Items     []CryptoKeyModel `json:"items,omitempty"`
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
