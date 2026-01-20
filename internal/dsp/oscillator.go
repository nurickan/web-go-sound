package dsp

type Waveform int

const (
	WaveSine Waveform = iota
	WaveSaw
	WaveSquare
	WaveTriangle
)

type Oscillator struct {
	phase      float64
	phaseStep  float64
	waveform   Waveform
	sampleRate float64
}

func NewOscillator(sr float64) *Oscillator {
	return &Oscillator{sampleRate: sr}
}

func (o *Oscillator) SetFreq(freq float64) {
	o.phaseStep = freq / o.sampleRate
}

func (o *Oscillator) SetWaveform(w Waveform) {
	o.waveform = w
}

func (o *Oscillator) Process(buf MonoBuffer) {
	for i := range buf {
		buf[i] = o.tick()
	}
}

func (o *Oscillator) tick() Sample {
	p := o.phase
	o.phase += o.phaseStep
	if o.phase >= 1.0 {
		o.phase -= 1.0
	}
	switch o.waveform {
	case WaveSine:
		return Sample(sin2pi(p))
	case WaveSaw:
		return Sample(2.0*p - 1.0)
	case WaveSquare:
		if p < 0.5 {
			return 1.0
		}
		return -1.0
	case WaveTriangle:
		if p < 0.5 {
			return Sample(4.0*p - 1.0)
		}
		return Sample(3.0 - 4.0*p)
	default:
		return 0
	}
}

func (o *Oscillator) Reset() {
	o.phase = 0
}

func sin2pi(x float64) float64 {
	x = x - float64(int(x))
	if x < 0 {
		x += 1.0
	}
	x = x*2.0 - 1.0
	xx := x * x
	return x * (3.1415926 + xx*(-5.1677128+xx*(2.550164+xx*(-0.599264+xx*0.082145))))
}
