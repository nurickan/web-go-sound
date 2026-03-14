package synth; import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
func TestPatchDrivenEngine(t *testing.T) {
	e := NewEngine(dsp.SampleRate, 8); p := patch.DefaultPatch(); p.Oscillators[0].Waveform = 1
	p.Filter.Enabled = true; p.Filter.Cutoff = 2000; p.Filter.Resonance = 0.4
	e.NoteOn(60, 0.9, p); e.NoteOn(64, 0.7, p); e.NoteOn(67, 0.5, p); n := 0
	for i := 0; i < 10; i++ {
		b := dsp.NewStereoBuffer(dsp.BlockSize); e.Render(&b); m := ComputeMeters(&b)
		if m.PeakL > 0 { n++ }
	}; if n == 0 { t.Error("no signal") }; e.AllNotesOff()
}
