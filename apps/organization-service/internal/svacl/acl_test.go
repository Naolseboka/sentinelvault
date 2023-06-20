package svacl

import (
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/acl"
)

func TestOrganizationGuard(t *testing.T) {
	g := NewGuard(acl.NewStore())
	if err := g.Authorize("o", "u", "viewer", ActionRead, "r1", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.Authorize("o", "u", "viewer", ActionDelete, "r1", nil); err == nil {
		t.Fatal("expected forbid")
	}
	if err := g.Authorize("o", "u", "admin", ActionConfigure, "r1", nil); err != nil {
		t.Fatal(err)
	}
	if len(AllActions()) == 0 {
		t.Fatal("actions")
	}
}
// maintenance note 36
