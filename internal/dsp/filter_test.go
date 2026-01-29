package dsp

import "testing"

func TestOnePoleSettles(t *testing.T) {
	f := NewOnePole(SampleRate)
	f.SetCutoff(1000.0)
	in := NewMonoBuffer(BlockSize)
	out := NewMonoBuffer(BlockSize)
	for i := range in { in[i] = 1.0 }
	f.Process(in, out)
	if out[BlockSize-1] < 0.9 {
		t.Error("output should approach 1.0")
	}
}

func TestOnePoleZeroInput(t *testing.T) {
	f := NewOnePole(SampleRate)
	f.SetCutoff(1000.0)
	in := NewMonoBuffer(BlockSize)
	out := NewMonoBuffer(BlockSize)
	f.Process(in, out)
	for _, s := range out {
		if s != 0 { t.Error("zero input should give zero output"); break }
	}
}
