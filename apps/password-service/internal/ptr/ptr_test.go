package ptr

import "testing"

func TestPasswordPtr(t *testing.T) {
	s := String("x")
	if StringVal(s) != "x" || StringVal(nil) != "" {
		t.Fatal("string")
	}
	if BoolVal(Bool(true)) != true {
		t.Fatal("bool")
	}
}
