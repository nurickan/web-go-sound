package synth; import ("math"; "github.com/nurickan/web-go-sound/internal/dsp")
type Meters struct { PeakL, PeakR, RMSL, RMSR float64 }
func ComputeMeters(buf *dsp.StereoBuffer) Meters {
	var sl, sr, pl, pr float64
	for i := range buf.Left {
		l, r := float64(buf.Left[i]), float64(buf.Right[i]); sl += l*l; sr += r*r
		if al := math.Abs(l); al > pl { pl = al }; if ar := math.Abs(r); ar > pr { pr = ar }
	}; n := float64(len(buf.Left))
	return Meters{PeakL: pl, PeakR: pr, RMSL: math.Sqrt(sl/n), RMSR: math.Sqrt(sr/n)}
}
