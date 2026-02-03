package dsp
type ModSource int32; type ModDest int32
type ModRoute struct { Source ModSource; Dest ModDest; Amount float64 }
type ModMatrix struct { sources map[ModSource]Sample; routes []ModRoute }
func NewModMatrix() *ModMatrix { return &ModMatrix{sources: make(map[ModSource]Sample)} }
func (m *ModMatrix) SetSource(src ModSource, val Sample) { m.sources[src] = val }
func (m *ModMatrix) AddRoute(route ModRoute) { m.routes = append(m.routes, route) }
func (m *ModMatrix) ClearRoutes() { m.routes = m.routes[:0] }
func (m *ModMatrix) GetModulation(dest ModDest) Sample {
	var sum Sample
	for _, r := range m.routes { if r.Dest == dest { sum += m.sources[r.Source] * Sample(r.Amount) } }
	return ClampSample(sum)
}
func ClampSample(s Sample) Sample { if s < -1.0 { return -1.0 }; if s > 1.0 { return 1.0 }; return s }
