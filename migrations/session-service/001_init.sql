CREATE SCHEMA IF NOT EXISTS sessions;

CREATE TABLE IF NOT EXISTS sessions.sessions (
    id            TEXT PRIMARY KEY,
    user_id       TEXT NOT NULL,
    org_id        TEXT NOT NULL,
    device_id     TEXT,
    refresh_hash  TEXT NOT NULL UNIQUE,
    expires_at    TIMESTAMPTZ NOT NULL,
    revoked_at    TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions.sessions (user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_org_id ON sessions.sessions (org_id);
