CREATE SCHEMA IF NOT EXISTS otp;

CREATE TABLE IF NOT EXISTS otp.totp_seeds (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    user_id     TEXT NOT NULL UNIQUE,
    label       TEXT NOT NULL DEFAULT '',
    secret_blob BYTEA NOT NULL,
    enabled     BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_otp_totp_org ON otp.totp_seeds (org_id);
