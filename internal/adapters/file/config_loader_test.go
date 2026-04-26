package file

import (
	"os"
	"path/filepath"
	"testing"
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
