CREATE SCHEMA IF NOT EXISTS kms;

CREATE TABLE IF NOT EXISTS kms.keys (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    name        TEXT NOT NULL,
    version     INT NOT NULL DEFAULT 1,
    status      TEXT NOT NULL DEFAULT 'active',
    wrapped_kek BYTEA NOT NULL,
    wrap_nonce  BYTEA NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_kms_keys_org_id ON kms.keys (org_id);
CREATE INDEX IF NOT EXISTS idx_kms_keys_status ON kms.keys (org_id, status);

CREATE TABLE IF NOT EXISTS kms.wrap_audit (
    id          TEXT PRIMARY KEY,
    key_id      TEXT NOT NULL REFERENCES kms.keys (id),
    operation   TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_kms_wrap_audit_key_id ON kms.wrap_audit (key_id);
