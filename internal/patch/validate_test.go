package patch
import "testing"
func TestValidateDefault(t *testing.T) { if err := Validate(DefaultPatch()); err != nil { t.Fatal(err) } }
func TestValidateEmpty(t *testing.T) {
	p := DefaultPatch(); p.Oscillators = nil
	if err := Validate(p); err != ErrNoOscillators { t.Errorf("got %v", err) }
}
func TestValidateBadWaveform(t *testing.T) {
	p := DefaultPatch(); p.Oscillators[0].Waveform = 99
	if err := Validate(p); err != ErrInvalidWaveform { t.Errorf("got %v", err) }
}
