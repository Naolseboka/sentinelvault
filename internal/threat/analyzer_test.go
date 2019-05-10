package threat

import "testing"

func TestAnalyzePassword(t *testing.T) {
	a := NewAnalyzer()
	findings := a.AnalyzePassword("org", "item1", "password")
	if len(findings) < 2 {
		t.Fatalf("expected multiple findings got %d", len(findings))
	}
	strong := a.AnalyzePassword("org", "item2", "Str0ng!VaultPass#2026")
	if len(strong) != 0 {
		t.Fatalf("unexpected findings %+v", strong)
	}
}

func TestReuseAndExposure(t *testing.T) {
	a := NewAnalyzer()
	reuse := a.DetectReuse("org", map[string]string{"a": "same", "b": "same", "c": "other"})
	if len(reuse) != 1 {
		t.Fatalf("reuse=%d", len(reuse))
	}
	db := MemoryDarkWeb{Exposed: map[string]string{"leaked": "collection-x"}}
	exp, err := a.CheckExposure("org", "r1", "leaked", db)
	if err != nil || len(exp) != 1 {
		t.Fatalf("exp=%v err=%v", exp, err)
	}
}
