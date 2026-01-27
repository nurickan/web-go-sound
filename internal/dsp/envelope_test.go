package dsp

import "testing"

func TestADSRPhases(t *testing.T) {
	e := NewADSR(SampleRate)
	e.AttackTime = 0.01
	e.DecayTime = 0.01
	e.SustainLevel = 0.5
	e.ReleaseTime = 0.01

	e.Gate(true)
	buf := NewMonoBuffer(BlockSize)
	e.Process(buf)

	if !e.IsActive() {
		t.Error("envelope should be active after gate on")
	}

	e.Gate(false)
	e.Process(buf)
}

func TestADSRIdle(t *testing.T) {
	e := NewADSR(SampleRate)
	buf := NewMonoBuffer(BlockSize)
	e.Process(buf)
	for i, s := range buf {
		if s != 0 {
			t.Errorf("idle envelope should output 0 at %d", i)
		}
	}
}
