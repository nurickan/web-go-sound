package dsp
type Chorus struct { delay *DelayLine; lfo *LFO; depth, rate, mix float64 }
func NewChorus(sr float64) *Chorus {
	return &Chorus{delay: NewDelayLine(0.05, sr), lfo: NewLFO(sr), depth: 0.005, rate: 0.5, mix: 0.3}
}
func (c *Chorus) SetDepth(d float64) { c.depth = Clamp(d, 0.0, 0.02) }
func (c *Chorus) SetRate(r float64) { c.rate = r; c.lfo.SetRate(r) }
func (c *Chorus) SetMix(m float64) { c.mix = Clamp(m, 0.0, 1.0) }
func (c *Chorus) Process(in, out MonoBuffer) {
	mod := NewMonoBuffer(BlockSize); c.lfo.Process(mod)
	for i := range in {
		ms := (float64(mod[i]) + 1.0) / 2.0 * c.depth; c.delay.SetDelay(ms)
		w := c.delay.tick(in[i]); out[i] = Sample(float64(in[i])*(1.0-c.mix) + float64(w)*c.mix)
	}
}
