package sequence; import "math"
type Clock struct { bpm, samplesPerBeat, sampleRate float64; tick int }
func NewClock(bpm, sr float64) *Clock { c := &Clock{sampleRate: sr}; c.SetBPM(bpm); return c }
func (c *Clock) SetBPM(bpm float64) { c.bpm = math.Max(20, math.Min(300, bpm)); c.samplesPerBeat = c.sampleRate / (c.bpm / 60.0) }
func (c *Clock) BPM() float64 { return c.bpm }
func (c *Clock) Advance(n int) int { c.tick++; return int(float64(n) / c.samplesPerBeat) }
func (c *Clock) Reset() { c.tick = 0 }
func (c *Clock) SamplesPerBeat() float64 { return c.samplesPerBeat }
