package synth
import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
func TestEngineRendersNonSilent(t *testing.T) {
	e := NewEngine(dsp.SampleRate, 8); p := patch.DefaultPatch(); e.NoteOn(69, 0.8, p)
	buf := dsp.NewStereoBuffer(dsp.BlockSize); e.Render(&buf); ok := false
	for i := range buf.Left { if buf.Left[i] != 0 || buf.Right[i] != 0 { ok = true; break } }
	if !ok { t.Error("silence only") }; e.AllNotesOff()
}
