package vault

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RequiredFields maps item types to mandatory payload field names.
var RequiredFields = map[ItemType][]string{
	ItemTypePassword:           {"password"},
	ItemTypeSecureNote:         {"note"},
	ItemTypeCreditCard:         {"number", "expiry"},
	ItemTypeBankAccount:        {"account_number", "routing_number"},
	ItemTypePassport:           {"passport_number", "country"},
	ItemTypeIdentity:           {"first_name", "last_name"},
	ItemTypeLicense:            {"license_number", "issuer"},
	ItemTypeRecoveryCodes:      {"codes"},
	ItemTypeAPIKey:             {"api_key"},
	ItemTypeSSHKey:             {"private_key"},
	ItemTypeCertificate:        {"certificate"},
	ItemTypeToken:              {"token"},
	ItemTypeSecret:             {"secret"},
	ItemTypeDatabaseCredential: {"username", "password", "host"},
	ItemTypeEnvVar:             {"name", "value"},
	ItemTypePrivateKey:         {"private_key"},
	ItemTypePGPKey:             {"private_key"},
	ItemTypeEncryptionKey:      {"key_material"},
	ItemTypeSoftwareLicense:    {"license_key"},
	ItemTypeOTPSeed:            {"seed"},
	ItemTypePasskey:            {"credential_id", "private_key"},
}

// ValidateItemType ensures the type is known.
func ValidateItemType(t ItemType) error {
	if _, ok := RequiredFields[t]; !ok {
		return common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("unknown item type: %s", t))
	}
	return nil
}

// ValidatePayload checks required secret fields for an item type.
func ValidatePayload(t ItemType, payload ItemPayload) error {
	if err := ValidateItemType(t); err != nil {
		return err
	}
	for _, field := range RequiredFields[t] {
		val, ok := payload[field]
		if !ok || strings.TrimSpace(val) == "" {
			return common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("missing required field: %s", field))
		}
	}
	return nil
}

// ValidateMetadata rejects plaintext secrets accidentally placed in metadata.
func ValidateMetadata(t ItemType, meta ItemMetadata) error {
	forbidden := secretFieldNames(t)
	for _, field := range forbidden {
		if valueFromMetadata(meta, field) != "" {
			return common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("sensitive field %q must not be stored in metadata", field))
		}
	}
	return nil
}

func secretFieldNames(t ItemType) []string {
	fields := RequiredFields[t]
	out := make([]string, len(fields))
	copy(out, fields)
	return out
}

func valueFromMetadata(meta ItemMetadata, field string) string {
	if meta.CustomFields != nil {
		if v, ok := meta.CustomFields[field]; ok {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

// IsHumanType reports whether the type belongs to password-service.
func IsHumanType(t ItemType) bool {
	for _, ht := range HumanItemTypes() {
		if ht == t {
			return true
		}
	}
	return false
}

// IsMachineType reports whether the type belongs to secret-service.
func IsMachineType(t ItemType) bool {
	for _, mt := range MachineItemTypes() {
		if mt == t {
			return true
		}
	}
	return false
}
