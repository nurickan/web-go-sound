package patch
import "testing"
func TestDefaultPatchValues(t *testing.T) {
	p := DefaultPatch()
	if p.Name != "Init" { t.Errorf("got %s", p.Name) }
	if len(p.Oscillators) != 1 { t.Errorf("got %d", len(p.Oscillators)) }
}
func TestDefaultPatchFilterOff(t *testing.T) { p := DefaultPatch(); if p.Filter.Enabled { t.Error("filter on") } }
