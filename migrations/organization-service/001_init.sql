CREATE SCHEMA IF NOT EXISTS organizations;

CREATE TABLE IF NOT EXISTS organizations.organizations (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL,
    slug        TEXT NOT NULL UNIQUE,
    status      TEXT NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organizations.memberships (
    id          TEXT PRIMARY KEY,
    user_id     TEXT NOT NULL,
    org_id      TEXT NOT NULL REFERENCES organizations.organizations(id),
    role        TEXT NOT NULL DEFAULT 'member',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, org_id)
);

CREATE INDEX IF NOT EXISTS idx_memberships_org_id ON organizations.memberships (org_id);
