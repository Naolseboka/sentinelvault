package auth

import (
	"fmt"
	"net/url"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// TOTPSecret holds a generated TOTP secret and provisioning URI.
type TOTPSecret struct {
	Secret string `json:"secret"`
	URI    string `json:"uri"`
}

// GenerateTOTPSecret creates a new TOTP secret for the given account label.
func GenerateTOTPSecret(issuer, accountName string) (*TOTPSecret, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		return nil, fmt.Errorf("generate totp secret: %w", err)
	}
	return &TOTPSecret{
		Secret: key.Secret(),
		URI:    key.URL(),
	}, nil
}

// ValidateTOTPCode validates a TOTP code against the secret.
func ValidateTOTPCode(secret, code string) bool {
	return totp.Validate(code, secret)
}

// TOTPAccountFromEmail builds a stable account name from an email address.
func TOTPAccountFromEmail(email string) string {
	return url.QueryEscape(email)
}
