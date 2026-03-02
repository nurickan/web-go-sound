package sequence; import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/synth")
func TestPlayerLoop(t *testing.T) {
	pat := NewEmptyPattern("t", 8); pat.Steps[0] = Step{Active: true, Note: 60, Velocity: 0.8, Gate: 0.5}
	clk := NewClock(120, dsp.SampleRate); eng := synth.NewEngine(dsp.SampleRate, 8)
	p := NewPlayer(&pat, clk, eng); p.Advance(int(dsp.SampleRate))
}
