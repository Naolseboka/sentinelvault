package semver

import "testing"

func TestSemver(t *testing.T) {
	v, err := Parse("1.2.3-alpha+001")
	if err != nil || v.String() == "" {
		t.Fatal(err, v)
	}
	a, _ := Parse("1.2.3")
	b, _ := Parse("1.2.4")
	if Compare(a, b) >= 0 {
		t.Fatal("compare")
	}
	if a.BumpPatch().Patch != 4 {
		t.Fatal(a.BumpPatch())
	}
}
