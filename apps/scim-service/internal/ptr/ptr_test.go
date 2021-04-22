package ptr

import "testing"

func TestScimPtr(t *testing.T) {
	s := String("x")
	if StringVal(s) != "x" || StringVal(nil) != "" {
		t.Fatal("string")
	}
	if BoolVal(Bool(true)) != true {
		t.Fatal("bool")
	}
}
