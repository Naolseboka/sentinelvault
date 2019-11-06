package saga

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func TestUserOffboardingSuccess(t *testing.T) {
	e := NewEngine()
	if err := RegisterUserOffboarding(e); err != nil {
		t.Fatal(err)
	}
	inst, err := e.Start(context.Background(), "UserOffboarding", map[string]any{"org_id": "o1"})
	if err != nil {
		t.Fatal(err)
	}
	if inst.Status != StatusCompleted {
		t.Fatalf("status=%s", inst.Status)
	}
}

func TestUserOffboardingCompensate(t *testing.T) {
	e := NewEngine()
	steps := []StepHandler{&UserOffboardingStep1{}, &UserOffboardingStep2{fail: true}, &UserOffboardingStep3{}}
	_ = RegisterHandlers(e, "UserOffboardingFail", steps...)
	inst, _ := e.Start(context.Background(), "UserOffboardingFail", map[string]any{})
	if inst.Status != StatusFailed {
		t.Fatalf("status=%s", inst.Status)
	}
	_, err := e.Start(context.Background(), "missing", nil)
	if err == nil || !common.IsCode(err, common.ErrCodeNotFound) {
		t.Fatalf("err=%v", err)
	}
}
