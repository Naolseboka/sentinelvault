package vault

import "time"

// Collection groups vault items within an organization.
type Collection struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Folder organizes items inside a collection.
type Folder struct {
	ID           string    `json:"id"`
	OrgID        string    `json:"org_id"`
	CollectionID string    `json:"collection_id"`
	ParentID     string    `json:"parent_id,omitempty"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ShareGrant grants time-limited access to a vault item.
type ShareGrant struct {
	ID         string     `json:"id"`
	OrgID      string     `json:"org_id"`
	ItemID     string     `json:"item_id"`
	GranteeID  string     `json:"grantee_id"`
	Permission string     `json:"permission"`
	ExpiresAt  time.Time  `json:"expires_at"`
	CreatedAt  time.Time  `json:"created_at"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
}

// IsActive reports whether the grant is currently valid.
func (g *ShareGrant) IsActive(now time.Time) bool {
	if g == nil || g.RevokedAt != nil {
		return false
	}
	return now.Before(g.ExpiresAt)
}
