package ports

import "dsp/internal/domain"

type ConfigLoader interface {
	Load(path string) (domain.Config, error)
}
