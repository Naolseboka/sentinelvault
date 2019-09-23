package common

import (
	"errors"
	"testing"
)

func TestAppError(t *testing.T) {
	cause := errors.New("db down")
	err := Wrap(ErrCodeUnavailable, "service unavailable", cause)
	if err.Error() == "" {
		t.Fatal("expected non-empty error")
	}
	if !errors.Is(err, cause) {
		t.Fatal("expected unwrap to work")
	}
	if !IsCode(err, ErrCodeUnavailable) {
		t.Fatal("expected unavailable code")
	}
	det := err.WithDetails(map[string]any{"retry": true})
	if det.Details["retry"] != true {
		t.Fatal("expected details")
	}
}

func TestAsAppError(t *testing.T) {
	raw := errors.New("boom")
	ae := AsAppError(raw)
	if ae.Code != ErrCodeInternal {
		t.Fatalf("got %s", ae.Code)
	}
	orig := New(ErrCodeNotFound, "missing")
	if AsAppError(orig) != orig {
		t.Fatal("expected same pointer")
	}
}
