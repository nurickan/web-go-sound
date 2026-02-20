package synth
import ("github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
type Voice struct { note int; velocity float64; gate bool; osc *dsp.Oscillator; ampEnv *dsp.ADSR; filter *dsp.SVF; filterEnv *dsp.ADSR }
func NewVoice(sr float64) *Voice {
	return &Voice{osc: dsp.NewOscillator(sr), ampEnv: dsp.NewADSR(sr), filter: dsp.NewSVF(sr), filterEnv: dsp.NewADSR(sr)}
}
func (v *Voice) NoteOn(note int, velocity float64, p patch.Patch) {
	v.note = note; v.velocity = velocity; v.gate = true
	v.osc.SetFreq(dsp.FreqFromMIDI(note))
	v.osc.SetWaveform(dsp.Waveform(p.Oscillators[0].Waveform))
	v.ampEnv.AttackTime = p.Amplitude.AttackTime; v.ampEnv.DecayTime = p.Amplitude.DecayTime
	v.ampEnv.SustainLevel = p.Amplitude.SustainLevel; v.ampEnv.ReleaseTime = p.Amplitude.ReleaseTime; v.ampEnv.Gate(true)
	if p.Filter.Enabled {
		v.filter.SetCutoff(p.Filter.Cutoff); v.filter.SetResonance(p.Filter.Resonance)
		v.filter.SetMode(dsp.SVFMode(p.Filter.Type))
		v.filterEnv.AttackTime = p.FilterEnv.AttackTime; v.filterEnv.DecayTime = p.FilterEnv.DecayTime
		v.filterEnv.SustainLevel = p.FilterEnv.SustainLevel; v.filterEnv.ReleaseTime = p.FilterEnv.ReleaseTime; v.filterEnv.Gate(true)
	}
}
func (v *Voice) NoteOff() { v.gate = false; v.ampEnv.Gate(false); v.filterEnv.Gate(false) }
func (v *Voice) Render(buf *dsp.StereoBuffer) {
	mono := dsp.NewMonoBuffer(dsp.BlockSize); v.osc.Process(mono)
	ampBuf := dsp.NewMonoBuffer(dsp.BlockSize); v.ampEnv.Process(ampBuf)
	for i := range mono { mono[i] = dsp.Sample(float64(mono[i]) * float64(ampBuf[i]) * v.velocity) }
	if v.ampEnv.IsActive() { for i := range buf.Left { buf.Left[i] += mono[i]; buf.Right[i] += mono[i] } }
}
func (v *Voice) IsActive() bool { return v.ampEnv.IsActive() }
func (v *Voice) Note() int { return v.note }
