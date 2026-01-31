package dsp
import "testing"
func TestSVFLowPass(t *testing.T) {
	f := NewSVF(SampleRate); f.SetCutoff(500.0); f.SetResonance(0.2); f.SetMode(SVFLowPass)
	in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 1.0 }
	f.Process(in, out)
}
func TestSVFBandPass(t *testing.T) {
	f := NewSVF(SampleRate); f.SetCutoff(1000.0); f.SetResonance(0.0); f.SetMode(SVFBandPass)
	in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 1.0 }
	f.Process(in, out)
}
