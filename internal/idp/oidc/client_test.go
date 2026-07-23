package oidc

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

type mockDoer struct {
	body string
	code int
}

func (m mockDoer) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: m.code,
		Body:       io.NopCloser(bytes.NewBufferString(m.body)),
		Header:     make(http.Header),
	}, nil
}

func TestClientExchange(t *testing.T) {
	c := &Client{HTTP: mockDoer{code: 200, body: `{"access_token":"at","id_token":"x.y.z","token_type":"Bearer","expires_in":3600}`}}
	out, err := c.ExchangeAuthorizationCode(context.Background(), TokenRequest{
		TokenEndpoint: "https://idp/token", GrantType: "authorization_code",
		Code: "code", RedirectURI: "https://app/cb", ClientID: "cid", CodeVerifier: "v",
	})
	if err != nil || out.AccessToken != "at" {
		t.Fatal(err)
	}
}

func TestClientExchangeError(t *testing.T) {
	c := &Client{HTTP: mockDoer{code: 401, body: `{"error":"invalid_grant"}`}}
	_, err := c.ExchangeAuthorizationCode(context.Background(), TokenRequest{
		TokenEndpoint: "https://idp/token", GrantType: "authorization_code",
		Code: "bad", RedirectURI: "https://app/cb", ClientID: "cid",
	})
	if err == nil {
		t.Fatal("expected error")
	}
}
