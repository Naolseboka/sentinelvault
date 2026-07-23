package reqctx

import "context"

type key int

const (
	keyRequestID key = iota + 1
	keyOrgID
	keyActorID
	keySessionID
	keyTraceID
)

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyRequestID, id)
}
func RequestID(ctx context.Context) string {
	v, _ := ctx.Value(keyRequestID).(string)
	return v
}
func WithOrgID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyOrgID, id)
}
func OrgID(ctx context.Context) string {
	v, _ := ctx.Value(keyOrgID).(string)
	return v
}
func WithActorID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyActorID, id)
}
func ActorID(ctx context.Context) string {
	v, _ := ctx.Value(keyActorID).(string)
	return v
}
func WithSessionID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keySessionID, id)
}
func SessionID(ctx context.Context) string {
	v, _ := ctx.Value(keySessionID).(string)
	return v
}
func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyTraceID, id)
}
func TraceID(ctx context.Context) string {
	v, _ := ctx.Value(keyTraceID).(string)
	return v
}

// maintenance note 88
