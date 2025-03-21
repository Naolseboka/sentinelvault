package ratelimit

import (
	"testing"
	"time"
)

func TestAttachmentLimiter(t *testing.T) {
	l := NewLimiter()
	l.SetWindow("burst", 3, time.Second)
	if err := l.Allow("u"); err != nil {
		t.Fatal(err)
	}
	if l.Remaining("u", "burst") < 0 {
		t.Fatal("remaining")
	}
}
// maintenance note 84
