package featureflags

import "testing"

func TestFeatureFlags(t *testing.T) {
	r := NewRegistry()
	if err := r.Upsert(Flag{Key: "new_ui", Enabled: true, Percentage: 100}); err != nil {
		t.Fatal(err)
	}
	on, err := r.Evaluate("new_ui", "o1", "u1")
	if err != nil || !on {
		t.Fatal(on, err)
	}
	_ = r.Upsert(Flag{Key: "beta", Enabled: true, Percentage: 0, OrgAllow: map[string]bool{"o2": true}})
	on, _ = r.Evaluate("beta", "o2", "u")
	if !on {
		t.Fatal("allow list")
	}
	on, _ = r.Evaluate("beta", "o3", "u")
	if on {
		t.Fatal("expected off")
	}
}
