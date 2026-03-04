package sequence; import "testing"
func TestAutomationInterp(t *testing.T) {
	l := NewAutomationLane(0, "linear"); l.AddPoint(0, 0); l.AddPoint(1, 1)
	if v := l.ValueAt(0.5); v != 0.5 { t.Errorf("got %f", v) }
}
func TestAutomationEdge(t *testing.T) {
	l := NewAutomationLane(0, "linear"); if v := l.ValueAt(0.5); v != 0 { t.Error("not 0") }
	l.AddPoint(0, 0.5); if v := l.ValueAt(100); v != 0.5 { t.Errorf("got %f", v) }
}
