package dsp

import "testing"

func TestNoiseReproducibility(t *testing.T) {
	n1 := NewNoise(42)
	n2 := NewNoise(42)
	buf1 := NewMonoBuffer(128)
	buf2 := NewMonoBuffer(128)
	n1.Process(buf1)
	n2.Process(buf2)
	for i := range buf1 {
		if buf1[i] != buf2[i] {
			t.Errorf("sample %d differs between seeded instances", i)
		}
	}
}

func TestNoiseRange(t *testing.T) {
	n := NewNoise(1)
	buf := NewMonoBuffer(256)
	n.Process(buf)
	for i, s := range buf {
		if s < -1.0 || s > 1.0 {
			t.Errorf("sample %d out of range: %f", i, s)
		}
	}
}
