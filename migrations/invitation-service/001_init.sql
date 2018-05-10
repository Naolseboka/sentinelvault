CREATE SCHEMA IF NOT EXISTS invitations;

CREATE TABLE IF NOT EXISTS invitations.invitations (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    email        TEXT NOT NULL,
    role         TEXT NOT NULL DEFAULT 'member',
    token        TEXT NOT NULL UNIQUE,
    status       TEXT NOT NULL DEFAULT 'pending',
    expires_at   TIMESTAMPTZ NOT NULL,
    accepted_at  TIMESTAMPTZ,
    rejected_at  TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_invitations_org_id ON invitations.invitations (org_id);
CREATE INDEX IF NOT EXISTS idx_invitations_email ON invitations.invitations (email);
