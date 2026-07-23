package apimodels

import "time"

// DirSyncModel is the API model for directory-sync-service.
type DirSyncModel struct {
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

// CreateDirSyncRequest is the Create request for directory-sync-service.
type CreateDirSyncRequest struct {
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

// CreateDirSyncResponse is the Create response for directory-sync-service.
type CreateDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// GetDirSyncRequest is the Get request for directory-sync-service.
type GetDirSyncRequest struct {
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

// GetDirSyncResponse is the Get response for directory-sync-service.
type GetDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ListDirSyncRequest is the List request for directory-sync-service.
type ListDirSyncRequest struct {
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

// ListDirSyncResponse is the List response for directory-sync-service.
type ListDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// UpdateDirSyncRequest is the Update request for directory-sync-service.
type UpdateDirSyncRequest struct {
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

// UpdateDirSyncResponse is the Update response for directory-sync-service.
type UpdateDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// DeleteDirSyncRequest is the Delete request for directory-sync-service.
type DeleteDirSyncRequest struct {
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

// DeleteDirSyncResponse is the Delete response for directory-sync-service.
type DeleteDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ArchiveDirSyncRequest is the Archive request for directory-sync-service.
type ArchiveDirSyncRequest struct {
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

// ArchiveDirSyncResponse is the Archive response for directory-sync-service.
type ArchiveDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// RestoreDirSyncRequest is the Restore request for directory-sync-service.
type RestoreDirSyncRequest struct {
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

// RestoreDirSyncResponse is the Restore response for directory-sync-service.
type RestoreDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ExportDirSyncRequest is the Export request for directory-sync-service.
type ExportDirSyncRequest struct {
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

// ExportDirSyncResponse is the Export response for directory-sync-service.
type ExportDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ImportDirSyncRequest is the Import request for directory-sync-service.
type ImportDirSyncRequest struct {
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

// ImportDirSyncResponse is the Import response for directory-sync-service.
type ImportDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ShareDirSyncRequest is the Share request for directory-sync-service.
type ShareDirSyncRequest struct {
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

// ShareDirSyncResponse is the Share response for directory-sync-service.
type ShareDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// RotateDirSyncRequest is the Rotate request for directory-sync-service.
type RotateDirSyncRequest struct {
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

// RotateDirSyncResponse is the Rotate response for directory-sync-service.
type RotateDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// ApproveDirSyncRequest is the Approve request for directory-sync-service.
type ApproveDirSyncRequest struct {
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

// ApproveDirSyncResponse is the Approve response for directory-sync-service.
type ApproveDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// DenyDirSyncRequest is the Deny request for directory-sync-service.
type DenyDirSyncRequest struct {
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

// DenyDirSyncResponse is the Deny response for directory-sync-service.
type DenyDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// SearchDirSyncRequest is the Search request for directory-sync-service.
type SearchDirSyncRequest struct {
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

// SearchDirSyncResponse is the Search response for directory-sync-service.
type SearchDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
	Total     int64          `json:"total,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

// StatsDirSyncRequest is the Stats request for directory-sync-service.
type StatsDirSyncRequest struct {
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

// StatsDirSyncResponse is the Stats response for directory-sync-service.
type StatsDirSyncResponse struct {
	Success   bool           `json:"success"`
	Error     string         `json:"error,omitempty"`
	Item      *DirSyncModel  `json:"item,omitempty"`
	Items     []DirSyncModel `json:"items,omitempty"`
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

// maintenance note 35
