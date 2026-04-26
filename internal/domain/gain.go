package domain

type Gain struct {
	factor float32
}

func NewGain(factor float32) Gain {
	return Gain{factor: factor}
}

func (g Gain) ProcessInPlace(buf []float32) {
	for i := range buf {
		buf[i] *= g.factor
	}
}
