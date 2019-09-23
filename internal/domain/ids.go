package domain

// OrgID identifies an organization.
type OrgID string

// TeamID identifies a team within an organization.
type TeamID string

// RoleID identifies an RBAC role.
type RoleID string

// UserID identifies a platform user.
type UserID string

// Org represents organization metadata shared across services.
type Org struct {
	ID   OrgID  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Team represents a team within an org.
type Team struct {
	ID    TeamID `json:"id"`
	OrgID OrgID  `json:"org_id"`
	Name  string `json:"name"`
}

// Role represents a named permission bundle.
type Role struct {
	ID          RoleID   `json:"id"`
	OrgID       OrgID    `json:"org_id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
