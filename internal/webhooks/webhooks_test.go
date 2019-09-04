package webhooks

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebhookSignedDelivery(t *testing.T) {
	var gotSig string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotSig = r.Header.Get("X-SV-Signature")
		b, _ := io.ReadAll(r.Body)
		_ = b
		w.WriteHeader(204)
	}))
	defer srv.Close()
	reg := NewRegistry()
	ep, err := reg.Register("o1", srv.URL, []string{"vault.item.created"})
	if err != nil {
		t.Fatal(err)
	}
	body := []byte(`{"id":"1"}`)
	atts := reg.Deliver(context.Background(), "o1", "vault.item.created", body)
	if len(atts) != 1 || !atts[0].Success {
		t.Fatalf("%+v", atts)
	}
	if gotSig != Sign(ep.Secret, body) {
		t.Fatal("signature mismatch")
	}
	if Backoff(1) <= 0 {
		t.Fatal("backoff")
	}
}
