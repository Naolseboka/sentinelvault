package webhooks

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeliverWithRetry(t *testing.T) {
	var hits int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		if hits < 2 {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	reg := NewRegistry()
	reg.client = srv.Client()
	ep, _ := reg.Register("o1", srv.URL, []string{"test.event"})
	_ = ep
	attempts := reg.DeliverWithRetry(context.Background(), "o1", "test.event", []byte(`{}`), 3)
	if len(attempts) == 0 {
		t.Fatal("no attempts")
	}
}
