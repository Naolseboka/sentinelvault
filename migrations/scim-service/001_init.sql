CREATE SCHEMA IF NOT EXISTS scim;

CREATE TABLE IF NOT EXISTS scim.users (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    external_id  TEXT,
    user_name    TEXT NOT NULL,
    active       BOOLEAN NOT NULL DEFAULT TRUE,
    name         JSONB NOT NULL DEFAULT '{}',
    emails       JSONB NOT NULL DEFAULT '[]',
    schemas      TEXT[] NOT NULL DEFAULT '{}',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_scim_users_org ON scim.users (org_id);
