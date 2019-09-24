package saml

import (
    "encoding/base64"
    "encoding/xml"
    "fmt"
    "strings"
    "time"

    "github.com/sentinelvault/sentinelvault/internal/common"
)

// AuthnRequest is a simplified SAML 2.0 AuthnRequest builder.
type AuthnRequest struct {
    ID           string
    IssueInstant time.Time
    Destination  string
    Issuer       string
    ACSURL       string
    ForceAuthn   bool
}

// BuildAuthnRequestXML returns deflated-ready XML (not compressed here).
func BuildAuthnRequestXML(r AuthnRequest) ([]byte, error) {
    if r.Destination == "" || r.Issuer == "" || r.ACSURL == "" {
        return nil, common.New(common.ErrCodeInvalidArgument, "destination, issuer, acs required")
    }
    if r.ID == "" {
        r.ID = "_" + common.NewID()
    }
    if r.IssueInstant.IsZero() {
        r.IssueInstant = common.NowUTC()
    }
    force := "false"
    if r.ForceAuthn {
        force = "true"
    }
    doc := fmt.Sprintf(`<?xml version="1.0"?><samlp:AuthnRequest xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol" xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion" ID="%s" Version="2.0" IssueInstant="%s" Destination="%s" AssertionConsumerServiceURL="%s" ForceAuthn="%s"><saml:Issuer>%s</saml:Issuer></samlp:AuthnRequest>`,
        xmlEscape(r.ID), r.IssueInstant.UTC().Format(time.RFC3339), xmlEscape(r.Destination), xmlEscape(r.ACSURL), force, xmlEscape(r.Issuer))
    return []byte(doc), nil
}

func xmlEscape(s string) string {
    s = strings.ReplaceAll(s, "&", "&amp;")
    s = strings.ReplaceAll(s, "<", "&lt;")
    s = strings.ReplaceAll(s, ">", "&gt;")
    s = strings.ReplaceAll(s, `"`, "&quot;")
    return s
}

// ResponseAssertion holds parsed assertion attributes (simplified).
type ResponseAssertion struct {
    Subject       string
    SessionIndex  string
    Audience      string
    Recipient     string
    NotBefore     time.Time
    NotOnOrAfter  time.Time
    Attributes    map[string][]string
}

// ParseResponseAttributes extracts attributes from simplified SAML Response XML.
func ParseResponseAttributes(xmlData []byte) (*ResponseAssertion, error) {
    type attrVal struct {
        Name string `xml:"Name,attr"`
        Val  string `xml:",chardata"`
    }
    type attrStmt struct {
        Attrs []attrVal `xml:"Attribute"`
    }
    type subject struct {
        NameID string `xml:"NameID"`
    }
    type conditions struct {
        NotBefore    string `xml:"NotBefore,attr"`
        NotOnOrAfter string `xml:"NotOnOrAfter,attr"`
        Audience     string `xml:"Audience"`
    }
    type assertion struct {
        Subject    subject    `xml:"Subject"`
        Conditions conditions `xml:"Conditions"`
        AttrStmt   attrStmt   `xml:"AttributeStatement"`
    }
    type response struct {
        Assertion assertion `xml:"Assertion"`
    }
    var resp response
    if err := xml.Unmarshal(xmlData, &resp); err != nil {
        return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid saml xml", err)
    }
    a := resp.Assertion
    out := &ResponseAssertion{
        Subject: a.Subject.NameID,
        Attributes: map[string][]string{},
    }
    if a.Conditions.Audience != "" {
        out.Audience = strings.TrimSpace(a.Conditions.Audience)
    }
    if a.Conditions.NotBefore != "" {
        t, err := time.Parse(time.RFC3339, a.Conditions.NotBefore)
        if err == nil {
            out.NotBefore = t
        }
    }
    if a.Conditions.NotOnOrAfter != "" {
        t, err := time.Parse(time.RFC3339, a.Conditions.NotOnOrAfter)
        if err == nil {
            out.NotOnOrAfter = t
        }
    }
    for _, at := range a.AttrStmt.Attrs {
        out.Attributes[at.Name] = append(out.Attributes[at.Name], strings.TrimSpace(at.Val))
    }
    return out, nil
}

// ValidateAssertion checks audience, recipient, and clock skew.
func ValidateAssertion(a *ResponseAssertion, wantAudience, wantRecipient string, now time.Time, skew time.Duration) error {
    if a == nil {
        return common.New(common.ErrCodeInvalidArgument, "assertion required")
    }
    if a.Subject == "" {
        return common.New(common.ErrCodeUnauthorized, "subject required")
    }
    if wantAudience != "" && a.Audience != wantAudience {
        return common.New(common.ErrCodeUnauthorized, "audience mismatch")
    }
    if wantRecipient != "" && a.Recipient != "" && a.Recipient != wantRecipient {
        return common.New(common.ErrCodeUnauthorized, "recipient mismatch")
    }
    if skew <= 0 {
        skew = 2 * time.Minute
    }
    if !a.NotBefore.IsZero() && now.Before(a.NotBefore.Add(-skew)) {
        return common.New(common.ErrCodeUnauthorized, "assertion not yet valid")
    }
    if !a.NotOnOrAfter.IsZero() && !now.Before(a.NotOnOrAfter.Add(skew)) {
        return common.New(common.ErrCodeUnauthorized, "assertion expired")
    }
    return nil
}

// EncodeRedirectRequest base64-encodes AuthnRequest for HTTP-Redirect binding.
func EncodeRedirectRequest(xmlData []byte) string {
    return base64.StdEncoding.EncodeToString(xmlData)
}