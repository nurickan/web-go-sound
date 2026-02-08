package dsp
type DelayLine struct { buffer []Sample; writePos, size int; feedback float64 }
func NewDelayLine(maxSec float64, sr float64) *DelayLine {
	size := SecondsToSamples(maxSec); if size < 1 { size = 1 }
	return &DelayLine{buffer: make([]Sample, size), size: size}
}
func (d *DelayLine) SetDelay(sec float64) {
	s := SecondsToSamples(sec); if s >= d.size { s = d.size - 1 }; d.writePos = s
}
func (d *DelayLine) SetFeedback(fb float64) { d.feedback = Clamp(fb, 0.0, 0.99) }
func (d *DelayLine) Process(in, out MonoBuffer) { for i := range in { out[i] = d.tick(in[i]) } }
func (d *DelayLine) tick(in Sample) Sample {
	rp := (d.writePos + 1) % d.size; o := d.buffer[rp]
	d.buffer[d.writePos] = in + Sample(d.feedback)*o; d.writePos = (d.writePos + 1) % d.size; return o
}
func (d *DelayLine) Reset() { for i := range d.buffer { d.buffer[i] = 0 }; d.writePos = 0 }
