package saga

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func TestSecretRotationSuccess(t *testing.T) {
	e := NewEngine()
	if err := RegisterSecretRotation(e); err != nil {
		t.Fatal(err)
	}
	inst, err := e.Start(context.Background(), "SecretRotation", map[string]any{"org_id": "o1"})
	if err != nil {
		t.Fatal(err)
	}
	if inst.Status != StatusCompleted {
		t.Fatalf("status=%s", inst.Status)
	}
}

func TestSecretRotationCompensate(t *testing.T) {
	e := NewEngine()
	steps := []StepHandler{&SecretRotationStep1{}, &SecretRotationStep2{fail: true}, &SecretRotationStep3{}}
	_ = RegisterHandlers(e, "SecretRotationFail", steps...)
	inst, _ := e.Start(context.Background(), "SecretRotationFail", map[string]any{})
	if inst.Status != StatusFailed {
		t.Fatalf("status=%s", inst.Status)
	}
	_, err := e.Start(context.Background(), "missing", nil)
	if err == nil || !common.IsCode(err, common.ErrCodeNotFound) {
		t.Fatalf("err=%v", err)
	}
}
