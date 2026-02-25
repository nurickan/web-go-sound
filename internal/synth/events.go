package synth
type EventType int; const (EventNoteOn EventType = iota; EventNoteOff; EventParamChange)
type TimestampedEvent struct { Type EventType; Note int; Velocity float64; ParamID int; ParamVal float64; BlockPos int }
type EventQueue struct { events []TimestampedEvent }
func NewEventQueue() *EventQueue { return &EventQueue{} }
func (q *EventQueue) Push(e TimestampedEvent) { q.events = append(q.events, e) }
func (q *EventQueue) Flush() []TimestampedEvent { e := q.events; q.events = nil; return e }
func (q *EventQueue) Len() int { return len(q.events) }
