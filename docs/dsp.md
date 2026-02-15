# DSP Primitives

## Conventions
- Sample rate: 48000 Hz
- Block size: 64 frames
- All samples use float64 internally
- Ranges are [-1.0, 1.0] unless noted

## Units
- Oscillator: phase accumulator with wavetable interpolation
- Envelope: ADSR with per-phase duration in seconds
- Filter: SVF with low/high/band/notch modes
- LFO: sine/saw/square/triangle/sample-hold
- ModMatrix: source accumulation to destinations
