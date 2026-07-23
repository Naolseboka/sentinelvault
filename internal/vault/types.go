package vault

import "time"

// ItemType classifies vault entries.
type ItemType string

const (
	ItemTypePassword           ItemType = "password"
	ItemTypeSecureNote         ItemType = "secure_note"
	ItemTypeCreditCard         ItemType = "credit_card"
	ItemTypeBankAccount        ItemType = "bank_account"
	ItemTypePassport           ItemType = "passport"
	ItemTypeIdentity           ItemType = "identity"
	ItemTypeLicense            ItemType = "license"
	ItemTypeRecoveryCodes      ItemType = "recovery_codes"
	ItemTypeAPIKey             ItemType = "api_key"
	ItemTypeSSHKey             ItemType = "ssh_key"
	ItemTypeCertificate        ItemType = "certificate"
	ItemTypeToken              ItemType = "token"
	ItemTypeSecret             ItemType = "secret"
	ItemTypeDatabaseCredential ItemType = "database_credential"
	ItemTypeEnvVar             ItemType = "env_var"
	ItemTypePrivateKey         ItemType = "private_key"
	ItemTypePGPKey             ItemType = "pgp_key"
	ItemTypeEncryptionKey      ItemType = "encryption_key"
	ItemTypeSoftwareLicense    ItemType = "software_license"
	ItemTypeOTPSeed            ItemType = "otp_seed"
	ItemTypePasskey            ItemType = "passkey"
)

// HumanItemTypes are vault entries managed by password-service.
func HumanItemTypes() []ItemType {
	return []ItemType{
		ItemTypePassword, ItemTypeSecureNote, ItemTypeCreditCard, ItemTypeBankAccount,
		ItemTypePassport, ItemTypeIdentity, ItemTypeLicense, ItemTypeRecoveryCodes,
		ItemTypeSoftwareLicense, ItemTypeOTPSeed, ItemTypePasskey,
	}
}

// MachineItemTypes are vault entries managed by secret-service.
func MachineItemTypes() []ItemType {
	return []ItemType{
		ItemTypeAPIKey, ItemTypeSSHKey, ItemTypeCertificate, ItemTypeToken,
		ItemTypeSecret, ItemTypeDatabaseCredential, ItemTypeEnvVar, ItemTypePrivateKey,
		ItemTypePGPKey, ItemTypeEncryptionKey,
	}
}

// ItemStatus tracks lifecycle of a vault item.
type ItemStatus string

const (
	ItemStatusActive   ItemStatus = "active"
	ItemStatusArchived ItemStatus = "archived"
	ItemStatusDeleted  ItemStatus = "deleted"
)

// ItemMetadata holds non-sensitive metadata. Never store plaintext secrets here.
type ItemMetadata struct {
	Username     string            `json:"username,omitempty"`
	URL          string            `json:"url,omitempty"`
	Issuer       string            `json:"issuer,omitempty"`
	LastFour     string            `json:"last_four,omitempty"`
	Brand        string            `json:"brand,omitempty"`
	Environment  string            `json:"environment,omitempty"`
	Host         string            `json:"host,omitempty"`
	Port         int               `json:"port,omitempty"`
	Database     string            `json:"database,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// VaultItem holds metadata for an encrypted vault entry.
type VaultItem struct {
	ID           string       `json:"id"`
	OrgID        string       `json:"org_id"`
	Type         ItemType     `json:"type"`
	Name         string       `json:"name"`
	Version      int          `json:"version"`
	Status       ItemStatus   `json:"status"`
	Metadata     ItemMetadata `json:"metadata"`
	CollectionID string       `json:"collection_id,omitempty"`
	FolderID     string       `json:"folder_id,omitempty"`
	KeyID        string       `json:"key_id"`
	CreatedBy    string       `json:"created_by,omitempty"`
	UpdatedBy    string       `json:"updated_by,omitempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at,omitempty"`
}

// StoredRecord is the persisted form: ciphertext only, no plaintext payload fields.
type StoredRecord struct {
	ID           string       `json:"id"`
	OrgID        string       `json:"org_id"`
	Type         ItemType     `json:"type"`
	Name         string       `json:"name"`
	Version      int          `json:"version"`
	Status       ItemStatus   `json:"status"`
	Metadata     ItemMetadata `json:"metadata"`
	Ciphertext   []byte       `json:"ciphertext"`
	KeyID        string       `json:"key_id"`
	CollectionID string       `json:"collection_id,omitempty"`
	FolderID     string       `json:"folder_id,omitempty"`
	CreatedBy    string       `json:"created_by,omitempty"`
	UpdatedBy    string       `json:"updated_by,omitempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at,omitempty"`
}

// ItemPayload is the decrypted secret fields for an item. Stored only in ciphertext at rest.
type ItemPayload map[string]string
