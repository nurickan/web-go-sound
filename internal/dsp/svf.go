package dsp

type SVFMode int
const ( SVFLowPass SVFMode = iota; SVFHighPass; SVFBandPass; SVFNotch )

type SVF struct { ic1eq, ic2eq, g, h, r float64; mode SVFMode; sr float64 }

func NewSVF(sr float64) *SVF { return &SVF{sr: sr, h: 1.0} }

func (f *SVF) SetCutoff(freq float64) {
	freq = Clamp(freq, 20.0, Nyquist*0.95)
	g := tan(freq * 3.14159265 / f.sr); f.g = g
	f.h = 1.0 / (1.0 + 2.0*f.r*g + g*g)
}

func (f *SVF) SetResonance(res float64) { f.r = Clamp(res, 0.0, 1.0) }
func (f *SVF) SetMode(m SVFMode) { f.mode = m }
func (f *SVF) Process(in, out MonoBuffer) { for i := range in { out[i] = f.tick(in[i]) } }

func (f *SVF) tick(in Sample) Sample {
	v3 := in - Sample(f.ic2eq)
	v1 := f.ic1eq + f.g*float64(v3)
	v2 := f.ic2eq + f.g*f.ic1eq
	f.ic1eq = 2.0*v1 - f.ic1eq; f.ic2eq = 2.0*v2 - f.ic2eq
	switch f.mode {
	case SVFLowPass: return Sample(float64(v2) * f.h)
	case SVFHighPass: return Sample((float64(in) - v1 - f.r*v2) * f.h)
	case SVFBandPass: return Sample(v1 * f.h)
	default: return Sample(float64(v2) * f.h)
	}
}

func (f *SVF) Reset() { f.ic1eq, f.ic2eq = 0, 0 }
func tan(x float64) float64 { s := sin2pi(x / 3.14159265); return s / (1.0 - s*s) }
