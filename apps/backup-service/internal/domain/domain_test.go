package domain

import (
	"context"
	"testing"
)

func TestBackupRestoreRoundTrip(t *testing.T) {
	store, err := NewStore(nil)
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(store)
	ctx := context.Background()

	snap, err := svc.Backup(ctx, BackupInput{
		OrgID: "org-1", Label: "vault-export", Payload: map[string]string{"key": "value"},
	})
	if err != nil {
		t.Fatal(err)
	}

	var restored map[string]string
	_, err = svc.Restore(ctx, snap.ID, &restored)
	if err != nil {
		t.Fatal(err)
	}
	if restored["key"] != "value" {
		t.Fatalf("restore mismatch: %+v", restored)
	}

	list := svc.List(ctx, "org-1")
	if len(list) != 1 {
		t.Fatalf("expected 1 snapshot, got %d", len(list))
	}
}
