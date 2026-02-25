package synth; import "testing"
func TestEventOrdering(t *testing.T) {
	q := NewEventQueue(); q.Push(TimestampedEvent{Type: EventNoteOn, Note: 60, BlockPos: 10})
	q.Push(TimestampedEvent{Type: EventNoteOff, Note: 60, BlockPos: 20}); e := q.Flush()
	if len(e) != 2 { t.Errorf("got %d", len(e)) }
}
func TestEventEmpty(t *testing.T) { q := NewEventQueue(); if q.Len() != 0 { t.Error("not empty") } }
