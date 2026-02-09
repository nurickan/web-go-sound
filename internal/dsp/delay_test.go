package dsp
import "testing"
func TestDelayTapTiming(t *testing.T) {
	d := NewDelayLine(0.5, SampleRate); d.SetDelay(0.01); d.SetFeedback(0.0)
	in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize); in[0] = 1.0; d.Process(in, out)
	if out[0] != 0 { t.Error("first should be silence") }
}
func TestDelayFeedbackClamp(t *testing.T) {
	d := NewDelayLine(0.5, SampleRate); d.SetFeedback(2.0)
	if d.feedback > 0.99 { t.Error("feedback not clamped") }
}
