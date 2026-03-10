package render; import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/synth")
func TestOfflineLen(t *testing.T) {
	r := NewOfflineRenderer(synth.NewEngine(dsp.SampleRate, 8)); buf := r.Render(0.1)
	if len(buf.Left) != dsp.SecondsToSamples(0.1) { t.Error("bad length") }
}
func TestOfflineChannels(t *testing.T) {
	r := NewOfflineRenderer(synth.NewEngine(dsp.SampleRate, 8)); buf := r.Render(0.1)
	if len(buf.Left) != len(buf.Right) { t.Error("channels differ") }
}
