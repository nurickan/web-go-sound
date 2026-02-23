package synth
import ("github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/patch")
type Engine struct { allocator *VoiceAllocator; sampleRate float64 }
func NewEngine(sr float64, poly int) *Engine { return &Engine{allocator: NewVoiceAllocator(poly, sr)} }
func (e *Engine) NoteOn(note int, vel float64, p patch.Patch) { v := e.allocator.Allocate(); v.NoteOn(note, vel, p) }
func (e *Engine) NoteOff(note int) {
	for _, v := range e.allocator.ActiveVoices() { if v.Note() == note { v.NoteOff(); return } }
}
func (e *Engine) AllNotesOff() { for _, v := range e.allocator.All() { if v.IsActive() { v.NoteOff() } } }
func (e *Engine) Render(buf *dsp.StereoBuffer) {
	buf.Clear()
	for _, v := range e.allocator.ActiveVoices() {
		s := dsp.NewStereoBuffer(dsp.BlockSize); v.Render(&s)
		for i := range buf.Left { buf.Left[i] += s.Left[i]; buf.Right[i] += s.Right[i] }
	}
}
func (e *Engine) ActiveVoiceCount() int { return len(e.allocator.ActiveVoices()) }
