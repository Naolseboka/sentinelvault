package apimodels

import "time"

// ThreatModel is the API model for threat-service.
type ThreatModel struct {
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

// CreateThreatRequest is the Create request for threat-service.
type CreateThreatRequest struct {
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

// CreateThreatResponse is the Create response for threat-service.
type CreateThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// GetThreatRequest is the Get request for threat-service.
type GetThreatRequest struct {
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

// GetThreatResponse is the Get response for threat-service.
type GetThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ListThreatRequest is the List request for threat-service.
type ListThreatRequest struct {
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

// ListThreatResponse is the List response for threat-service.
type ListThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// UpdateThreatRequest is the Update request for threat-service.
type UpdateThreatRequest struct {
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

// UpdateThreatResponse is the Update response for threat-service.
type UpdateThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DeleteThreatRequest is the Delete request for threat-service.
type DeleteThreatRequest struct {
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

// DeleteThreatResponse is the Delete response for threat-service.
type DeleteThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ArchiveThreatRequest is the Archive request for threat-service.
type ArchiveThreatRequest struct {
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

// ArchiveThreatResponse is the Archive response for threat-service.
type ArchiveThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RestoreThreatRequest is the Restore request for threat-service.
type RestoreThreatRequest struct {
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

// RestoreThreatResponse is the Restore response for threat-service.
type RestoreThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ExportThreatRequest is the Export request for threat-service.
type ExportThreatRequest struct {
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

// ExportThreatResponse is the Export response for threat-service.
type ExportThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ImportThreatRequest is the Import request for threat-service.
type ImportThreatRequest struct {
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

// ImportThreatResponse is the Import response for threat-service.
type ImportThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ShareThreatRequest is the Share request for threat-service.
type ShareThreatRequest struct {
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

// ShareThreatResponse is the Share response for threat-service.
type ShareThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RotateThreatRequest is the Rotate request for threat-service.
type RotateThreatRequest struct {
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

// RotateThreatResponse is the Rotate response for threat-service.
type RotateThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ApproveThreatRequest is the Approve request for threat-service.
type ApproveThreatRequest struct {
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

// ApproveThreatResponse is the Approve response for threat-service.
type ApproveThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DenyThreatRequest is the Deny request for threat-service.
type DenyThreatRequest struct {
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

// DenyThreatResponse is the Deny response for threat-service.
type DenyThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// SearchThreatRequest is the Search request for threat-service.
type SearchThreatRequest struct {
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

// SearchThreatResponse is the Search response for threat-service.
type SearchThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// StatsThreatRequest is the Stats request for threat-service.
type StatsThreatRequest struct {
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

// StatsThreatResponse is the Stats response for threat-service.
type StatsThreatResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *ThreatModel  `json:"item,omitempty"`
	Items     []ThreatModel `json:"items,omitempty"`
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
