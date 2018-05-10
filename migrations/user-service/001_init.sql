CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.users (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    email        TEXT NOT NULL,
    display_name TEXT NOT NULL,
    status       TEXT NOT NULL DEFAULT 'active',
    mfa_enabled  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_org_email ON users.users (org_id, email);
CREATE INDEX IF NOT EXISTS idx_users_org_id ON users.users (org_id);
