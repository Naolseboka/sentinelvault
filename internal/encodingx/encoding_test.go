package encodingx

import "testing"

func TestEncoding(t *testing.T) {
	in := []byte("hello")
	h := HexEncode(in)
	b, err := HexDecode(h)
	if err != nil || string(b) != "hello" {
		t.Fatal(err, b)
	}
	u := Base64URLEncode(in)
	b, err = Base64URLDecode(u)
	if err != nil || string(b) != "hello" {
		t.Fatal(err, b)
	}
	s := Base64StdEncode(in)
	b, err = Base64StdDecode(s)
	if err != nil || string(b) != "hello" {
		t.Fatal(err, b)
	}
}
