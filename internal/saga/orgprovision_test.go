package saga

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func TestOrgProvisionSuccess(t *testing.T) {
	e := NewEngine()
	if err := RegisterOrgProvision(e); err != nil {
		t.Fatal(err)
	}
	inst, err := e.Start(context.Background(), "OrgProvision", map[string]any{"org_id": "o1"})
	if err != nil {
		t.Fatal(err)
	}
	if inst.Status != StatusCompleted {
		t.Fatalf("status=%s", inst.Status)
	}
}

func TestOrgProvisionCompensate(t *testing.T) {
	e := NewEngine()
	steps := []StepHandler{&OrgProvisionStep1{}, &OrgProvisionStep2{fail: true}, &OrgProvisionStep3{}}
	_ = RegisterHandlers(e, "OrgProvisionFail", steps...)
	inst, _ := e.Start(context.Background(), "OrgProvisionFail", map[string]any{})
	if inst.Status != StatusFailed {
		t.Fatalf("status=%s", inst.Status)
	}
	_, err := e.Start(context.Background(), "missing", nil)
	if err == nil || !common.IsCode(err, common.ErrCodeNotFound) {
		t.Fatalf("err=%v", err)
	}
}
