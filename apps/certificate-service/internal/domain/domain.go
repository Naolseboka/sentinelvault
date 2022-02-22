package domain

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/pki"
)

// CertificateRecord re-exports pki.CertificateRecord.
type CertificateRecord = pki.CertificateRecord

// IssueInput holds certificate issuance fields.
type IssueInput struct {
	OrgID      string        `json:"org_id"`
	CommonName string        `json:"common_name"`
	DNSNames   []string      `json:"dns_names"`
	TTL        time.Duration `json:"ttl"`
}

// Service wraps the PKI authority.
type Service struct {
	authority *pki.Authority
}

// NewService creates a certificate domain service.
func NewService(a *pki.Authority) *Service {
	return &Service{authority: a}
}

// NewAuthority creates a self-signed CA.
func NewAuthority(cn string) (*pki.Authority, error) {
	return pki.NewAuthority(cn)
}

// Issue creates an end-entity certificate.
func (s *Service) Issue(_ context.Context, in IssueInput) (*CertificateRecord, error) {
	return s.authority.Issue(in.OrgID, in.CommonName, in.DNSNames, in.TTL)
}

// Revoke marks a certificate revoked.
func (s *Service) Revoke(_ context.Context, id string) (*CertificateRecord, error) {
	return s.authority.Revoke(id)
}

// Get returns certificate metadata without private key.
func (s *Service) Get(_ context.Context, id string) (*CertificateRecord, error) {
	return s.authority.Get(id)
}

// CACertPEM returns the CA certificate PEM.
func (s *Service) CACertPEM(_ context.Context) string {
	return s.authority.CACertPEM()
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "certificate" }

// TableName returns the primary table name.
func TableName() string { return "certificates" }
