package pki

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// CertificateRecord stores issued cert metadata and PEM.
type CertificateRecord struct {
	ID         string    `json:"id"`
	OrgID      string    `json:"org_id"`
	CommonName string    `json:"common_name"`
	Serial     string    `json:"serial"`
	NotBefore  time.Time `json:"not_before"`
	NotAfter   time.Time `json:"not_after"`
	CertPEM    string    `json:"cert_pem"`
	KeyPEM     string    `json:"key_pem,omitempty"`
	Revoked    bool      `json:"revoked"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
}

// Authority is an in-process CA.
type Authority struct {
	mu       sync.Mutex
	caCert   *x509.Certificate
	caKey    *ecdsa.PrivateKey
	caPEM    string
	issued   map[string]*CertificateRecord
	serial   int64
}

// NewAuthority creates a self-signed CA for org/dev use.
func NewAuthority(commonName string) (*Authority, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	serial, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, err
	}
	tmpl := &x509.Certificate{
		SerialNumber: serial,
		Subject:      pkix.Name{CommonName: commonName, Organization: []string{"SentinelVault"}},
		NotBefore:    time.Now().UTC().Add(-time.Hour),
		NotAfter:     time.Now().UTC().Add(10 * 365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		IsCA:         true,
		BasicConstraintsValid: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}
	caPEM := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
	return &Authority{
		caCert: cert,
		caKey:  key,
		caPEM:  caPEM,
		issued: map[string]*CertificateRecord{},
		serial: 1,
	}, nil
}

// CACertPEM returns the CA certificate.
func (a *Authority) CACertPEM() string { return a.caPEM }

// Issue creates an end-entity certificate.
func (a *Authority) Issue(orgID, cn string, dnsNames []string, ttl time.Duration) (*CertificateRecord, error) {
	if orgID == "" || cn == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and common_name required")
	}
	if ttl <= 0 {
		ttl = 90 * 24 * time.Hour
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.serial++
	serial := big.NewInt(a.serial)
	now := time.Now().UTC()
	tmpl := &x509.Certificate{
		SerialNumber: serial,
		Subject:      pkix.Name{CommonName: cn, Organization: []string{orgID}},
		NotBefore:    now,
		NotAfter:     now.Add(ttl),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		DNSNames:     dnsNames,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, a.caCert, &key.PublicKey, a.caKey)
	if err != nil {
		return nil, err
	}
	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}
	rec := &CertificateRecord{
		ID: common.NewID(), OrgID: orgID, CommonName: cn, Serial: fmt.Sprintf("%d", a.serial),
		NotBefore: now, NotAfter: now.Add(ttl),
		CertPEM: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})),
		KeyPEM:  string(pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})),
	}
	a.issued[rec.ID] = rec
	cp := *rec
	return &cp, nil
}

// Revoke marks a certificate revoked.
func (a *Authority) Revoke(id string) (*CertificateRecord, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	rec, ok := a.issued[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "certificate not found")
	}
	now := time.Now().UTC()
	rec.Revoked = true
	rec.RevokedAt = &now
	cp := *rec
	cp.KeyPEM = ""
	return &cp, nil
}

// Get returns a certificate without private key.
func (a *Authority) Get(id string) (*CertificateRecord, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	rec, ok := a.issued[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "certificate not found")
	}
	cp := *rec
	cp.KeyPEM = ""
	return &cp, nil
}
