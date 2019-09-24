package saml

import (
    "testing"
    "time"
)

func TestAuthnRequestAndAssertion(t *testing.T) {
    xml, err := BuildAuthnRequestXML(AuthnRequest{
        Destination: "https://idp/saml/sso", Issuer: "sp", ACSURL: "https://sp/acs",
    })
    if err != nil || len(xml) == 0 {
        t.Fatal(err)
    }
    enc := EncodeRedirectRequest(xml)
    if enc == "" {
        t.Fatal("empty encode")
    }
    sample := []byte(`<Response><Assertion><Subject><NameID>user@example.com</NameID></Subject><Conditions NotBefore="2020-01-01T00:00:00Z" NotOnOrAfter="2030-01-01T00:00:00Z"><Audience>sp</Audience></Conditions><AttributeStatement><Attribute Name="email">user@example.com</Attribute></AttributeStatement></Assertion></Response>`)
    a, err := ParseResponseAttributes(sample)
    if err != nil || a.Subject == "" {
        t.Fatal(err)
    }
    if err := ValidateAssertion(a, "sp", "", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), time.Minute); err != nil {
        t.Fatal(err)
    }
}

func TestValidateAssertionErrors(t *testing.T) {
    a := &ResponseAssertion{Subject: "u", Audience: "wrong", NotOnOrAfter: time.Now().Add(-time.Hour)}
    if err := ValidateAssertion(a, "expected", "", time.Now(), 0); err == nil {
        t.Fatal("expected audience error")
    }
}