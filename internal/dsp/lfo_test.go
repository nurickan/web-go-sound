package dsp
import "testing"
func TestLFOBounds(t *testing.T) {
	shapes := []LFOShape{LFOSine, LFOSaw, LFOSquare, LFOTriangle}
	for _, s := range shapes {
		l := NewLFO(SampleRate); l.SetRate(1.0); l.SetShape(s)
		buf := NewMonoBuffer(BlockSize * 4); l.Process(buf)
		for i, v := range buf {
			if v < -1.0 || v > 1.0 { t.Errorf("shape %d sample %d out: %f", s, i, v) }
		}
	}
}
