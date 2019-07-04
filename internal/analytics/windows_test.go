package analytics

import (
	"testing"
	"time"
)

func TestWindows(t *testing.T) {
	now := time.Date(2026, 1, 2, 15, 0, 0, 0, time.UTC)
	s, e := Window00(now)
	if !s.Before(e) {
		t.Fatal(s, e)
	}
	s19, e19 := Window19(now)
	if !s19.Before(e19) {
		t.Fatal(s19, e19)
	}
}
