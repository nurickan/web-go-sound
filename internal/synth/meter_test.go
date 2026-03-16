package synth; import ("testing"; "github.com/nurickan/web-go-sound/internal/dsp")
func TestMeters(t *testing.T) {
	b := dsp.NewStereoBuffer(dsp.BlockSize)
	for i := range b.Left { b.Left[i] = 0.5; b.Right[i] = 0.3 }
	m := ComputeMeters(&b)
	if m.PeakL < 0.49 || m.PeakL > 0.51 { t.Errorf("peak L: %f", m.PeakL) }
	if m.RMSL < 0.49 || m.RMSL > 0.51 { t.Errorf("RMS L: %f", m.RMSL) }
}
