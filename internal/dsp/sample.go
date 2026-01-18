package dsp

type Sample float64

const (
	SampleRate      = 48000.0
	BlockSize       = 64
	Nyquist         = SampleRate / 2.0
	MaxMIDINote     = 127
	MinMIDINote     = 0
)

func SecondsToSamples(sec float64) int {
	return int(sec * SampleRate)
}

func SamplesToSeconds(n int) float64 {
	return float64(n) / SampleRate
}
