package ports

import "context"

type AudioProcessFunc func([]float32)

type AudioStream interface {
	Start(ctx context.Context, process AudioProcessFunc) error
}
