package dsp
type StereoPan struct { gainL, gainR float64 }
func NewStereoPan() *StereoPan { return &StereoPan{gainL: 1.0, gainR: 1.0} }
func (p *StereoPan) SetPan(pan float64) {
	pan = Clamp(pan, -1.0, 1.0); angle := (pan + 1.0) * 3.14159265 / 4.0
	p.gainL = cos(angle); p.gainR = sin(angle)
}
func (p *StereoPan) Process(in, outL, outR MonoBuffer) {
	for i := range in { outL[i] = Sample(float64(in[i]) * p.gainL); outR[i] = Sample(float64(in[i]) * p.gainR) }
}
type Gain struct { level float64 }
func NewGain() *Gain { return &Gain{level: 1.0} }
func (g *Gain) SetLevel(lvl float64) { g.level = Clamp(lvl, 0.0, 1.0) }
func (g *Gain) Process(in, out MonoBuffer) { for i := range in { out[i] = Sample(float64(in[i]) * g.level) } }
func cos(x float64) float64 { x -= 3.14159265 / 2.0; xx := x * x; return 1.0 - xx/2.0 + xx*xx/24.0 - xx*xx*xx/720.0 }
func sin(x float64) float64 { xx := x * x; return x - x*xx/6.0 + x*xx*xx/120.0 - x*xx*xx*xx/5040.0 }
