package fake

import (
	"context"

	"dsp/internal/ports"
)

type Stream struct {
	buffer []float32
}

func NewStream(buffer []float32) Stream {
	return Stream{buffer: buffer}
}

func (s Stream) Start(_ context.Context, process ports.AudioProcessFunc) error {
	process(s.buffer)
	return nil
}

func (s Stream) Buffer() []float32 {
	return s.buffer
}
