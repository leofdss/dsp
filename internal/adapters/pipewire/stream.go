package pipewire

import (
	"context"
	"errors"

	"dsp/internal/ports"
)

var ErrNotImplemented = errors.New("pipewire adapter not implemented")

type Stream struct{}

func (Stream) Start(_ context.Context, _ ports.AudioProcessFunc) error {
	return ErrNotImplemented
}
