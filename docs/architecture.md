# Architecture

Web-go-sound is a modular browser-based audio synthesizer.

## Data flow

User controls in React UI > patchStore > WASM API > Go synth engine > stereo blocks > ring buffer > AudioWorklet > device

## Boundaries

- Go core knows nothing about DOM
- React knows no DSP details
- WASM API is a versioned JSON contract

## Directories

- internal/dsp � pure DSP primitives
- internal/synth � voice allocation, engine, events
- internal/patch � schema, validation, presets
- internal/sequence � clock, pattern, player, automation
- internal/render � offline render, WAV encoding
- pkg/wasmapi � WASM API layer
- cmd/wasm � WASM entry point
- web/src � React UI
