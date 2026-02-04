package dsp
type Smoother struct { current, target, step, sampleRate float64 }
func NewSmoother(sr float64) *Smoother { return &Smoother{sampleRate: sr} }
func (s *Smoother) SetTarget(tgt, timeSec float64) {
	s.target = tgt
	samples := SecondsToSamples(timeSec)
	if samples < 1 { samples = 1 }
	s.step = (tgt - s.current) / float64(samples)
}
func (s *Smoother) Process(buf MonoBuffer) { for i := range buf { buf[i] = Sample(s.tick()) } }
func (s *Smoother) tick() float64 {
	diff := s.target - s.current
	if diff*diff < s.step*s.step { s.current = s.target } else { s.current += s.step }
	return s.current
}
func (s *Smoother) Reset() { s.current = 0; s.target = 0; s.step = 0 }
func (s *Smoother) Current() float64 { return s.current }
