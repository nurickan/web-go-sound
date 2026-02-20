package synth
import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
func TestVoiceGateLifecycle(t *testing.T) {
	v := NewVoice(dsp.SampleRate); p := patch.DefaultPatch(); v.NoteOn(69, 0.8, p)
	if !v.IsActive() { t.Error("not active after NoteOn") }; v.NoteOff()
	buf := dsp.NewStereoBuffer(dsp.BlockSize); v.Render(&buf)
}
