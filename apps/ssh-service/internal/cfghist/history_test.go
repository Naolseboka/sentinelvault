package cfghist

import "testing"

func TestSshCfgHist(t *testing.T) {
	s := New()
	snap, err := s.Commit("admin", "init", map[string]any{"enabled": true})
	if err != nil || snap.Version != 1 {
		t.Fatal(err, snap)
	}
	got, ok := s.Latest()
	if !ok || got.ID != snap.ID {
		t.Fatal(got, ok)
	}
}
