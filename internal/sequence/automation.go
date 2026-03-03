package sequence
type AutomationPoint struct { Position, Value float64 }
type AutomationLane struct { Target int; Points []AutomationPoint; Interpolation string }
func NewAutomationLane(target int, interp string) AutomationLane { return AutomationLane{Target: target, Points: []AutomationPoint{}, Interpolation: interp} }
func (l *AutomationLane) AddPoint(pos, val float64) { l.Points = append(l.Points, AutomationPoint{Position: pos, Value: val}) }
func (l *AutomationLane) ValueAt(pos float64) float64 {
	if len(l.Points) == 0 { return 0 }; if len(l.Points) == 1 { return l.Points[0].Value }
	for i := 1; i < len(l.Points); i++ {
		if l.Points[i].Position >= pos {
			p, n := l.Points[i-1], l.Points[i]; t := (pos - p.Position) / (n.Position - p.Position); return p.Value + t*(n.Value-p.Value)
		}
	}; return l.Points[len(l.Points)-1].Value
}
