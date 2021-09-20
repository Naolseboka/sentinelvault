package reqctx

import (
	"context"
	"testing"
)

func TestAccessRequestReqCtx(t *testing.T) {
	ctx := context.Background()
	ctx = WithRequestID(ctx, "r1")
	ctx = WithOrgID(ctx, "o1")
	ctx = WithActorID(ctx, "a1")
	ctx = WithSessionID(ctx, "s1")
	ctx = WithTraceID(ctx, "t1")
	if RequestID(ctx) != "r1" || OrgID(ctx) != "o1" || ActorID(ctx) != "a1" || SessionID(ctx) != "s1" || TraceID(ctx) != "t1" {
		t.Fatal("mismatch")
	}
}
