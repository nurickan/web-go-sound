# Synth Engine

## Lifecycle
1. NoteOn -> voice allocator finds voice -> voice configures oscillator+envelope
2. Render -> each active voice renders block -> mixed to stereo output
3. NoteOff -> envelope enters release -> voice freed when complete
4. AllNotesOff -> panic / stop all voices

## Patch Flow
Patch from UI -> WASM API validates -> engine renders with validated patch
