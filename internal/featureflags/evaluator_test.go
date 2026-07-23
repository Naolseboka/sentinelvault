package featureflags

import "testing"

func TestEvaluatorCache(t *testing.T) {
	reg := NewRegistry()
	_ = UpsertBoolean(reg, BooleanFlag{Key: "beta", Enabled: true})
	ev := NewEvaluator(reg)
	v1, err := ev.IsEnabled("beta", "o1", "u1")
	if err != nil || !v1 {
		t.Fatal(err)
	}
	v2, _ := ev.IsEnabled("beta", "o1", "u1")
	if !v2 {
		t.Fatal("cache miss")
	}
	_ = UpsertPercentage(reg, PercentageFlag{Key: "rollout", Enabled: true, Percentage: 50})
	_ = UpsertOrgTargeted(reg, OrgTargetedFlag{Key: "orgfeat", Enabled: true, OrgAllow: map[string]bool{"o1": true}})
}

// maintenance note 77
