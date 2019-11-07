package setx

import "testing"

func TestSet(t *testing.T) {
	a := NewString("a", "b")
	b := NewString("b", "c")
	if !a.Has("a") || a.Intersect(b).Len() != 1 {
		t.Fatal("intersect")
	}
	if a.Union(b).Len() != 3 {
		t.Fatal("union")
	}
	if a.Diff(b).Len() != 1 || !a.Diff(b).Has("a") {
		t.Fatal("diff")
	}
}
