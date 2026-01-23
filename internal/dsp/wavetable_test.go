package dsp

import "testing"

func TestWavetableRanges(t *testing.T) {
	tables := []*Wavetable{NewSineTable(), NewSawTable(), NewSquareTable(), NewTriangleTable()}
	for _, wt := range tables {
		for i := 0; i < TableSize; i++ {
			v := wt.Read(i)
			if v < -1.0 || v > 1.0 {
				t.Errorf("table value out of range at %d: %f", i, v)
			}
		}
	}
}

func TestWavetableContinuity(t *testing.T) {
	wt := NewSineTable()
	last := wt.Read(TableSize - 1)
	first := wt.Read(0)
	diff := last - first
	if diff*diff > 0.01 {
		t.Errorf("discontinuity at table wrap: %f", diff)
	}
}
