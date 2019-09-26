package outbox

import "testing"

func TestOutbox(t *testing.T) {
	s := NewStore()
	m, err := s.Enqueue("vault.events", "k", map[string]string{"a": "b"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	pending := s.Pending(10)
	if len(pending) != 1 {
		t.Fatal(pending)
	}
	if err := s.MarkPublished(m.ID); err != nil {
		t.Fatal(err)
	}
}
