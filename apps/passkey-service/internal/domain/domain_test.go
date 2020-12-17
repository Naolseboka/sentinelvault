package domain

import (
	"context"
	"testing"
)

func TestPasskeyRegisterFlow(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	ch, err := svc.CreateChallenge(ctx, CreateChallengeInput{UserID: "user-1", OrgID: "org-1", Kind: "enroll"})
	if err != nil {
		t.Fatal(err)
	}
	rec, err := svc.Register(ctx, RegisterInput{
		UserID: "user-1", OrgID: "org-1", ChallengeID: ch.ChallengeID, ClientChallenge: ch.Challenge,
		Name: "MacBook", CredentialID: "Y3JlZA", PublicKey: "cHVi",
	})
	if err != nil {
		t.Fatal(err)
	}
	if rec.Name != "MacBook" {
		t.Fatal("credential not stored")
	}
	list, err := svc.List(ctx, "user-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 passkey, got %d", len(list))
	}
}

func TestPasskeyAssertFlow(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	enrollCh, err := svc.CreateChallenge(ctx, CreateChallengeInput{UserID: "user-1", Kind: "enroll"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Register(ctx, RegisterInput{
		UserID: "user-1", ChallengeID: enrollCh.ChallengeID, ClientChallenge: enrollCh.Challenge,
		CredentialID: "Y3JlZA", PublicKey: "cHVi", Name: "Phone",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertCh, err := svc.CreateChallenge(ctx, CreateChallengeInput{UserID: "user-1", Kind: "assert"})
	if err != nil {
		t.Fatal(err)
	}
	result, err := svc.Assert(ctx, AssertInput{
		UserID: "user-1", ChallengeID: assertCh.ChallengeID, ClientChallenge: assertCh.Challenge,
		CredentialID: "Y3JlZA", SignCount: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Verified || result.SignCount != 1 {
		t.Fatalf("assert failed: %+v", result)
	}
}

func TestPasskeyAssertRequiresCredential(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	_, err := svc.CreateChallenge(ctx, CreateChallengeInput{UserID: "user-2", Kind: "assert"})
	if err == nil {
		t.Fatal("expected error when no credentials registered")
	}
}
