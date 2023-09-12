package domain

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/notify"
)

func TestNotificationSendAndList(t *testing.T) {
	svc := NewService(NewDispatcher())
	ctx := context.Background()

	msg, err := svc.Send(ctx, SendInput{
		OrgID: "org-1", Channel: notify.ChannelEmail,
		Recipient: "user@example.com", Subject: "Alert", Body: "test", Priority: "high",
	})
	if err != nil {
		t.Fatal(err)
	}
	if msg.Status != "sent" {
		t.Fatalf("expected sent, got %q", msg.Status)
	}

	items := svc.List(ctx, "org-1")
	if len(items) != 1 {
		t.Fatalf("expected 1 message, got %d", len(items))
	}
}

func TestNotificationUnknownChannel(t *testing.T) {
	svc := NewService(NewDispatcher())
	_, err := svc.Send(context.Background(), SendInput{
		OrgID: "org-1", Channel: "carrier_pigeon", Recipient: "x",
	})
	if err == nil {
		t.Fatal("expected channel not found error")
	}
}
