package domain

import (
	"context"
	"testing"
)

func TestAuditAppendListVerify(t *testing.T) {
	chain, err := NewChain()
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(chain)
	ctx := context.Background()

	ev, err := svc.Append(ctx, AppendInput{
		OrgID: "org-1", ActorID: "user-1", Action: "login",
		ResourceType: "session", ResourceID: "sess-1", Severity: "info",
	})
	if err != nil {
		t.Fatal(err)
	}
	if ev.Hash == "" || ev.Signature == "" {
		t.Fatal("expected hash and signature")
	}

	items, total, err := svc.List(ctx, "org-1", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 || len(items) != 1 {
		t.Fatalf("expected 1 event, got %d total %d", len(items), total)
	}
	if err := svc.VerifyIntegrity(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestAuditAppendRequiresOrgAndAction(t *testing.T) {
	chain, _ := NewChain()
	svc := NewService(chain)
	_, err := svc.Append(context.Background(), AppendInput{OrgID: "", Action: ""})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
