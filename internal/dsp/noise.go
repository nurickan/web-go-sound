package dsp

type Noise struct {
	state uint32
}

func NewNoise(seed uint32) *Noise {
	return &Noise{state: seed}
}

func (n *Noise) Process(buf MonoBuffer) {
	for i := range buf {
		buf[i] = Sample(n.next())*2.0 - 1.0
	}
}

func (n *Noise) next() float64 {
	n.state ^= n.state << 13
	n.state ^= n.state >> 17
	n.state ^= n.state << 5
	return float64(n.state%1000000) / 1000000.0
}

func (n *Noise) Reset() {
	n.state = 12345
}
