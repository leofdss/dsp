package file

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"dsp/internal/domain"
)

func TestConfigLoaderLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")

	if err := os.WriteFile(path, []byte(`{"gain":1.5}`), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	loader := ConfigLoader{}
	cfg, err := loader.Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.Gain != 1.5 {
		t.Fatalf("got gain %v want %v", cfg.Gain, float32(1.5))
	}
}

func TestConfigLoaderLoadYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	if err := os.WriteFile(path, []byte("gain: 2.25\n"), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	loader := ConfigLoader{}
	cfg, err := loader.Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.Gain != 2.25 {
		t.Fatalf("got gain %v want %v", cfg.Gain, float32(2.25))
	}
}

func TestConfigLoaderLoadInvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yml")

	if err := os.WriteFile(path, []byte("gain: [\n"), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	loader := ConfigLoader{}
	if _, err := loader.Load(path); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestConfigLoaderLoadReadError(t *testing.T) {
	loader := ConfigLoader{}

	if _, err := loader.Load(filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Fatal("expected read error")
	}
}

func TestUnmarshalConfigYML(t *testing.T) {
	var cfg domain.Config
	err := unmarshalConfig("config.yml", []byte("gain: 3.5\n"), &cfg)
	if err != nil {
		t.Fatalf("unmarshal config: %v", err)
	}

	if cfg.Gain != 3.5 {
		t.Fatalf("got gain %v want %v", cfg.Gain, float32(3.5))
	}
}

func TestUnmarshalConfigDefaultJSON(t *testing.T) {
	var cfg domain.Config
	err := unmarshalConfig("config.conf", []byte(`{"gain":4.5}`), &cfg)
	if err != nil {
		t.Fatalf("unmarshal config: %v", err)
	}

	if cfg.Gain != 4.5 {
		t.Fatalf("got gain %v want %v", cfg.Gain, float32(4.5))
	}
}

func TestUnmarshalConfigInvalidJSON(t *testing.T) {
	var cfg domain.Config
	err := unmarshalConfig("config.json", []byte(`{"gain":`), &cfg)
	if err == nil {
		t.Fatal("expected decode error")
	}

	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Fatalf("expected concrete error, got %T", err)
	}
}
