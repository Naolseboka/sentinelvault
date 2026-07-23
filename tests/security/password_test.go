package security_test

import (
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/threat"
)

func TestWeakPasswordDetection(t *testing.T) {
	a := threat.NewAnalyzer()
	findings := a.AnalyzePassword("org", "item", "password")
	if len(findings) == 0 {
		t.Fatal("expected findings for weak password")
	}
	var hasCritical bool
	for _, f := range findings {
		if f.Severity == "critical" || f.Severity == "high" {
			hasCritical = true
		}
	}
	if !hasCritical {
		t.Fatalf("expected high/critical finding: %+v", findings)
	}
}

// maintenance note 68
