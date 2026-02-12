package dsp
import "testing"
func TestLimiterPreventsClipping(t *testing.T) {
	l := NewLimiter(); in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 2.0 }; l.Process(in, out)
	for i, s := range out { if s > 1.0 || s < -1.0 { t.Errorf("%d clipped: %f", i, s) } }
}
func TestLimiterPassthrough(t *testing.T) {
	l := NewLimiter(); in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 0.5 }; l.Process(in, out)
	for i, s := range out { if s != in[i] { t.Errorf("%d fail: %f", i, s) } }
}
