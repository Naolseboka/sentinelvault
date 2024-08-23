#!/usr/bin/env bash
set -euo pipefail

echo "Checking SentinelVault bootstrap prerequisites..."

need() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "MISSING: $1"
    exit 1
  fi
  echo "OK: $1 ($($1 --version 2>&1 | head -1))"
}

need go
GO_VERSION=$(go env GOVERSION | tr -d 'go')
MAJOR=$(echo "$GO_VERSION" | cut -d. -f1)
MINOR=$(echo "$GO_VERSION" | cut -d. -f2)
if [ "$MAJOR" -lt 1 ] || { [ "$MAJOR" -eq 1 ] && [ "$MINOR" -lt 22 ]; }; then
  echo "Go 1.22+ required, found $GO_VERSION"
  exit 1
fi

if command -v golangci-lint >/dev/null 2>&1; then
  echo "OK: golangci-lint"
else
  echo "WARN: golangci-lint not installed (optional for local dev)"
fi

if command -v docker >/dev/null 2>&1; then
  echo "OK: docker"
else
  echo "WARN: docker not installed (optional for local stack)"
fi

echo "Bootstrap check passed."
