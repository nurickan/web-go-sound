# ADR-001: WASM Audio Architecture

## Decision
Use Go compiled to WASM for synth engine, AudioWorklet for output.

## Rationale
- Go offers safe concurrency, suitable for realtime audio
- AudioWorklet runs on dedicated thread, avoids main-thread jank
- WASM bridges allow JS UI to control engine without shared memory complexity

## Consequences
- WASM binary ~1�2MB gzipped
- Requires SharedArrayBuffer or copy-based message passing
