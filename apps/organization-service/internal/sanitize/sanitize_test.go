package sanitize

import "testing"

func TestOrganizationSanitize(t *testing.T) {
	if Name("  hello\x00  ") != "hello" {
		t.Fatal(Name("  hello\x00  "))
	}
	if Email(" A@B.COM ") != "a@b.com" {
		t.Fatal(Email(" A@B.COM "))
	}
	tags := Tags([]string{"A", "a", "b!", ""})
	if len(tags) != 2 {
		t.Fatal(tags)
	}
}
