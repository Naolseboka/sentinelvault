CREATE SCHEMA IF NOT EXISTS versioning;

CREATE TABLE IF NOT EXISTS versioning.item_versions (
    id          TEXT PRIMARY KEY,
    item_id     TEXT NOT NULL,
    org_id      TEXT NOT NULL,
    version     INT NOT NULL,
    type        TEXT NOT NULL,
    name        TEXT NOT NULL,
    metadata    JSONB NOT NULL DEFAULT '{}',
    ciphertext  BYTEA NOT NULL,
    key_id      TEXT NOT NULL,
    created_by  TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_item_versions_item_id ON versioning.item_versions (item_id);
CREATE INDEX IF NOT EXISTS idx_item_versions_org_id ON versioning.item_versions (org_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_item_versions_item_version ON versioning.item_versions (item_id, version);
