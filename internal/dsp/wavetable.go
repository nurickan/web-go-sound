package dsp

const TableSize = 2048

type Wavetable struct {
	data [TableSize]Sample
}

func NewSineTable() *Wavetable {
	wt := &Wavetable{}
	for i := range wt.data {
		wt.data[i] = Sample(sin2pi(float64(i) / float64(TableSize)))
	}
	return wt
}

func NewSawTable() *Wavetable {
	wt := &Wavetable{}
	for i := range wt.data {
		wt.data[i] = Sample(2.0*float64(i)/float64(TableSize) - 1.0)
	}
	return wt
}

func NewSquareTable() *Wavetable {
	wt := &Wavetable{}
	half := TableSize / 2
	for i := range wt.data {
		if i < half {
			wt.data[i] = 1.0
		} else {
			wt.data[i] = -1.0
		}
	}
	return wt
}

func NewTriangleTable() *Wavetable {
	wt := &Wavetable{}
	half := TableSize / 2
	for i := range wt.data {
		if i < half {
			wt.data[i] = Sample(4.0*float64(i)/float64(TableSize) - 1.0)
		} else {
			wt.data[i] = Sample(3.0 - 4.0*float64(i)/float64(TableSize))
		}
	}
	return wt
}

func (wt *Wavetable) Read(idx int) Sample {
	return wt.data[idx%TableSize]
}
