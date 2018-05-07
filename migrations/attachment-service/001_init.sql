CREATE SCHEMA IF NOT EXISTS attachments;

CREATE TABLE IF NOT EXISTS attachments.attachment_records (
    id              TEXT PRIMARY KEY,
    org_id          TEXT NOT NULL,
    vault_item_id   TEXT NOT NULL,
    name            TEXT NOT NULL,
    content_type    TEXT NOT NULL DEFAULT 'application/octet-stream',
    size_bytes      BIGINT NOT NULL,
    storage_path    TEXT NOT NULL,
    key_id          TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_attachment_records_org_id ON attachments.attachment_records (org_id);
CREATE INDEX IF NOT EXISTS idx_attachment_records_vault_item ON attachments.attachment_records (vault_item_id);
