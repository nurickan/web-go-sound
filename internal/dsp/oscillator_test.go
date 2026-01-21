package dsp

import "testing"

func TestOscillatorPhaseWrapping(t *testing.T) {
	o := NewOscillator(SampleRate)
	o.SetFreq(200.0)
	o.SetWaveform(WaveSine)
	buf := NewMonoBuffer(BlockSize)
	o.Process(buf)
	for i, s := range buf {
		if s < -1.0 || s > 1.0 {
			t.Errorf("sample %d out of range: %f", i, s)
		}
	}
}

func TestOscillatorReset(t *testing.T) {
	o := NewOscillator(SampleRate)
	o.SetFreq(440.0)
	buf := NewMonoBuffer(BlockSize)
	o.Process(buf)
	o.Reset()
	if o.phase != 0 {
		t.Error("phase not reset")
	}
}
