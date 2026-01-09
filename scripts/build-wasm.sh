#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p web/public/wasm
GOOS=js GOARCH=wasm go build -o web/public/wasm/synth.wasm ./cmd/wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/public/wasm/
echo "WASM build complete."
