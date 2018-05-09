CREATE SCHEMA IF NOT EXISTS billing;

CREATE TABLE IF NOT EXISTS billing.licenses (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL UNIQUE,
    plan       TEXT NOT NULL DEFAULT 'free',
    seats      INT NOT NULL DEFAULT 1,
    features   JSONB NOT NULL DEFAULT '{}',
    valid_from TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    valid_to   TIMESTAMPTZ NOT NULL,
    active     BOOLEAN NOT NULL DEFAULT TRUE,
    metadata   JSONB
);

CREATE INDEX IF NOT EXISTS idx_billing_licenses_org ON billing.licenses (org_id);
