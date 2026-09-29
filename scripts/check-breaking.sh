#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
base="${1:-$root/openapi/baseline.yaml}"
current="$(mktemp)"
trap 'rm -f "$current"' EXIT

cd "$root"
go run . openapi > "$current"

if command -v oasdiff >/dev/null 2>&1; then
	"$(command -v oasdiff)" breaking "$base" "$current" --fail-on ERR
else
	go run github.com/oasdiff/oasdiff@v1.20.0 breaking "$base" "$current" --fail-on ERR
fi
