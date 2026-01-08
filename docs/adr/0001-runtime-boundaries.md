# ADR-0001: Runtime Boundaries

Status: Accepted

## Context

Go WASM runs in a separate runtime from the browser AudioWorklet. Go may introduce unpredictable pauses.

## Decision

- Go renders audio blocks (64-256 frames) into a ring buffer
- AudioWorklet reads from the ring buffer on the audio thread
- The worklet plays silence on underrun
- No Go code runs on the audio thread
