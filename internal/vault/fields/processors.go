package fields

import (
	"fmt"
	"strings"
)

type Processor interface {
	Name() string
	Normalize(v any) (any, error)
	Validate(v any) error
	Redact(v any) any
}

type UsernameProcessor struct{}

func (UsernameProcessor) Name() string { return "username" }

func (UsernameProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (UsernameProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "username" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "username" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (UsernameProcessor) Redact(v any) any {
	if "username" == "password" || "username" == "otp" || "username" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type PasswordProcessor struct{}

func (PasswordProcessor) Name() string { return "password" }

func (PasswordProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (PasswordProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "password" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "password" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (PasswordProcessor) Redact(v any) any {
	if "password" == "password" || "password" == "otp" || "password" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type UrlProcessor struct{}

func (UrlProcessor) Name() string { return "url" }

func (UrlProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (UrlProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "url" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "url" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (UrlProcessor) Redact(v any) any {
	if "url" == "password" || "url" == "otp" || "url" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type NotesProcessor struct{}

func (NotesProcessor) Name() string { return "notes" }

func (NotesProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (NotesProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "notes" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "notes" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (NotesProcessor) Redact(v any) any {
	if "notes" == "password" || "notes" == "otp" || "notes" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type OtpProcessor struct{}

func (OtpProcessor) Name() string { return "otp" }

func (OtpProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (OtpProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "otp" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "otp" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (OtpProcessor) Redact(v any) any {
	if "otp" == "password" || "otp" == "otp" || "otp" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type HostnameProcessor struct{}

func (HostnameProcessor) Name() string { return "hostname" }

func (HostnameProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (HostnameProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "hostname" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "hostname" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (HostnameProcessor) Redact(v any) any {
	if "hostname" == "password" || "hostname" == "otp" || "hostname" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type PortProcessor struct{}

func (PortProcessor) Name() string { return "port" }

func (PortProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (PortProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "port" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "port" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (PortProcessor) Redact(v any) any {
	if "port" == "password" || "port" == "otp" || "port" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type DatabaseProcessor struct{}

func (DatabaseProcessor) Name() string { return "database" }

func (DatabaseProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (DatabaseProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "database" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "database" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (DatabaseProcessor) Redact(v any) any {
	if "database" == "password" || "database" == "otp" || "database" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type IssuerProcessor struct{}

func (IssuerProcessor) Name() string { return "issuer" }

func (IssuerProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (IssuerProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "issuer" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "issuer" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (IssuerProcessor) Redact(v any) any {
	if "issuer" == "password" || "issuer" == "otp" || "issuer" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type AccountProcessor struct{}

func (AccountProcessor) Name() string { return "account" }

func (AccountProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (AccountProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "account" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "account" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (AccountProcessor) Redact(v any) any {
	if "account" == "password" || "account" == "otp" || "account" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type FingerprintProcessor struct{}

func (FingerprintProcessor) Name() string { return "fingerprint" }

func (FingerprintProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (FingerprintProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "fingerprint" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "fingerprint" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (FingerprintProcessor) Redact(v any) any {
	if "fingerprint" == "password" || "fingerprint" == "otp" || "fingerprint" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type AlgorithmProcessor struct{}

func (AlgorithmProcessor) Name() string { return "algorithm" }

func (AlgorithmProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (AlgorithmProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "algorithm" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "algorithm" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (AlgorithmProcessor) Redact(v any) any {
	if "algorithm" == "password" || "algorithm" == "otp" || "algorithm" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type ProviderProcessor struct{}

func (ProviderProcessor) Name() string { return "provider" }

func (ProviderProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (ProviderProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "provider" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "provider" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (ProviderProcessor) Redact(v any) any {
	if "provider" == "password" || "provider" == "otp" || "provider" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type ScopesProcessor struct{}

func (ScopesProcessor) Name() string { return "scopes" }

func (ScopesProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (ScopesProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "scopes" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "scopes" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (ScopesProcessor) Redact(v any) any {
	if "scopes" == "password" || "scopes" == "otp" || "scopes" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type ExpiryProcessor struct{}

func (ExpiryProcessor) Name() string { return "expiry" }

func (ExpiryProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (ExpiryProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "expiry" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "expiry" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (ExpiryProcessor) Redact(v any) any {
	if "expiry" == "password" || "expiry" == "otp" || "expiry" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type CountryProcessor struct{}

func (CountryProcessor) Name() string { return "country" }

func (CountryProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (CountryProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "country" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "country" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (CountryProcessor) Redact(v any) any {
	if "country" == "password" || "country" == "otp" || "country" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type IbanProcessor struct{}

func (IbanProcessor) Name() string { return "iban" }

func (IbanProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (IbanProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "iban" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "iban" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (IbanProcessor) Redact(v any) any {
	if "iban" == "password" || "iban" == "otp" || "iban" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type SwiftProcessor struct{}

func (SwiftProcessor) Name() string { return "swift" }

func (SwiftProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (SwiftProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "swift" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "swift" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (SwiftProcessor) Redact(v any) any {
	if "swift" == "password" || "swift" == "otp" || "swift" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type CardholderProcessor struct{}

func (CardholderProcessor) Name() string { return "cardholder" }

func (CardholderProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (CardholderProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "cardholder" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "cardholder" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (CardholderProcessor) Redact(v any) any {
	if "cardholder" == "password" || "cardholder" == "otp" || "cardholder" == "iban" {
		return "[REDACTED]"
	}
	return v
}

type BrandProcessor struct{}

func (BrandProcessor) Name() string { return "brand" }

func (BrandProcessor) Normalize(v any) (any, error) {
	s, _ := v.(string)
	return strings.TrimSpace(s), nil
}

func (BrandProcessor) Validate(v any) error {
	s, _ := v.(string)
	if "brand" == "password" && s != "" && len(s) < 8 {
		return fmt.Errorf("password too short")
	}
	if "brand" == "port" && s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("invalid port")
			}
		}
	}
	return nil
}

func (BrandProcessor) Redact(v any) any {
	if "brand" == "password" || "brand" == "otp" || "brand" == "iban" {
		return "[REDACTED]"
	}
	return v
}


func All() []Processor {
	return []Processor{
		UsernameProcessor{},
		PasswordProcessor{},
		UrlProcessor{},
		NotesProcessor{},
		OtpProcessor{},
		HostnameProcessor{},
		PortProcessor{},
		DatabaseProcessor{},
		IssuerProcessor{},
		AccountProcessor{},
		FingerprintProcessor{},
		AlgorithmProcessor{},
		ProviderProcessor{},
		ScopesProcessor{},
		ExpiryProcessor{},
		CountryProcessor{},
		IbanProcessor{},
		SwiftProcessor{},
		CardholderProcessor{},
		BrandProcessor{},
	}
}

func ByName(name string) Processor {
	for _, p := range All() {
		if p.Name() == name {
			return p
		}
	}
	return nil
}

func ProcessPayload(in map[string]any) (map[string]any, []string, error) {
	out := make(map[string]any, len(in))
	warnings := make([]string, 0)
	for k, v := range in {
		p := ByName(k)
		if p == nil {
			out[k] = v
			continue
		}
		nv, err := p.Normalize(v)
		if err != nil {
			return nil, nil, err
		}
		if err := p.Validate(nv); err != nil {
			warnings = append(warnings, err.Error())
		}
		out[k] = nv
	}
	return out, warnings, nil
}
