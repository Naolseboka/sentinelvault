CREATE SCHEMA IF NOT EXISTS devices;

CREATE TABLE IF NOT EXISTS devices.devices (
    id           TEXT PRIMARY KEY,
    user_id      TEXT NOT NULL,
    org_id       TEXT NOT NULL,
    fingerprint  TEXT NOT NULL,
    name         TEXT,
    trust_level  TEXT NOT NULL DEFAULT 'pending',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, fingerprint)
);

CREATE INDEX IF NOT EXISTS idx_devices_user_id ON devices.devices (user_id);
