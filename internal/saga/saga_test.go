package saga

import (
	"context"
	"testing"
)

func TestSagaSuccessAndCompensate(t *testing.T) {
	e := NewEngine()
	Builtins(e)
	inst, err := e.Start(context.Background(), "OrgProvision", map[string]any{"org_name": "Acme"})
	if err != nil {
		t.Fatal(err)
	}
	if inst.Status != StatusCompleted {
		t.Fatalf("%+v", inst)
	}
	inst2, err := e.Start(context.Background(), "EmergencyAccess", map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if inst2.Status != StatusFailed {
		t.Fatalf("expected fail %+v", inst2)
	}
}
