CREATE SCHEMA IF NOT EXISTS approvals;

CREATE TABLE IF NOT EXISTS approvals.approval_requests (
    id            TEXT PRIMARY KEY,
    org_id        TEXT NOT NULL,
    requester_id  TEXT NOT NULL,
    approver_id   TEXT,
    resource_type TEXT NOT NULL DEFAULT '',
    resource_id   TEXT NOT NULL,
    action        TEXT NOT NULL DEFAULT '',
    reason        TEXT NOT NULL DEFAULT '',
    status        TEXT NOT NULL DEFAULT 'pending',
    expires_at    TIMESTAMPTZ NOT NULL,
    decided_at    TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_approval_requests_org_status ON approvals.approval_requests (org_id, status);
