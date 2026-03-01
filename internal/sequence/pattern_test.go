package sequence; import "testing"
func TestPatternBounds(t *testing.T) { p := NewEmptyPattern("t", 16); if p.NumSteps() != 16 { t.Errorf("got %d", p.NumSteps()) } }
func TestPatternSet(t *testing.T) { p := NewEmptyPattern("t", 8); p.SetStep(3, Step{Active: true, Note: 72, Velocity: 1, Gate: 0.75}); if !p.Steps[3].Active { t.Error("not active") } }
