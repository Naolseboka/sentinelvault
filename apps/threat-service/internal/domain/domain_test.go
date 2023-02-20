package domain

import (
	"context"
	"testing"
)

func TestThreatAnalyzeWeakPassword(t *testing.T) {
	svc := NewService(NewAnalyzer())
	findings := svc.AnalyzePassword(context.Background(), AnalyzePasswordInput{
		OrgID: "org-1", ResourceID: "pwd-1", Password: "123456",
	})
	if len(findings) == 0 {
		t.Fatal("expected findings for weak password")
	}
}

func TestThreatDetectReuse(t *testing.T) {
	svc := NewService(NewAnalyzer())
	findings := svc.DetectReuse(context.Background(), DetectReuseInput{
		OrgID: "org-1", ItemPasswords: map[string]string{"a": "same", "b": "same"},
	})
	if len(findings) != 1 || findings[0].Type != "password_reuse" {
		t.Fatal("expected password reuse finding")
	}
}
