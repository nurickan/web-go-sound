package render; import ("github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/synth")
type OfflineRenderer struct { engine *synth.Engine }
func NewOfflineRenderer(eng *synth.Engine) *OfflineRenderer { return &OfflineRenderer{engine: eng} }
func (r *OfflineRenderer) Render(dur float64) dsp.StereoBuffer {
	total := dsp.SecondsToSamples(dur); out := dsp.NewStereoBuffer(total); block := dsp.NewStereoBuffer(dsp.BlockSize)
	for pos := 0; pos < total; pos += dsp.BlockSize {
		r.engine.Render(&block); rem := total - pos; if rem > dsp.BlockSize { rem = dsp.BlockSize }
		for i := 0; i < rem; i++ { out.Left[pos+i] += block.Left[i]; out.Right[pos+i] += block.Right[i] }
	}; return out
}
