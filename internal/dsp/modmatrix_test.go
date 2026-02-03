package dsp
import "testing"
func TestModMatrixScaling(t *testing.T) {
	m := NewModMatrix(); m.SetSource(1, 0.5); m.AddRoute(ModRoute{Source: 1, Dest: 10, Amount: 0.5})
	v := m.GetModulation(10)
	if v != 0.25 { t.Errorf("expected 0.25, got %f", v) }
}
func TestModMatrixSumming(t *testing.T) {
	m := NewModMatrix(); m.SetSource(1, 0.5); m.SetSource(2, -0.5)
	m.AddRoute(ModRoute{Source: 1, Dest: 10, Amount: 1.0})
	m.AddRoute(ModRoute{Source: 2, Dest: 10, Amount: 1.0})
	v := m.GetModulation(10)
	if v != 0.0 { t.Errorf("expected 0.0, got %f", v) }
}
