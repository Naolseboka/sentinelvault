package oidc

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"
)

func TestDiscoveryPKCEAuthURL(t *testing.T) {
	doc, err := ParseDiscovery([]byte(`{"issuer":"https://idp.example","authorization_endpoint":"https://idp.example/auth","token_endpoint":"https://idp.example/token"}`))
	if err != nil {
		t.Fatal(err)
	}
	pkce, err := NewPKCES256()
	if err != nil {
		t.Fatal(err)
	}
	if !VerifyPKCE(pkce.Verifier, pkce.Challenge, "S256") {
		t.Fatal("pkce verify")
	}
	u, err := BuildAuthCodeURL(AuthCodeURLInput{
		AuthorizationEndpoint: doc.AuthorizationEndpoint,
		ClientID: "client", RedirectURI: "https://app/callback",
		Scope: []string{"openid", "email"}, State: "s", Nonce: "n",
		PKCEChallenge: pkce.Challenge, PKCEMethod: "S256",
	})
	if err != nil || u == "" {
		t.Fatal(err)
	}
	claims := IDTokenClaims{Iss: doc.Issuer, Sub: "u1", Aud: "client", Exp: time.Now().Add(time.Hour).Unix(), Iat: time.Now().Unix(), Nonce: "n"}
	if err := ValidateClaims(claims, doc.Issuer, []string{"client"}, "n", time.Now(), 0); err != nil {
		t.Fatal(err)
	}
	payload, _ := json.Marshal(claims)
	jwt := "hdr." + base64.RawURLEncoding.EncodeToString(payload) + ".sig"
	decoded, err := DecodeClaimsPayload(jwt)
	if err != nil || decoded.Sub != "u1" {
		t.Fatal(err)
	}
	tr := TokenRequest{TokenEndpoint: doc.TokenEndpoint, GrantType: "authorization_code", Code: "c", RedirectURI: "https://app/callback", ClientID: "client", CodeVerifier: pkce.Verifier}
	if _, err := tr.FormValues(); err != nil {
		t.Fatal(err)
	}
}
