package textutil

import "testing"

func TestTextUtil(t *testing.T) {
	if CollapseWS(" a  b ") != "a b" {
		t.Fatal(CollapseWS(" a  b "))
	}
	if Slugify("Hello World!") != "hello-world" {
		t.Fatal(Slugify("Hello World!"))
	}
	if !ContainsFold("AbC", "b") {
		t.Fatal("contains")
	}
	if Ellipsize("abcdef", 5) != "ab..." {
		t.Fatal(Ellipsize("abcdef", 5))
	}
	if len(UniqueStrings([]string{"a", "a", "b"})) != 2 {
		t.Fatal("unique")
	}
}
