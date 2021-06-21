package batch

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func TestUserBatch(t *testing.T) {
	p := NewProcessor()
	if err := ValidateBatchSize(0, 10); err == nil {
		t.Fatal("empty")
	}
	res := p.CreateMany(context.Background(), []Item{{OrgID: "o", Name: "a"}, {Name: "bad"}}, func(ctx context.Context, item Item) (string, error) {
		return common.NewID(), nil
	})
	if len(res) != 2 || !res[0].OK || res[1].OK {
		t.Fatalf("%+v", res)
	}
	c, f := p.Stats()
	if c != 1 || f != 1 {
		t.Fatalf("%d %d", c, f)
	}
}
