package notify

import (
	"context"
	"testing"
)

func TestDispatcher(t *testing.T) {
	d := NewDispatcher()
	RegisterDefaults(d)
	msg, err := d.Send(context.Background(), "org1", ChannelSlack, "#sec", "Alert", "rotation failed", "high", nil)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Status != "sent" {
		t.Fatalf("status=%s", msg.Status)
	}
	if len(d.Sent()) != 1 {
		t.Fatal("expected history")
	}
	if _, err := d.Send(context.Background(), "org1", "fax", "x", "", "", "", nil); err == nil {
		t.Fatal("expected missing channel error")
	}
}
