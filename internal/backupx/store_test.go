package backupx

import "testing"

func TestBackupRestore(t *testing.T) {
	store, err := NewStore(nil)
	if err != nil {
		t.Fatal(err)
	}
	payload := map[string]any{"secrets": []string{"a", "b"}, "version": 1}
	snap, err := store.Backup("org1", "nightly", payload)
	if err != nil {
		t.Fatal(err)
	}
	if !snap.Encrypted || snap.SizeBytes == 0 {
		t.Fatalf("snap=%+v", snap)
	}
	var dest map[string]any
	got, err := store.Restore(snap.ID, &dest)
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != snap.ID {
		t.Fatal("id mismatch")
	}
	if dest["version"].(float64) != 1 {
		t.Fatalf("dest=%v", dest)
	}
	if len(store.List("org1")) != 1 {
		t.Fatal("list")
	}
}
