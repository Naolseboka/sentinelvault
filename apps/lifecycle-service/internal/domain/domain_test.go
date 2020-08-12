package domain

import (
	"context"
	"testing"
)

func TestLifecycleInvitedToActive(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	st, err := svc.Invite(ctx, "org-1", "user-1")
	if err != nil || st.State != StateInvited {
		t.Fatal("invite failed")
	}
	st, err = svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-1", To: StateActive})
	if err != nil || st.State != StateActive {
		t.Fatal("activate failed")
	}
}

func TestLifecycleSuspendAndReactivate(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	_, _ = svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-1", To: StateActive})
	st, err := svc.Transition(ctx, TransitionInput{
		OrgID: "org-1", UserID: "user-1", To: StateSuspended, Reason: "leave",
	})
	if err != nil || st.State != StateSuspended {
		t.Fatal("suspend failed")
	}
	_, err = svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-1", To: StateActive})
	if err != nil {
		t.Fatal("reactivate failed")
	}
}

func TestLifecycleOffboard(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	_, _ = svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-2", To: StateActive})
	st, err := svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-2", To: StateOffboarded, Reason: "terminated"})
	if err != nil || st.State != StateOffboarded {
		t.Fatal("offboard failed")
	}
	_, err = svc.Transition(ctx, TransitionInput{OrgID: "org-1", UserID: "user-2", To: StateActive})
	if err == nil {
		t.Fatal("expected invalid transition from offboarded")
	}
}

func TestLifecycleDefaultInvited(t *testing.T) {
	svc := NewService()
	st, err := svc.Get(context.Background(), "org-1", "new-user")
	if err != nil || st.State != StateInvited {
		t.Fatal("expected default invited state")
	}
}
