CREATE SCHEMA IF NOT EXISTS secrets;

CREATE TABLE IF NOT EXISTS secrets.vault_secrets (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    type        TEXT NOT NULL,
    name        TEXT NOT NULL,
    version     INT NOT NULL DEFAULT 1,
    status      TEXT NOT NULL DEFAULT 'active',
    metadata    JSONB NOT NULL DEFAULT '{}',
    ciphertext  BYTEA NOT NULL,
    key_id      TEXT NOT NULL,
    created_by  TEXT,
    updated_by  TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_vault_secrets_org_id ON secrets.vault_secrets (org_id);
CREATE INDEX IF NOT EXISTS idx_vault_secrets_status ON secrets.vault_secrets (org_id, status);
CREATE INDEX IF NOT EXISTS idx_vault_secrets_type ON secrets.vault_secrets (org_id, type);
