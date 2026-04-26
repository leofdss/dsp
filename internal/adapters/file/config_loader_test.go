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
