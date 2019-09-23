package common

import (
	"errors"
	"fmt"
)

// ErrorCode classifies application errors for API responses and logging.
type ErrorCode string

const (
	ErrCodeInvalidArgument ErrorCode = "invalid_argument"
	ErrCodeNotFound        ErrorCode = "not_found"
	ErrCodeConflict        ErrorCode = "conflict"
	ErrCodeUnauthorized    ErrorCode = "unauthorized"
	ErrCodeForbidden       ErrorCode = "forbidden"
	ErrCodeInternal        ErrorCode = "internal"
	ErrCodeUnavailable     ErrorCode = "unavailable"
	ErrCodeNotImplemented  ErrorCode = "not_implemented"
)

// AppError is a typed error with machine-readable code and optional details.
type AppError struct {
	Code    ErrorCode
	Message string
	Details map[string]any
	Cause   error
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error { return e.Cause }

// New creates an AppError with the given code and message.
func New(code ErrorCode, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Wrap attaches a cause to an AppError.
func Wrap(code ErrorCode, message string, cause error) *AppError {
	return &AppError{Code: code, Message: message, Cause: cause}
}

// WithDetails returns a copy with additional details.
func (e *AppError) WithDetails(details map[string]any) *AppError {
	cp := *e
	if e.Details != nil {
		cp.Details = make(map[string]any, len(e.Details)+len(details))
		for k, v := range e.Details {
			cp.Details[k] = v
		}
	} else {
		cp.Details = make(map[string]any, len(details))
	}
	for k, v := range details {
		cp.Details[k] = v
	}
	return &cp
}

// AsAppError extracts an AppError from err, optionally wrapping unknown errors.
func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	var ae *AppError
	if errors.As(err, &ae) {
		return ae
	}
	return Wrap(ErrCodeInternal, "unexpected error", err)
}

// IsCode reports whether err is an AppError with the given code.
func IsCode(err error, code ErrorCode) bool {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae.Code == code
	}
	return false
}
