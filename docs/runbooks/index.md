# Runbooks

Operational guides for common incidents and maintenance:

- [Backup & Disaster Recovery](backup-dr.md)
- [Incident Response](incident-response.md)

## Service Won't Start
1. Check `CONFIG_PATH` points to valid YAML
2. Verify `SERVICE_NAME` is set
3. Check logs for config validation errors

## Database Connection Failed
1. Confirm Postgres is running: `docker compose -f deploy/docker/docker-compose.yml ps`
2. Verify `DATABASE_URL`
3. Run migrations: apply `migrations/<service>/001_init.sql`

## Gateway 502 Upstream
1. Confirm target service is healthy: `curl localhost:<port>/healthz`
2. Check upstream mapping in `apps/api-gateway/internal/gateway/gateway.go`

## High Error Rate
1. Check Prometheus/Grafana dashboards
2. Review structured logs for `request_id`
3. Trace via Jaeger/Tempo if OTel endpoint configured
<!-- touch 7861 -->
