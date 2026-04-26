package domain

import "testing"

func TestGainProcessInPlace(t *testing.T) {
	processor := NewGain(2)
	buf := []float32{0.5, -1, 2}
	original := &buf[0]

	processor.ProcessInPlace(buf)

	if &buf[0] != original {
		t.Fatal("expected in-place processing")
	}

	want := []float32{1, -2, 4}
	for i := range want {
		if buf[i] != want[i] {
			t.Fatalf("sample %d: got %v want %v", i, buf[i], want[i])
		}
	}
}

func TestGainZeroLengthBuffer(t *testing.T) {
	processor := NewGain(2)
	buf := []float32{}

	processor.ProcessInPlace(buf)
}
