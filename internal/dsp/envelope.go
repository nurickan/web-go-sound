package dsp

type ADSRPhase int

const (
	PhaseIdle ADSRPhase = iota
	PhaseAttack
	PhaseDecay
	PhaseSustain
	PhaseRelease
)

type ADSR struct {
	AttackTime   float64
	DecayTime    float64
	SustainLevel float64
	ReleaseTime  float64
	phase        ADSRPhase
	level        float64
	samples      int
	phasePos     int
	sr           float64
}

func NewADSR(sr float64) *ADSR {
	return &ADSR{sr: sr}
}

func (e *ADSR) Gate(on bool) {
	if on {
		e.phase = PhaseAttack
		e.phasePos = 0
		e.samples = SecondsToSamples(e.AttackTime)
		if e.samples < 1 { e.samples = 1 }
	} else if e.phase != PhaseIdle {
		e.phase = PhaseRelease
		e.phasePos = 0
		e.samples = SecondsToSamples(e.ReleaseTime)
		if e.samples < 1 { e.samples = 1 }
	}
}

func (e *ADSR) Process(buf MonoBuffer) {
	for i := range buf {
		buf[i] = e.tick()
	}
}

func (e *ADSR) tick() Sample {
	switch e.phase {
	case PhaseIdle:
		return 0
	case PhaseAttack:
		e.phasePos++
		e.level = float64(e.phasePos) / float64(e.samples)
		if e.phasePos >= e.samples {
			e.phase = PhaseDecay
			e.phasePos = 0
			e.samples = SecondsToSamples(e.DecayTime)
			if e.samples < 1 { e.samples = 1 }
		}
	case PhaseDecay:
		e.phasePos++
		t := float64(e.phasePos) / float64(e.samples)
		e.level = 1.0 - (1.0-e.SustainLevel)*t
		if e.phasePos >= e.samples {
			e.phase = PhaseSustain
			e.level = e.SustainLevel
		}
	case PhaseSustain:
		e.level = e.SustainLevel
	case PhaseRelease:
		e.phasePos++
		t := float64(e.phasePos) / float64(e.samples)
		e.level = e.level * (1.0 - t)
		if e.level < 0.001 || e.phasePos >= e.samples {
			e.phase = PhaseIdle
			e.level = 0
		}
	}
	if e.level < 0 { e.level = 0 }
	return Sample(e.level)
}

func (e *ADSR) IsActive() bool { return e.phase != PhaseIdle }

func (e *ADSR) Reset() {
	e.phase = PhaseIdle
	e.level = 0
	e.phasePos = 0
}
