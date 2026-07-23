package apimodels

import "time"

// PolicyModel is the API model for policy-service.
type PolicyModel struct {
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

// CreatePolicyRequest is the Create request for policy-service.
type CreatePolicyRequest struct {
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

// CreatePolicyResponse is the Create response for policy-service.
type CreatePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// GetPolicyRequest is the Get request for policy-service.
type GetPolicyRequest struct {
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

// GetPolicyResponse is the Get response for policy-service.
type GetPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ListPolicyRequest is the List request for policy-service.
type ListPolicyRequest struct {
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

// ListPolicyResponse is the List response for policy-service.
type ListPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// UpdatePolicyRequest is the Update request for policy-service.
type UpdatePolicyRequest struct {
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

// UpdatePolicyResponse is the Update response for policy-service.
type UpdatePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DeletePolicyRequest is the Delete request for policy-service.
type DeletePolicyRequest struct {
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

// DeletePolicyResponse is the Delete response for policy-service.
type DeletePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ArchivePolicyRequest is the Archive request for policy-service.
type ArchivePolicyRequest struct {
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

// ArchivePolicyResponse is the Archive response for policy-service.
type ArchivePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RestorePolicyRequest is the Restore request for policy-service.
type RestorePolicyRequest struct {
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

// RestorePolicyResponse is the Restore response for policy-service.
type RestorePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ExportPolicyRequest is the Export request for policy-service.
type ExportPolicyRequest struct {
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

// ExportPolicyResponse is the Export response for policy-service.
type ExportPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ImportPolicyRequest is the Import request for policy-service.
type ImportPolicyRequest struct {
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

// ImportPolicyResponse is the Import response for policy-service.
type ImportPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// SharePolicyRequest is the Share request for policy-service.
type SharePolicyRequest struct {
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

// SharePolicyResponse is the Share response for policy-service.
type SharePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// RotatePolicyRequest is the Rotate request for policy-service.
type RotatePolicyRequest struct {
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

// RotatePolicyResponse is the Rotate response for policy-service.
type RotatePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// ApprovePolicyRequest is the Approve request for policy-service.
type ApprovePolicyRequest struct {
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

// ApprovePolicyResponse is the Approve response for policy-service.
type ApprovePolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// DenyPolicyRequest is the Deny request for policy-service.
type DenyPolicyRequest struct {
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

// DenyPolicyResponse is the Deny response for policy-service.
type DenyPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// SearchPolicyRequest is the Search request for policy-service.
type SearchPolicyRequest struct {
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

// SearchPolicyResponse is the Search response for policy-service.
type SearchPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
	Total     int64         `json:"total,omitempty"`
	Warnings  []string      `json:"warnings,omitempty"`
	RequestID string        `json:"request_id,omitempty"`
}

// StatsPolicyRequest is the Stats request for policy-service.
type StatsPolicyRequest struct {
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

// StatsPolicyResponse is the Stats response for policy-service.
type StatsPolicyResponse struct {
	Success   bool          `json:"success"`
	Error     string        `json:"error,omitempty"`
	Item      *PolicyModel  `json:"item,omitempty"`
	Items     []PolicyModel `json:"items,omitempty"`
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
