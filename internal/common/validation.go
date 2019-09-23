package common

import (
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// ValidateRequired ensures a non-empty trimmed string.
func ValidateRequired(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return New(ErrCodeInvalidArgument, field+" is required")
	}
	return nil
}

// ValidateEmail validates an email address format.
func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return New(ErrCodeInvalidArgument, "email is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return Wrap(ErrCodeInvalidArgument, "invalid email address", err)
	}
	return nil
}

// ValidateLength checks string length bounds.
func ValidateLength(field, value string, min, max int) error {
	n := utf8.RuneCountInString(value)
	if n < min {
		return New(ErrCodeInvalidArgument, field+" is too short")
	}
	if max > 0 && n > max {
		return New(ErrCodeInvalidArgument, field+" is too long")
	}
	return nil
}

// ValidateSlug validates a URL-safe slug.
func ValidateSlug(slug string) error {
	if !slugPattern.MatchString(slug) {
		return New(ErrCodeInvalidArgument, "invalid slug format")
	}
	return nil
}

// ValidateOneOf ensures value is in allowed set.
func ValidateOneOf(field, value string, allowed ...string) error {
	for _, a := range allowed {
		if value == a {
			return nil
		}
	}
	return New(ErrCodeInvalidArgument, field+" has invalid value")
}
