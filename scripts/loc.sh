#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "== Go files =="
# Sum lines across xargs batches (path may contain spaces; wc "total" is per-batch).
find . -name '*.go' -not -path './vendor/*' -print0 | xargs -0 cat | wc -l

echo "== SQL migrations =="
find migrations -name '*.sql' -print0 2>/dev/null | xargs -0 cat 2>/dev/null | wc -l || echo 0

echo "== Proto =="
find proto -name '*.proto' -print0 2>/dev/null | xargs -0 cat 2>/dev/null | wc -l || echo 0

echo "== Docs =="
find docs -name '*.md' -print0 2>/dev/null | xargs -0 cat 2>/dev/null | wc -l || echo 0
