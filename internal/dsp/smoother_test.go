package dsp
import "testing"
func TestSmootherConvergence(t *testing.T) {
	s := NewSmoother(SampleRate); s.SetTarget(1.0, 0.01)
	buf := NewMonoBuffer(BlockSize); s.Process(buf)
	if buf[BlockSize-1] < 0.5 { t.Errorf("not converge: %f", buf[BlockSize-1]) }
}
func TestSmootherReset(t *testing.T) {
	s := NewSmoother(SampleRate); s.SetTarget(1.0, 0.01); buf := NewMonoBuffer(BlockSize); s.Process(buf)
	s.Reset(); if s.Current() != 0 { t.Error("reset failed") }
}
