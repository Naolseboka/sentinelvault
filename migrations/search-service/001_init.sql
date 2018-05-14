CREATE SCHEMA IF NOT EXISTS search;

CREATE TABLE IF NOT EXISTS search.documents (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    index_name TEXT NOT NULL DEFAULT 'default',
    title      TEXT NOT NULL DEFAULT '',
    body       TEXT NOT NULL DEFAULT '',
    tags       TEXT[] NOT NULL DEFAULT '{}',
    fields     JSONB,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_search_documents_org_index ON search.documents (org_id, index_name);
