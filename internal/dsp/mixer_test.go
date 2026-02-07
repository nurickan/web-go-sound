package dsp
import "testing"
func TestEqualPowerPan(t *testing.T) {
	p := NewStereoPan(); p.SetPan(0.0); diff := p.gainL - p.gainR
	if diff*diff > 0.001 { t.Errorf("center pan: L=%f R=%f", p.gainL, p.gainR) }
}
func TestPanHardLeft(t *testing.T) { p := NewStereoPan(); p.SetPan(-1.0); if p.gainL < p.gainR { t.Error("L<R") } }
func TestPanHardRight(t *testing.T) { p := NewStereoPan(); p.SetPan(1.0); if p.gainR < p.gainL { t.Error("R<L") } }
