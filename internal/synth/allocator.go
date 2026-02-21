package synth
type VoiceAllocator struct { voices []*Voice; numVoices, nextVoice int }
func NewVoiceAllocator(num int, sr float64) *VoiceAllocator {
	v := make([]*Voice, num); for i := range v { v[i] = NewVoice(sr) }
	return &VoiceAllocator{voices: v, numVoices: num}
}
func (a *VoiceAllocator) Allocate() *Voice {
	for i := 0; i < a.numVoices; i++ {
		idx := (a.nextVoice + i) % a.numVoices
		if !a.voices[idx].IsActive() { a.nextVoice = (idx + 1) % a.numVoices; return a.voices[idx] }
	}
	stolen := a.voices[a.nextVoice]; stolen.NoteOff()
	a.nextVoice = (a.nextVoice + 1) % a.numVoices; return stolen
}
func (a *VoiceAllocator) ActiveVoices() []*Voice {
	var act []*Voice; for _, v := range a.voices { if v.IsActive() { act = append(act, v) } }; return act
}
func (a *VoiceAllocator) All() []*Voice { return a.voices }
