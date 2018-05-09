CREATE SCHEMA IF NOT EXISTS certificate;

CREATE TABLE IF NOT EXISTS certificate.certificates (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    common_name TEXT NOT NULL,
    serial      TEXT NOT NULL,
    not_before  TIMESTAMPTZ NOT NULL,
    not_after   TIMESTAMPTZ NOT NULL,
    cert_pem    TEXT NOT NULL,
    revoked     BOOLEAN NOT NULL DEFAULT FALSE,
    revoked_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_certificate_certs_org ON certificate.certificates (org_id);
