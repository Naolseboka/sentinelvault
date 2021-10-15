package gqltypes

import (
	"testing"
	"time"
)

func TestApprovalConnect(t *testing.T) {
	items := []Approval{{ID: "1", Name: "a", CreatedAt: time.Now()}, {ID: "2", Name: "b", CreatedAt: time.Now()}}
	c := Connect(items, 0, 1, 2)
	if len(c.Nodes) != 1 || !c.PageInfo.HasNextPage {
		t.Fatalf("%+v", c)
	}
}
