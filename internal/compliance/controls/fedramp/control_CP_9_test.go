package fedramp

import (
	"context"
	"testing"
)

func TestControlCP_9Evaluate(t *testing.T) {
	c := ControlCP_9{OrgID: "org-1"}
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
