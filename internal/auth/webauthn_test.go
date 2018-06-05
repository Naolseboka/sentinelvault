package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
)

func TestWebAuthnChallengeStore(t *testing.T) {
	t.Parallel()
	store := auth.NewMemoryChallengeStore()
	ctx := context.Background()
	ch, err := auth.NewWebAuthnChallenge("user-1", "enroll", time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(ctx, ch); err != nil {
		t.Fatal(err)
	}
	got, err := store.Consume(ctx, ch.ID, "user-1")
	if err != nil {
		t.Fatal(err)
	}
	if string(got.Challenge) != string(ch.Challenge) {
		t.Fatal("challenge mismatch")
	}
	if _, err := store.Consume(ctx, ch.ID, "user-1"); err == nil {
		t.Fatal("expected challenge not found on second consume")
	}
}

func TestVerifyAssertionBinding(t *testing.T) {
	t.Parallel()
	a := []byte("challenge-a")
	b := []byte("challenge-a")
	if err := auth.VerifyAssertionBinding(a, b); err != nil {
		t.Fatal(err)
	}
	if err := auth.VerifyAssertionBinding(a, []byte("other")); err == nil {
		t.Fatal("expected mismatch")
	}
}

func TestEncodeDecodeChallenge(t *testing.T) {
	t.Parallel()
	raw := []byte("secret-challenge-bytes")
	enc := auth.EncodeChallenge(raw)
	dec, err := auth.DecodeChallenge(enc)
	if err != nil {
		t.Fatal(err)
	}
	if string(dec) != string(raw) {
		t.Fatal("roundtrip failed")
	}
}

func TestMemoryCredentialStore(t *testing.T) {
	t.Parallel()
	store := auth.NewMemoryCredentialStore()
	ctx := context.Background()
	cred := &auth.WebAuthnCredential{
		ID: "cred-1", UserID: "user-1",
		CredentialID: []byte("id-bytes"), PublicKey: []byte("pk"),
		SignCount: 0, CreatedAt: time.Now().UTC(),
	}
	if err := store.Save(ctx, cred); err != nil {
		t.Fatal(err)
	}
	list, err := store.ListByUser(ctx, "user-1")
	if err != nil || len(list) != 1 {
		t.Fatal("expected one credential")
	}
	got, err := store.GetByCredentialID(ctx, "user-1", []byte("id-bytes"))
	if err != nil || got.ID != "cred-1" {
		t.Fatal("get by credential id failed")
	}
	if err := store.UpdateSignCount(ctx, "cred-1", 3); err != nil {
		t.Fatal(err)
	}
	got, _ = store.GetByCredentialID(ctx, "user-1", []byte("id-bytes"))
	if got.SignCount != 3 {
		t.Fatal("sign count not updated")
	}
}

func TestAuthEventHelpers(t *testing.T) {
	t.Parallel()
	ev := auth.NewAuthEvent(auth.EventLoginSucceeded).
		WithUser("u").WithOrg("o").WithSession("s").WithEmail("e@x.com").WithReason("ok")
	if ev.UserID != "u" || ev.OrgID != "o" || ev.SessionID != "s" || ev.Email != "e@x.com" {
		t.Fatalf("event=%+v", ev)
	}
}
