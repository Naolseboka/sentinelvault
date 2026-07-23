package idempotency

import (
	"testing"
	"time"
)

func TestPasskeyIdempotency(t *testing.T) {
	s := NewStore(time.Hour)
	if err := s.Remember("k", "fp", 200, []byte("ok")); err != nil {
		t.Fatal(err)
	}
	rec, ok := s.Lookup("k")
	if !ok || string(rec.Response) != "ok" {
		t.Fatal(rec, ok)
	}
	if err := s.Remember("k", "other", 200, []byte("x")); err == nil {
		t.Fatal("expected conflict")
	}
}

// maintenance note 15
