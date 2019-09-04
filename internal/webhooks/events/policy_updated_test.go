package webhooks

import "testing"

func TestPolicyUpdatedPayload(t *testing.T) {
	p, err := NewPolicyUpdatedPayload("o1", "u1", "r1", map[string]string{"k": "v"})
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Validate(); err != nil {
		t.Fatal(err)
	}
	b, err := p.JSON()
	if err != nil || len(b) == 0 {
		t.Fatal(err)
	}
}
