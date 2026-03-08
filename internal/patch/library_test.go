package patch; import "testing"
func TestLibrary(t *testing.T) {
	l := NewLibrary(); p := DefaultPatch(); l.Add(p)
	if len(l.List()) != 1 { t.Errorf("got %d", len(l.List())) }
	if _, ok := l.Get("Init"); !ok { t.Error("Init missing") }
}
