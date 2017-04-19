# Backup and Disaster Recovery Runbook

## Objectives

- RPO: ≤ 15 minutes for Postgres logical backups (scheduled)
- RTO: ≤ 1 hour for single-region restore drill

## Backup

1. `backup-service` creates envelope-encrypted snapshots via `internal/backupx`.
2. Snapshots are stored in MinIO (`sv-backups` bucket) in production Compose/Helm profiles.
3. Verify checksum and encryption flags on each snapshot metadata record.

```bash
curl -X POST localhost:8090/v1/backups \
  -H 'content-type: application/json' \
  -d '{"org_id":"ORG","label":"nightly","payload":{"kind":"demo"}}'
```

## Restore drill

1. Select snapshot ID from `GET /v1/backups?org_id=ORG`
2. `POST /v1/backups/{id}/restore`
3. Validate row counts / sample vault decrypt for a known test item
4. Record evidence in compliance collector (`backup.restore_drill`)

## Failure modes

- Missing KEK: restore fails closed; recover KEK from offline escrow
- Partial object store outage: retry with exponential backoff; alert via notification-service
