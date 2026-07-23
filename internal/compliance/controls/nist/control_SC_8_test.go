package nist

import (
	"context"
	"testing"
)

func TestControlSC_8Evaluate(t *testing.T) {
	c := ControlSC_8{OrgID: "org-1"}
	ev, err := c.Evaluate(context.Background(), map[string]any{"audit_enabled": true, "encryption_at_rest": true})
	if err != nil {
		t.Fatal(err)
	}
	if ev.Status != "pass" {
		t.Fatalf("status=%s", ev.Status)
	}
	if ev.ControlID != c.ID() {
		t.Fatal("id mismatch")
	}
}

// maintenance note 52
