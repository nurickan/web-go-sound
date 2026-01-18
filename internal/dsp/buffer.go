package dsp

type MonoBuffer []Sample

type StereoBuffer struct {
	Left  MonoBuffer
	Right MonoBuffer
}

func NewMonoBuffer(size int) MonoBuffer {
	return make(MonoBuffer, size)
}

func NewStereoBuffer(size int) StereoBuffer {
	return StereoBuffer{
		Left:  NewMonoBuffer(size),
		Right: NewMonoBuffer(size),
	}
}

func (b MonoBuffer) Clear() {
	for i := range b {
		b[i] = 0
	}
}

func (b *StereoBuffer) Clear() {
	b.Left.Clear()
	b.Right.Clear()
}
