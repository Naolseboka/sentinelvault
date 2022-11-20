package domain

import (
	"context"
	"testing"
)

func TestRiskEvaluateHighScore(t *testing.T) {
	svc := NewService(NewEngine())
	score := svc.Evaluate(context.Background(), EvaluateInput{
		ImpossibleTravel: true, NewDevice: true, SuspiciousIP: true, CredentialStuff: true,
	})
	if score.Level != "critical" || !score.Block {
		t.Fatalf("expected critical block, got level=%q block=%v", score.Level, score.Block)
	}
}

func TestRiskEvaluateLowScore(t *testing.T) {
	svc := NewService(NewEngine())
	score := svc.Evaluate(context.Background(), EvaluateInput{})
	if score.Level != "low" || score.Value != 0 {
		t.Fatalf("expected low zero score, got %d level=%q", score.Value, score.Level)
	}
}
