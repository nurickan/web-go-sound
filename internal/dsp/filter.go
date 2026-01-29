package dsp

type OnePole struct {
	a0, b1 float64
	z1     float64
	sr     float64
}

func NewOnePole(sr float64) *OnePole {
	return &OnePole{sr: sr}
}

func (f *OnePole) SetCutoff(freq float64) {
	freq = Clamp(freq, 20.0, Nyquist)
	g := freq / (freq + f.sr)
	f.a0 = g
	f.b1 = 1.0 - g
}

func (f *OnePole) Process(in, out MonoBuffer) {
	for i := range in {
		out[i] = f.tick(in[i])
	}
}

func (f *OnePole) tick(in Sample) Sample {
	out := in*f.a0 + Sample(f.z1)*f.b1
	f.z1 = float64(out)
	return out
}

func (f *OnePole) Reset() { f.z1 = 0 }
