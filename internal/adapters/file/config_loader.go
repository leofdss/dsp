package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"dsp/internal/domain"

	"gopkg.in/yaml.v3"
)

type ConfigLoader struct{}

func (ConfigLoader) Load(path string) (domain.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return domain.Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg domain.Config
	if err := unmarshalConfig(path, data, &cfg); err != nil {
		return domain.Config{}, fmt.Errorf("decode config: %w", err)
	}

	return cfg, nil
}

func unmarshalConfig(path string, data []byte, cfg *domain.Config) error {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, cfg)
	default:
		return json.Unmarshal(data, cfg)
	}
}
