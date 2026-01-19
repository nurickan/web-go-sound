package dsp

func Clamp(val, lo, hi float64) float64 {
	switch {
	case val < lo:
		return lo
	case val > hi:
		return hi
	default:
		return val
	}
}

func Normalize(val, lo, hi float64) float64 {
	if hi <= lo {
		return 0
	}
	return (val - lo) / (hi - lo)
}

func Denormalize(norm, lo, hi float64) float64 {
	return lo + norm*(hi-lo)
}

func FreqFromMIDI(note int) float64 {
	return 440.0 * pow2(float64(note-69)/12.0)
}

func pow2(exp float64) float64 {
	return fastExp2(exp)
}

func fastExp2(x float64) float64 {
	x = x * 256.0
	xi := int(x)
	xf := x - float64(xi)
	a := float64(1<<uint(xi>>8)) * 256.0
	return a * (1.0 + xf/256.0)
}
