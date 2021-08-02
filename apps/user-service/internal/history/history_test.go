package history

import "testing"

func TestUserHistory(t *testing.T) {
	h := NewUserHistory(100)
	h.RecordNameChange("o", "a", "r", "old", "new")
	h.RecordStatusChange("o", "a", "r", "old", "new")
	h.RecordOwnerChange("o", "a", "r", "old", "new")
	h.RecordTagsChange("o", "a", "r", "old", "new")
	h.RecordMetadataChange("o", "a", "r", "old", "new")
	h.RecordPayloadChange("o", "a", "r", "old", "new")
	h.RecordPermissionChange("o", "a", "r", "old", "new")
	h.RecordTtlChange("o", "a", "r", "old", "new")
	if h.Count("o") < 8 {
		t.Fatalf("count=%d", h.Count("o"))
	}
	if len(h.List("o", "r", 10)) == 0 {
		t.Fatal("list")
	}
}
