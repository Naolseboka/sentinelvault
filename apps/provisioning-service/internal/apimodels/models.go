package apimodels

import "time"

// ProvisionModel is the API model for provisioning-service.
type ProvisionModel struct {
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

// CreateProvisionRequest is the Create request for provisioning-service.
type CreateProvisionRequest struct {
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

// CreateProvisionResponse is the Create response for provisioning-service.
type CreateProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// GetProvisionRequest is the Get request for provisioning-service.
type GetProvisionRequest struct {
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

// GetProvisionResponse is the Get response for provisioning-service.
type GetProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ListProvisionRequest is the List request for provisioning-service.
type ListProvisionRequest struct {
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

// ListProvisionResponse is the List response for provisioning-service.
type ListProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// UpdateProvisionRequest is the Update request for provisioning-service.
type UpdateProvisionRequest struct {
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

// UpdateProvisionResponse is the Update response for provisioning-service.
type UpdateProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DeleteProvisionRequest is the Delete request for provisioning-service.
type DeleteProvisionRequest struct {
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

// DeleteProvisionResponse is the Delete response for provisioning-service.
type DeleteProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ArchiveProvisionRequest is the Archive request for provisioning-service.
type ArchiveProvisionRequest struct {
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

// ArchiveProvisionResponse is the Archive response for provisioning-service.
type ArchiveProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RestoreProvisionRequest is the Restore request for provisioning-service.
type RestoreProvisionRequest struct {
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

// RestoreProvisionResponse is the Restore response for provisioning-service.
type RestoreProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ExportProvisionRequest is the Export request for provisioning-service.
type ExportProvisionRequest struct {
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

// ExportProvisionResponse is the Export response for provisioning-service.
type ExportProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ImportProvisionRequest is the Import request for provisioning-service.
type ImportProvisionRequest struct {
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

// ImportProvisionResponse is the Import response for provisioning-service.
type ImportProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ShareProvisionRequest is the Share request for provisioning-service.
type ShareProvisionRequest struct {
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

// ShareProvisionResponse is the Share response for provisioning-service.
type ShareProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// RotateProvisionRequest is the Rotate request for provisioning-service.
type RotateProvisionRequest struct {
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

// RotateProvisionResponse is the Rotate response for provisioning-service.
type RotateProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// ApproveProvisionRequest is the Approve request for provisioning-service.
type ApproveProvisionRequest struct {
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

// ApproveProvisionResponse is the Approve response for provisioning-service.
type ApproveProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// DenyProvisionRequest is the Deny request for provisioning-service.
type DenyProvisionRequest struct {
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

// DenyProvisionResponse is the Deny response for provisioning-service.
type DenyProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// SearchProvisionRequest is the Search request for provisioning-service.
type SearchProvisionRequest struct {
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

// SearchProvisionResponse is the Search response for provisioning-service.
type SearchProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
	Total     int64            `json:"total,omitempty"`
	Warnings  []string         `json:"warnings,omitempty"`
	RequestID string           `json:"request_id,omitempty"`
}

// StatsProvisionRequest is the Stats request for provisioning-service.
type StatsProvisionRequest struct {
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

// StatsProvisionResponse is the Stats response for provisioning-service.
type StatsProvisionResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error,omitempty"`
	Item      *ProvisionModel  `json:"item,omitempty"`
	Items     []ProvisionModel `json:"items,omitempty"`
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
