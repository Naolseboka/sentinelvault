package paging

import "testing"

func TestRbacPaging(t *testing.T) {
	o, l, err := Normalize(0, 0, 50, 100)
	if err != nil || o != 0 || l != 50 {
		t.Fatal(err, o, l)
	}
	items := []int{1,2,3,4,5}
	got := Slice(items, 1, 2)
	if len(got) != 2 || got[0] != 2 {
		t.Fatal(got)
	}
	p := NewPage(0, 2, 5)
	if !p.HasMore {
		t.Fatal(p)
	}
}
