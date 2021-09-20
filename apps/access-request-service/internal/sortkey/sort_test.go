package sortkey

import (
	"testing"
	"time"
)

func TestAccessRequestSort(t *testing.T) {
	now := time.Now()
	recs := []Record{{ID: "2", Name: "b", CreatedAt: now}, {ID: "1", Name: "a", CreatedAt: now.Add(-time.Second)}}
	Sort(recs, "name", false)
	if recs[0].Name != "a" {
		t.Fatal(recs)
	}
	if len(AllowedSorts()) == 0 {
		t.Fatal("empty")
	}
}
