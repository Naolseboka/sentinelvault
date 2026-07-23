package apimodels

import "time"

// DeviceModel is the API model for device-service.
type DeviceModel struct {
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

// CreateDeviceRequest is the Create request for device-service.
type CreateDeviceRequest struct {
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

// CreateDeviceResponse is the Create response for device-service.
type CreateDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// GetDeviceRequest is the Get request for device-service.
type GetDeviceRequest struct {
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

// GetDeviceResponse is the Get response for device-service.
type GetDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ListDeviceRequest is the List request for device-service.
type ListDeviceRequest struct {
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

// ListDeviceResponse is the List response for device-service.
type ListDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// UpdateDeviceRequest is the Update request for device-service.
type UpdateDeviceRequest struct {
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

// UpdateDeviceResponse is the Update response for device-service.
type UpdateDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DeleteDeviceRequest is the Delete request for device-service.
type DeleteDeviceRequest struct {
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

// DeleteDeviceResponse is the Delete response for device-service.
type DeleteDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ArchiveDeviceRequest is the Archive request for device-service.
type ArchiveDeviceRequest struct {
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

// ArchiveDeviceResponse is the Archive response for device-service.
type ArchiveDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RestoreDeviceRequest is the Restore request for device-service.
type RestoreDeviceRequest struct {
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

// RestoreDeviceResponse is the Restore response for device-service.
type RestoreDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ExportDeviceRequest is the Export request for device-service.
type ExportDeviceRequest struct {
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

// ExportDeviceResponse is the Export response for device-service.
type ExportDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ImportDeviceRequest is the Import request for device-service.
type ImportDeviceRequest struct {
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

// ImportDeviceResponse is the Import response for device-service.
type ImportDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ShareDeviceRequest is the Share request for device-service.
type ShareDeviceRequest struct {
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

// ShareDeviceResponse is the Share response for device-service.
type ShareDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RotateDeviceRequest is the Rotate request for device-service.
type RotateDeviceRequest struct {
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

// RotateDeviceResponse is the Rotate response for device-service.
type RotateDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ApproveDeviceRequest is the Approve request for device-service.
type ApproveDeviceRequest struct {
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

// ApproveDeviceResponse is the Approve response for device-service.
type ApproveDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DenyDeviceRequest is the Deny request for device-service.
type DenyDeviceRequest struct {
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

// DenyDeviceResponse is the Deny response for device-service.
type DenyDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// SearchDeviceRequest is the Search request for device-service.
type SearchDeviceRequest struct {
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

// SearchDeviceResponse is the Search response for device-service.
type SearchDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// StatsDeviceRequest is the Stats request for device-service.
type StatsDeviceRequest struct {
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

// StatsDeviceResponse is the Stats response for device-service.
type StatsDeviceResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *DeviceModel  `json:"item,omitempty"`
	Items     []DeviceModel `json:"items,omitempty"`
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
