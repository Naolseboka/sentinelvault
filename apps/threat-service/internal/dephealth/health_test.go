package dephealth

import (
	"context"
	"errors"
	"testing"
)

func TestThreatDepHealth(t *testing.T) {
	r := New()
	r.Register("ok", func(ctx context.Context) error { return nil })
	r.Register("bad", func(ctx context.Context) error { return errors.New("down") })
	res := r.Run(context.Background())
	if len(res) != 2 {
		t.Fatal(res)
	}
	if r.Ready(context.Background()) {
		t.Fatal("expected not ready")
	}
}
