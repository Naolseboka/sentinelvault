package localcache

import (
	"testing"
	"time"
)

func TestBackupCache(t *testing.T) {
	c := New()
	c.Set("a", 1, time.Minute)
	v, ok := c.Get("a")
	if !ok || v.(int) != 1 {
		t.Fatal(v, ok)
	}
	got, err := c.GetOrSet("b", time.Minute, func() (any, error) { return "x", nil })
	if err != nil || got != "x" {
		t.Fatal(err, got)
	}
}
