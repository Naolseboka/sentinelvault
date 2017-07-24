package policy

import "testing"

func TestNewClient(t *testing.T) {
	c := New("http://localhost:8080/", "tok")
	if c.BaseURL != "http://localhost:8080" || c.Token != "tok" {
		t.Fatalf("%+v", c)
	}
}
