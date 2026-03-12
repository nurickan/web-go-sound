package render; import ("bytes"; "testing"; "github.com/nurickan/web-go-sound/internal/dsp")
func TestWAVHeader(t *testing.T) {
	enc := NewWAVEncoder(48000); buf := dsp.NewStereoBuffer(256); var b bytes.Buffer
	if err := enc.Encode(&b, buf); err != nil { t.Fatal(err) }
	if b.Len() < 44 { t.Error("too small") }; if string(b.Bytes()[:4]) != "RIFF" { t.Error("no RIFF") }
}
func TestWAVSamples(t *testing.T) {
	enc := NewWAVEncoder(48000); buf := dsp.NewStereoBuffer(128)
	for i := range buf.Left { buf.Left[i] = 0.5; buf.Right[i] = -0.5 }
	var b bytes.Buffer; if err := enc.Encode(&b, buf); err != nil { t.Fatal(err) }
}
