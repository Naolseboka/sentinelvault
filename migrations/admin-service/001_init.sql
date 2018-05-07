CREATE SCHEMA IF NOT EXISTS admin;

CREATE TABLE IF NOT EXISTS admin.feature_flags (
    name        TEXT PRIMARY KEY,
    enabled     BOOLEAN NOT NULL DEFAULT FALSE,
    description TEXT NOT NULL DEFAULT '',
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS admin.org_states (
    org_id     TEXT PRIMARY KEY,
    suspended  BOOLEAN NOT NULL DEFAULT FALSE,
    reason     TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
