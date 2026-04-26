package domain

type Processor interface {
	ProcessInPlace(buf []float32)
}
