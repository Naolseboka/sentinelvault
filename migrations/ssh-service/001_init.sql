CREATE SCHEMA IF NOT EXISTS ssh;

CREATE TABLE IF NOT EXISTS ssh.keys (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    name        TEXT NOT NULL,
    type        TEXT NOT NULL DEFAULT 'user',
    public_key  TEXT NOT NULL,
    fingerprint TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    rotated_at  TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS ssh.signed_certs (
    id           TEXT PRIMARY KEY,
    key_id       TEXT NOT NULL REFERENCES ssh.keys(id),
    org_id       TEXT NOT NULL,
    principals   TEXT[] NOT NULL DEFAULT '{}',
    valid_after  TIMESTAMPTZ NOT NULL,
    valid_before TIMESTAMPTZ NOT NULL,
    cert         TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_ssh_keys_org ON ssh.keys (org_id);
