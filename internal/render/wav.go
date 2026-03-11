package render; import ("encoding/binary"; "io"; "github.com/nurickan/web-go-sound/internal/dsp")
type WAVEncoder struct { sampleRate, bitDepth int }
func NewWAVEncoder(sr int) *WAVEncoder { return &WAVEncoder{sampleRate: sr, bitDepth: 16} }
func (e *WAVEncoder) Encode(w io.Writer, buf dsp.StereoBuffer) error {
	n := len(buf.Left); dataSize := n * 2 * 2; h := make([]byte, 44)
	copy(h[0:4], []byte("RIFF")); binary.LittleEndian.PutUint32(h[4:8], uint32(44+dataSize-8))
	copy(h[8:12], []byte("WAVE")); copy(h[12:16], []byte("fmt "))
	binary.LittleEndian.PutUint32(h[16:20], 16); binary.LittleEndian.PutUint16(h[20:22], 1)
	binary.LittleEndian.PutUint16(h[22:24], 2); binary.LittleEndian.PutUint32(h[24:28], uint32(e.sampleRate))
	binary.LittleEndian.PutUint32(h[28:32], uint32(e.sampleRate*2*2))
	binary.LittleEndian.PutUint16(h[32:34], 4); binary.LittleEndian.PutUint16(h[34:36], 16)
	copy(h[36:40], []byte("data")); binary.LittleEndian.PutUint32(h[40:44], uint32(dataSize))
	if _, err := w.Write(h); err != nil { return err }
	s := make([]int16, n*2)
	for i := 0; i < n; i++ { s[i*2] = int16(dsp.ClampSample(buf.Left[i]) * 32767); s[i*2+1] = int16(dsp.ClampSample(buf.Right[i]) * 32767) }
	return binary.Write(w, binary.LittleEndian, s)
}
