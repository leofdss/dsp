package app

import (
	"context"

	"dsp/internal/domain"
	"dsp/internal/ports"
)

type Engine struct {
	stream    ports.AudioStream
	processor domain.Processor
}

func NewEngine(stream ports.AudioStream, processor domain.Processor) Engine {
	return Engine{
		stream:    stream,
		processor: processor,
	}
}

func (e Engine) Run(ctx context.Context) error {
	return e.stream.Start(ctx, e.processor.ProcessInPlace)
}
