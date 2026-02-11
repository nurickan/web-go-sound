package dsp
type Limiter struct { threshold, gain float64 }
func NewLimiter() *Limiter { return &Limiter{threshold: 0.95, gain: 1.0} }
func (l *Limiter) SetThreshold(t float64) { l.threshold = Clamp(t, 0.1, 1.0) }
func (l *Limiter) Process(in, out MonoBuffer) {
	peak := Sample(0)
	for _, s := range in { a := s; if a < 0 { a = -a }; if a > peak { peak = a } }
	if peak > Sample(l.threshold) { l.gain = l.threshold / float64(peak) } else { l.gain = 1.0 }
	for i := range in { out[i] = Sample(float64(in[i]) * l.gain) }
}
