package validators

// PayloadValidator validates vault item payloads by type.
type PayloadValidator interface {
	Type() string
	Validate(p map[string]any) (normalized map[string]any, warnings []string, err error)
	Redact(p map[string]any) map[string]any
	RequiredFields() []string
	OptionalFields() []string
}

// ForType returns a validator for a vault item type slug.
func ForType(itemType string) PayloadValidator {
	switch itemType {
	case "password":
		return PasswordValidator{}
	case "secure_note":
		return SecureNoteValidator{}
	case "credit_card":
		return CreditCardValidator{}
	case "bank_account":
		return BankAccountValidator{}
	case "passport":
		return PassportValidator{}
	case "identity":
		return IdentityValidator{}
	case "license":
		return LicenseValidator{}
	case "recovery_codes":
		return RecoveryCodesValidator{}
	case "api_key":
		return APIKeyValidator{}
	case "ssh_key":
		return SSHKeyValidator{}
	case "certificate":
		return CertificateValidator{}
	case "token":
		return TokenValidator{}
	case "secret":
		return GenericSecretValidator{}
	case "database_credential":
		return DatabaseCredentialValidator{}
	case "env_var":
		return EnvVarValidator{}
	case "private_key":
		return PrivateKeyValidator{}
	case "pgp_key":
		return PGPKeyValidator{}
	case "encryption_key":
		return EncryptionKeyValidator{}
	case "software_license":
		return SoftwareLicenseValidator{}
	case "otp_seed":
		return OTPSeedValidator{}
	case "passkey":
		return PasskeyValidator{}
	default:
		return GenericSecretValidator{}
	}
}

// AllTypes returns supported item type slugs.
func AllTypes() []string {
	return []string{"password", "secure_note", "credit_card", "bank_account", "passport", "identity", "license", "recovery_codes", "api_key", "ssh_key", "certificate", "token", "secret", "database_credential", "env_var", "private_key", "pgp_key", "encryption_key", "software_license", "otp_seed", "passkey"}
}
