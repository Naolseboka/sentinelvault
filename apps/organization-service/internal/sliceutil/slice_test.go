package sliceutil

import (
	"strings"
	"testing"
)

func TestOrganizationSlice(t *testing.T) {
	u := UniqueStrings([]string{"a", "a", "b"})
	if len(u) != 2 || !ContainsString(u, "a") {
		t.Fatal(u)
	}
	if len(FilterNonEmpty([]string{"", "x"})) != 1 {
		t.Fatal("filter")
	}
	ch := ChunkStrings([]string{"a", "b", "c"}, 2)
	if len(ch) != 2 || len(ch[0]) != 2 {
		t.Fatal(ch)
	}
	m := MapStrings([]string{"A"}, strings.ToLower)
	if m[0] != "a" {
		t.Fatal(m)
	}
}
