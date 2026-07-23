package searchx

func ExtractPassword(payload map[string]any) map[string]string {
	out := map[string]string{"type": "password"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractSecureNote(payload map[string]any) map[string]string {
	out := map[string]string{"type": "secure_note"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractCreditCard(payload map[string]any) map[string]string {
	out := map[string]string{"type": "credit_card"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractBankAccount(payload map[string]any) map[string]string {
	out := map[string]string{"type": "bank_account"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractPassport(payload map[string]any) map[string]string {
	out := map[string]string{"type": "passport"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractIdentity(payload map[string]any) map[string]string {
	out := map[string]string{"type": "identity"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractLicense(payload map[string]any) map[string]string {
	out := map[string]string{"type": "license"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractRecoveryCodes(payload map[string]any) map[string]string {
	out := map[string]string{"type": "recovery_codes"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractApiKey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "api_key"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractSshKey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "ssh_key"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractCertificate(payload map[string]any) map[string]string {
	out := map[string]string{"type": "certificate"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractToken(payload map[string]any) map[string]string {
	out := map[string]string{"type": "token"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractSecret(payload map[string]any) map[string]string {
	out := map[string]string{"type": "secret"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractDatabaseCredential(payload map[string]any) map[string]string {
	out := map[string]string{"type": "database_credential"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractEnvVar(payload map[string]any) map[string]string {
	out := map[string]string{"type": "env_var"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractPrivateKey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "private_key"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractPgpKey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "pgp_key"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractEncryptionKey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "encryption_key"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractSoftwareLicense(payload map[string]any) map[string]string {
	out := map[string]string{"type": "software_license"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractOtpSeed(payload map[string]any) map[string]string {
	out := map[string]string{"type": "otp_seed"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractPasskey(payload map[string]any) map[string]string {
	out := map[string]string{"type": "passkey"}
	for _, k := range []string{"name", "username", "url", "issuer", "host", "provider", "title", "product", "account"} {
		if v, ok := payload[k].(string); ok && v != "" {
			out[k] = v
		}
	}
	return out
}

func ExtractForType(itemType string, payload map[string]any) map[string]string {
	switch itemType {
	case "password":
		return ExtractPassword(payload)
	case "secure_note":
		return ExtractSecureNote(payload)
	case "credit_card":
		return ExtractCreditCard(payload)
	case "bank_account":
		return ExtractBankAccount(payload)
	case "passport":
		return ExtractPassport(payload)
	case "identity":
		return ExtractIdentity(payload)
	case "license":
		return ExtractLicense(payload)
	case "recovery_codes":
		return ExtractRecoveryCodes(payload)
	case "api_key":
		return ExtractApiKey(payload)
	case "ssh_key":
		return ExtractSshKey(payload)
	case "certificate":
		return ExtractCertificate(payload)
	case "token":
		return ExtractToken(payload)
	case "secret":
		return ExtractSecret(payload)
	case "database_credential":
		return ExtractDatabaseCredential(payload)
	case "env_var":
		return ExtractEnvVar(payload)
	case "private_key":
		return ExtractPrivateKey(payload)
	case "pgp_key":
		return ExtractPgpKey(payload)
	case "encryption_key":
		return ExtractEncryptionKey(payload)
	case "software_license":
		return ExtractSoftwareLicense(payload)
	case "otp_seed":
		return ExtractOtpSeed(payload)
	case "passkey":
		return ExtractPasskey(payload)
	default:
		return map[string]string{"type": itemType}
	}
}
