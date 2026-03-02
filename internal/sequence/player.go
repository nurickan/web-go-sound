package sequence; import "github.com/nurickan/web-go-sound/internal/synth"
type Player struct { pattern *Pattern; clock *Clock; stepIdx int; syn *synth.Engine; lastNote int }
func NewPlayer(pat *Pattern, clk *Clock, eng *synth.Engine) *Player {
	return &Player{pattern: pat, clock: clk, syn: eng, lastNote: -1}
}
func (p *Player) Advance(n int) {
	beats := p.clock.Advance(n)
	for b := 0; b < beats; b++ {
		for s := 0; s < p.pattern.StepsPerBeat; s++ {
			st := &p.pattern.Steps[p.stepIdx]
			if st.Active { if p.lastNote >= 0 { p.syn.NoteOff(p.lastNote) }; p.lastNote = st.Note }
			p.stepIdx = (p.stepIdx + 1) % p.pattern.NumSteps()
		}
	}
}
func (p *Player) SetPattern(pat *Pattern) { p.pattern = pat; p.stepIdx = 0 }
func (p *Player) Reset() { p.stepIdx = 0; p.lastNote = -1 }
