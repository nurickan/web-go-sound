package synth
import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
func TestVoiceAllocatorStealing(t *testing.T) {
	a := NewVoiceAllocator(4, dsp.SampleRate); p := patch.DefaultPatch()
	for i := 0; i < 8; i++ { v := a.Allocate(); v.NoteOn(60+i, 0.8, p) }
	if n := len(a.ActiveVoices()); n != 4 { t.Errorf("expected 4, got %d", n) }
}
