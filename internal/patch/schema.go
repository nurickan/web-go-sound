package patch
type OscillatorConfig struct { Waveform int; Octave int; Detune, Level, Pan float64 }
type FilterConfig struct { Type int; Cutoff, Resonance float64; Enabled bool }
type EnvelopeConfig struct { AttackTime, DecayTime, SustainLevel, ReleaseTime float64 }
type LFOMapping struct { Target int; Shape int; Rate, Depth float64 }
type Patch struct {
	Name string; Oscillators []OscillatorConfig; Filter FilterConfig
	Amplitude EnvelopeConfig; FilterEnv EnvelopeConfig; LFOs []LFOMapping
}
func DefaultPatch() Patch {
	return Patch{Name: "Init",
		Oscillators: []OscillatorConfig{{Waveform: 0, Octave: 4, Detune: 0, Level: 0.8, Pan: 0}},
		Filter: FilterConfig{Type: 0, Cutoff: 18000, Resonance: 0, Enabled: false},
		Amplitude: EnvelopeConfig{AttackTime: 0.01, DecayTime: 0.1, SustainLevel: 0.8, ReleaseTime: 0.3},
		FilterEnv: EnvelopeConfig{AttackTime: 0.02, DecayTime: 0.2, SustainLevel: 0.5, ReleaseTime: 0.5},
	}
}
