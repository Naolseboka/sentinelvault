package acl

import (
	"testing"
	"time"
)

func TestInheritanceResolver(t *testing.T) {
	st := NewStore()
	r := NewInheritanceResolver(st)
	_, _ = r.GrantCollection("o1", "col1", "u1", PermViewer, time.Hour)
	r.LinkItem("item1", "col1")
	perm, ok := r.EffectiveForItem("o1", "item1", "u1")
	if !ok || perm != PermViewer {
		t.Fatalf("perm=%s ok=%v", perm, ok)
	}
	_, _ = st.Grant("o1", "item", "item1", "u2", PermEditor, 0)
	entries := st.ListEntries("o1", "item", "item1")
	if len(entries) != 1 {
		t.Fatalf("len=%d", len(entries))
	}
	if st.RevokeByPrincipal("o1", "item", "item1", "u2") != 1 {
		t.Fatal("revoke count")
	}
}
