#!/usr/bin/env bash
# SentinelVault MVP demo path (Phases 1–3).
# Usage: start auth-service and api-gateway, then ./scripts/demo-mvp.sh
set -euo pipefail

AUTH_URL="${AUTH_URL:-http://127.0.0.1:8081}"
GATEWAY_URL="${GATEWAY_URL:-http://127.0.0.1:8080}"
EMAIL="${EMAIL:-demo@sentinelvault.local}"
PASSWORD="${PASSWORD:-Str0ng!DemoPass#2026}"
ORG="${ORG:-Demo Org}"

echo "== health =="
curl -sf "$GATEWAY_URL/healthz" | tee /tmp/sv-health.json
echo

echo "== register =="
curl -sf -X POST "$AUTH_URL/v1/auth/register" \
  -H 'content-type: application/json' \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\",\"org_name\":\"$ORG\"}" | tee /tmp/sv-register.json
echo

echo "== login =="
TOKENS=$(curl -sf -X POST "$AUTH_URL/v1/auth/login" \
  -H 'content-type: application/json' \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}")
echo "$TOKENS" | tee /tmp/sv-login.json
ACCESS=$(echo "$TOKENS" | python3 -c 'import sys,json; print(json.load(sys.stdin).get("access_token",""))' 2>/dev/null || true)
echo "access_token_len=${#ACCESS}"

echo "== gateway openapi =="
curl -sf "$GATEWAY_URL/openapi.json" | head -c 200
echo
echo "Demo script complete."
