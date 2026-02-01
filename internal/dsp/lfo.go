package dsp
type LFOShape int
const ( LFOSine LFOShape = iota; LFOSaw; LFOSquare; LFOTriangle; LFOSampleHold )
type LFO struct { phase, rate, hold, sampleRate float64; shape LFOShape }
func NewLFO(sr float64) *LFO { return &LFO{sampleRate: sr, hold: 1.0} }
func (l *LFO) SetRate(rate float64) { l.rate = rate / l.sampleRate }
func (l *LFO) SetShape(s LFOShape) { l.shape = s }
func (l *LFO) Process(buf MonoBuffer) { for i := range buf { buf[i] = l.tick() } }
func (l *LFO) tick() Sample {
	l.phase += l.rate
	if l.phase >= 1.0 { l.phase -= 1.0 }
	switch l.shape {
	case LFOSine: return Sample(sin2pi(l.phase))
	case LFOSaw: return Sample(2.0*l.phase - 1.0)
	case LFOSquare: if l.phase < 0.5 { return 1.0 }; return -1.0
	case LFOTriangle:
		if l.phase < 0.5 { return Sample(4.0*l.phase - 1.0) }
		return Sample(3.0 - 4.0*l.phase)
	default: return 0
	}
}
func (l *LFO) Reset() { l.phase = 0 }
