package dsp
import "testing"
func BenchmarkOscillatorProcess(b *testing.B) {
	o := NewOscillator(SampleRate); o.SetFreq(440.0); buf := NewMonoBuffer(BlockSize); b.ResetTimer()
	for i := 0; i < b.N; i++ { o.Process(buf) }
}
func BenchmarkSVFProcess(b *testing.B) {
	f := NewSVF(SampleRate); f.SetCutoff(1000.0); in, out := NewMonoBuffer(BlockSize), NewMonoBuffer(BlockSize); b.ResetTimer()
	for i := 0; i < b.N; i++ { f.Process(in, out) }
}
