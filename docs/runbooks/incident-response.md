# Incident Response Runbook

## Severity

| Level | Examples | Response |
|-------|----------|----------|
| SEV1 | Auth outage, KMS unavailable, audit chain break | Page on-call, freeze rotations |
| SEV2 | Elevated risk scores, mass failed logins | Enable step-up MFA, review threat findings |
| SEV3 | Single org misconfig | Ticket + admin-service remediation |

## First 15 minutes

1. Confirm blast radius via Grafana dashboards and audit search
2. Revoke compromised sessions via session-service
3. Emergency rotate affected secrets via rotation-service (`mode=emergency`)
4. Open emergency access only with justification + audit event

## Evidence

- Export signed audit chain segment
- Capture risk/threat findings for the window
- Attach to compliance incident control evidence

## Communications

Use notification-service channels: PagerDuty/Opsgenie for SEV1, Slack/Teams for SEV2+.
