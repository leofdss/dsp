package file

import (
	"encoding/json"
	"fmt"
	"os"

	"dsp/internal/domain"
)

type ConfigLoader struct{}

func (ConfigLoader) Load(path string) (domain.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return domain.Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg domain.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return domain.Config{}, fmt.Errorf("decode config: %w", err)
	}

	return cfg, nil
}
