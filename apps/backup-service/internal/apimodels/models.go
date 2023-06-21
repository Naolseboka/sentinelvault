package apimodels

import "time"

// BackupModel is the API model for backup-service.
type BackupModel struct {
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

// CreateBackupRequest is the Create request for backup-service.
type CreateBackupRequest struct {
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

// CreateBackupResponse is the Create response for backup-service.
type CreateBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// GetBackupRequest is the Get request for backup-service.
type GetBackupRequest struct {
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

// GetBackupResponse is the Get response for backup-service.
type GetBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ListBackupRequest is the List request for backup-service.
type ListBackupRequest struct {
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

// ListBackupResponse is the List response for backup-service.
type ListBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// UpdateBackupRequest is the Update request for backup-service.
type UpdateBackupRequest struct {
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

// UpdateBackupResponse is the Update response for backup-service.
type UpdateBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DeleteBackupRequest is the Delete request for backup-service.
type DeleteBackupRequest struct {
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

// DeleteBackupResponse is the Delete response for backup-service.
type DeleteBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ArchiveBackupRequest is the Archive request for backup-service.
type ArchiveBackupRequest struct {
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

// ArchiveBackupResponse is the Archive response for backup-service.
type ArchiveBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RestoreBackupRequest is the Restore request for backup-service.
type RestoreBackupRequest struct {
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

// RestoreBackupResponse is the Restore response for backup-service.
type RestoreBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ExportBackupRequest is the Export request for backup-service.
type ExportBackupRequest struct {
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

// ExportBackupResponse is the Export response for backup-service.
type ExportBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ImportBackupRequest is the Import request for backup-service.
type ImportBackupRequest struct {
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

// ImportBackupResponse is the Import response for backup-service.
type ImportBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ShareBackupRequest is the Share request for backup-service.
type ShareBackupRequest struct {
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

// ShareBackupResponse is the Share response for backup-service.
type ShareBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// RotateBackupRequest is the Rotate request for backup-service.
type RotateBackupRequest struct {
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

// RotateBackupResponse is the Rotate response for backup-service.
type RotateBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// ApproveBackupRequest is the Approve request for backup-service.
type ApproveBackupRequest struct {
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

// ApproveBackupResponse is the Approve response for backup-service.
type ApproveBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// DenyBackupRequest is the Deny request for backup-service.
type DenyBackupRequest struct {
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

// DenyBackupResponse is the Deny response for backup-service.
type DenyBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// SearchBackupRequest is the Search request for backup-service.
type SearchBackupRequest struct {
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

// SearchBackupResponse is the Search response for backup-service.
type SearchBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

// StatsBackupRequest is the Stats request for backup-service.
type StatsBackupRequest struct {
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

// StatsBackupResponse is the Stats response for backup-service.
type StatsBackupResponse struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Item *BackupModel `json:"item,omitempty"`
	Items []BackupModel `json:"items,omitempty"`
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
