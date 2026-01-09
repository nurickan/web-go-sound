$ErrorActionPreference = "Stop"
$projectRoot = Split-Path -Parent $PSScriptRoot
Set-Location $projectRoot
$outDir = "web\\public\\wasm"
if (-not (Test-Path $outDir)) { New-Item -ItemType Directory -Path $outDir -Force }
$env:GOOS = "js"
$env:GOARCH = "wasm"
go build -o "$outDir\\synth.wasm" .\\cmd\\wasm
$goroot = go env GOROOT
Copy-Item "$goroot\\misc\\wasm\\wasm_exec.js" -Destination "$outDir\\"
Write-Host "WASM build complete."
