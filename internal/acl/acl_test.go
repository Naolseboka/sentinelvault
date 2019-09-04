package acl

import (
	"testing"
	"time"
)

func TestACLGrantEffective(t *testing.T) {
	s := NewStore()
	if _, err := s.Grant("o", "vault_item", "i1", "u1", PermEditor, time.Hour); err != nil {
		t.Fatal(err)
	}
	if !s.Can("o", "vault_item", "i1", "u1", PermViewer, nil) {
		t.Fatal("expected viewer allowed via editor")
	}
	if s.Can("o", "vault_item", "i1", "u1", PermOwner, nil) {
		t.Fatal("owner should be denied")
	}
	_, err := s.Grant("o", "collection", "c1", "u2", PermViewer, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !s.Can("o", "vault_item", "i2", "u2", PermViewer, []string{"c1"}) {
		t.Fatal("expected inherited")
	}
}
