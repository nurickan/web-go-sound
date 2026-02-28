package sequence
type Step struct { Active bool; Note int; Velocity, Gate float64 }
type Pattern struct { Name string; Steps []Step; StepsPerBeat int }
func NewEmptyPattern(name string, n int) Pattern {
	s := make([]Step, n); for i := range s { s[i] = Step{Active: false, Note: 60, Velocity: 0.8, Gate: 0.5} }
	return Pattern{Name: name, Steps: s, StepsPerBeat: 4}
}
func (p *Pattern) SetStep(i int, s Step) { if i >= 0 && i < len(p.Steps) { p.Steps[i] = s } }
func (p *Pattern) NumSteps() int { return len(p.Steps) }
