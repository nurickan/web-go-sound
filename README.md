# web-go-sound

Modular browser-based audio synthesizer with Go DSP core compiled to WASM.

## Quick start

```bash
make setup
make build-wasm
make dev
```

## Architecture

Go DSP > WASM > WebAudio AudioWorklet > Output

See [docs/architecture.md](docs/architecture.md).
