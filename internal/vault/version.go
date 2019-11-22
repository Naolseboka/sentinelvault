package vault

import "time"

// VersionSnapshot captures an immutable point-in-time copy of a vault item.
type VersionSnapshot struct {
	ID         string       `json:"id"`
	ItemID     string       `json:"item_id"`
	OrgID      string       `json:"org_id"`
	Version    int          `json:"version"`
	Type       ItemType     `json:"type"`
	Name       string       `json:"name"`
	Metadata   ItemMetadata `json:"metadata"`
	Ciphertext []byte       `json:"ciphertext"`
	KeyID      string       `json:"key_id"`
	CreatedBy  string       `json:"created_by,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
}
