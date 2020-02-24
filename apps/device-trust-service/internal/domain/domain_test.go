package domain

import (
	"context"
	"testing"
)

func TestDeviceTrustEvaluateNewDevice(t *testing.T) {
	svc := NewService(NewEngine())
	rec, err := svc.EvaluateTrust(context.Background(), EvaluateTrustInput{
		OrgID: "org-1", DeviceID: "dev-1", UserID: "user-1",
		Fingerprint: "fp-abc", NewDevice: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if rec.TrustLevel != TrustPending {
		t.Fatalf("expected pending, got %q", rec.TrustLevel)
	}
}

func TestDeviceTrustCriticalRiskUntrusted(t *testing.T) {
	svc := NewService(NewEngine())
	rec, err := svc.EvaluateTrust(context.Background(), EvaluateTrustInput{
		OrgID: "org-1", DeviceID: "dev-2", UserID: "user-1", Fingerprint: "fp-bad",
		ImpossibleTravel: true, SuspiciousIP: true, FailedAttempts: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	if rec.TrustLevel != TrustUntrusted {
		t.Fatalf("expected untrusted for critical risk, got %q (score=%d)", rec.TrustLevel, rec.Score)
	}
}

func TestDeviceTrustFingerprintBinding(t *testing.T) {
	svc := NewService(NewEngine())
	ctx := context.Background()
	_, err := svc.EvaluateTrust(ctx, EvaluateTrustInput{
		OrgID: "org-1", DeviceID: "dev-3", UserID: "user-1", Fingerprint: "fp-same",
	})
	if err != nil {
		t.Fatal(err)
	}
	got, err := svc.GetByFingerprint(ctx, "user-1", "fp-same")
	if err != nil || got.DeviceID != "dev-3" {
		t.Fatal("fingerprint binding failed")
	}
}

func TestDeviceTrustTrustedDevice(t *testing.T) {
	svc := NewService(NewEngine())
	rec, err := svc.EvaluateTrust(context.Background(), EvaluateTrustInput{
		OrgID: "org-1", DeviceID: "dev-4", UserID: "user-1",
		Fingerprint: "fp-known", CurrentLevel: TrustTrusted,
	})
	if err != nil {
		t.Fatal(err)
	}
	if rec.TrustLevel != TrustTrusted {
		t.Fatalf("expected trusted, got %q", rec.TrustLevel)
	}
}
