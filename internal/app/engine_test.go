package app

import (
	"context"
	"testing"

	"dsp/internal/adapters/fake"
	"dsp/internal/domain"
)

func TestEngineRunsProcessorAgainstAudioBuffer(t *testing.T) {
	stream := fake.NewStream([]float32{0.5, -1, 2})
	engine := NewEngine(stream, domain.NewGain(2))

	if err := engine.Run(context.Background()); err != nil {
		t.Fatalf("run engine: %v", err)
	}

	got := stream.Buffer()
	want := []float32{1, -2, 4}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("sample %d: got %v want %v", i, got[i], want[i])
		}
	}
}
