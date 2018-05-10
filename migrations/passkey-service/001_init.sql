CREATE SCHEMA IF NOT EXISTS passkey;

CREATE TABLE IF NOT EXISTS passkey.credentials (
    id            TEXT PRIMARY KEY,
    org_id        TEXT NOT NULL DEFAULT '',
    user_id       TEXT NOT NULL,
    name          TEXT NOT NULL DEFAULT '',
    credential_id TEXT NOT NULL,
    public_key    TEXT NOT NULL,
    sign_count    INT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_passkey_credentials_user ON passkey.credentials (user_id);
