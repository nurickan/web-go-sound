package patch
import "errors"
var (
	ErrNoOscillators = errors.New("need >=1 oscillator")
	ErrInvalidWaveform = errors.New("waveform 0-3")
	ErrCutoffRange = errors.New("cutoff 20-20000")
	ErrTimeRange = errors.New("time >= 0.001")
)
func Validate(p Patch) error {
	if len(p.Oscillators) == 0 { return ErrNoOscillators }
	for _, o := range p.Oscillators {
		if o.Waveform < 0 || o.Waveform > 3 { return ErrInvalidWaveform }
		if o.Level < 0 || o.Level > 1 { return ErrInvalidWaveform }
	}
	if p.Filter.Cutoff < 20 || p.Filter.Cutoff > 20000 { return ErrCutoffRange }
	if p.Amplitude.AttackTime < 0.001 || p.Amplitude.ReleaseTime < 0.001 { return ErrTimeRange }
	return nil
}
