package apimodels

import "time"

// RevModel is the API model for versioning-service.
type RevModel struct {
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

// CreateRevRequest is the Create request for versioning-service.
type CreateRevRequest struct {
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

// CreateRevResponse is the Create response for versioning-service.
type CreateRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// GetRevRequest is the Get request for versioning-service.
type GetRevRequest struct {
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

// GetRevResponse is the Get response for versioning-service.
type GetRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ListRevRequest is the List request for versioning-service.
type ListRevRequest struct {
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

// ListRevResponse is the List response for versioning-service.
type ListRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// UpdateRevRequest is the Update request for versioning-service.
type UpdateRevRequest struct {
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

// UpdateRevResponse is the Update response for versioning-service.
type UpdateRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// DeleteRevRequest is the Delete request for versioning-service.
type DeleteRevRequest struct {
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

// DeleteRevResponse is the Delete response for versioning-service.
type DeleteRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ArchiveRevRequest is the Archive request for versioning-service.
type ArchiveRevRequest struct {
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

// ArchiveRevResponse is the Archive response for versioning-service.
type ArchiveRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// RestoreRevRequest is the Restore request for versioning-service.
type RestoreRevRequest struct {
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

// RestoreRevResponse is the Restore response for versioning-service.
type RestoreRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ExportRevRequest is the Export request for versioning-service.
type ExportRevRequest struct {
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

// ExportRevResponse is the Export response for versioning-service.
type ExportRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ImportRevRequest is the Import request for versioning-service.
type ImportRevRequest struct {
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

// ImportRevResponse is the Import response for versioning-service.
type ImportRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ShareRevRequest is the Share request for versioning-service.
type ShareRevRequest struct {
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

// ShareRevResponse is the Share response for versioning-service.
type ShareRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// RotateRevRequest is the Rotate request for versioning-service.
type RotateRevRequest struct {
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

// RotateRevResponse is the Rotate response for versioning-service.
type RotateRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ApproveRevRequest is the Approve request for versioning-service.
type ApproveRevRequest struct {
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

// ApproveRevResponse is the Approve response for versioning-service.
type ApproveRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// DenyRevRequest is the Deny request for versioning-service.
type DenyRevRequest struct {
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

// DenyRevResponse is the Deny response for versioning-service.
type DenyRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// SearchRevRequest is the Search request for versioning-service.
type SearchRevRequest struct {
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

// SearchRevResponse is the Search response for versioning-service.
type SearchRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// StatsRevRequest is the Stats request for versioning-service.
type StatsRevRequest struct {
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

// StatsRevResponse is the Stats response for versioning-service.
type StatsRevResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error,omitempty"`
	Item      *RevModel  `json:"item,omitempty"`
	Items     []RevModel `json:"items,omitempty"`
	Total     int64      `json:"total,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

// ErrorBody is a standard API error.
type ErrorBody struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
