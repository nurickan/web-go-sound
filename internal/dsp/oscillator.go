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
	table      *Wavetable
}

func NewOscillator(sr float64) *Oscillator {
	return &Oscillator{sampleRate: sr, table: NewSineTable()}
}

func (o *Oscillator) SetFreq(freq float64) {
	o.phaseStep = freq / o.sampleRate
}

func (o *Oscillator) SetWaveform(w Waveform) {
	o.waveform = w
	switch w {
	case WaveSine:
		o.table = NewSineTable()
	case WaveSaw:
		o.table = NewSawTable()
	case WaveSquare:
		o.table = NewSquareTable()
	case WaveTriangle:
		o.table = NewTriangleTable()
	}
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
	idx := int(p * TableSize)
	frac := (p*TableSize - float64(idx))
	a := o.table.Read(idx)
	b := o.table.Read(idx + 1)
	return a + Sample(frac)*float64(b-a)
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
