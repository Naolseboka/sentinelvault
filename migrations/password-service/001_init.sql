CREATE SCHEMA IF NOT EXISTS passwords;

CREATE TABLE IF NOT EXISTS passwords.vault_items (
    id              TEXT PRIMARY KEY,
    org_id          TEXT NOT NULL,
    type            TEXT NOT NULL,
    name            TEXT NOT NULL,
    version         INT NOT NULL DEFAULT 1,
    status          TEXT NOT NULL DEFAULT 'active',
    metadata        JSONB NOT NULL DEFAULT '{}',
    ciphertext      BYTEA NOT NULL,
    key_id          TEXT NOT NULL,
    collection_id   TEXT,
    folder_id       TEXT,
    created_by      TEXT,
    updated_by      TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_vault_items_org_id ON passwords.vault_items (org_id);
CREATE INDEX IF NOT EXISTS idx_vault_items_status ON passwords.vault_items (org_id, status);
CREATE INDEX IF NOT EXISTS idx_vault_items_collection ON passwords.vault_items (collection_id);

CREATE TABLE IF NOT EXISTS passwords.collections (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    name        TEXT NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_collections_org_id ON passwords.collections (org_id);

CREATE TABLE IF NOT EXISTS passwords.folders (
    id              TEXT PRIMARY KEY,
    org_id          TEXT NOT NULL,
    collection_id   TEXT NOT NULL REFERENCES passwords.collections (id),
    parent_id       TEXT,
    name            TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_folders_collection ON passwords.folders (collection_id);

CREATE TABLE IF NOT EXISTS passwords.share_grants (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    item_id     TEXT NOT NULL REFERENCES passwords.vault_items (id),
    grantee_id  TEXT NOT NULL,
    permission  TEXT NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    revoked_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_share_grants_item_id ON passwords.share_grants (item_id);
