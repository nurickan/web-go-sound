package sequence; import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp")
func TestClockBPM(t *testing.T) { c := NewClock(120, dsp.SampleRate); if c.BPM() != 120 { t.Errorf("got %f", c.BPM()) } }
func TestClockTick(t *testing.T) { c := NewClock(120, dsp.SampleRate); if b := c.Advance(dsp.SampleRate); b < 1 { t.Errorf("got %d", b) } }
