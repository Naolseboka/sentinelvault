package durationx

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	d, err := Parse("7d")
	if err != nil || d != 7*24*time.Hour {
		t.Fatal(err, d)
	}
	if Humanize(2*time.Hour) != "2h" {
		t.Fatal(Humanize(2 * time.Hour))
	}
}
