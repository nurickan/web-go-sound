package dsp
import "testing"
func TestChorusBounds(t *testing.T) {
	c := NewChorus(SampleRate); in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 0.5 }; c.Process(in, out)
	for i, s := range out { if s < -1.0 || s > 1.0 { t.Errorf("%d out: %f", i, s) } }
}
