CREATE SCHEMA IF NOT EXISTS encryption;

CREATE TABLE IF NOT EXISTS encryption.operations (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    key_id      TEXT NOT NULL,
    blob_size   INT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_encryption_operations_org_id ON encryption.operations (org_id);
CREATE INDEX IF NOT EXISTS idx_encryption_operations_key_id ON encryption.operations (key_id);
