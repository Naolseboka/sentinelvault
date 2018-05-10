CREATE SCHEMA IF NOT EXISTS notification;

CREATE TABLE IF NOT EXISTS notification.messages (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    channel    TEXT NOT NULL,
    recipient  TEXT NOT NULL,
    subject    TEXT NOT NULL DEFAULT '',
    body       TEXT NOT NULL DEFAULT '',
    priority   TEXT NOT NULL DEFAULT 'normal',
    metadata   JSONB,
    status     TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notification_messages_org_id ON notification.messages (org_id);
