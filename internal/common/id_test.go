package common

import "testing"

func TestNewID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()
	if id1 == id2 {
		t.Fatal("expected unique ids")
	}
	if len(id1) != 26 {
		t.Fatalf("expected ulid length 26, got %d", len(id1))
	}
	parsed, err := ParseID(id1)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.String() != id1 {
		t.Fatal("round trip failed")
	}
}

func TestParseIDInvalid(t *testing.T) {
	if _, err := ParseID("not-a-ulid"); err == nil {
		t.Fatal("expected error")
	}
}
