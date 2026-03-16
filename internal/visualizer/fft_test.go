package visualizer

import (
	"fmt"
	"math"
	"testing"
)

func generateSine(freq, sampleRate float64, numSamples int) []float64 {
	samples := make([]float64, numSamples)
	for i := range samples {
		samples[i] = math.Sin(2 * math.Pi * freq * float64(i) / sampleRate)
	}
	return samples
}

func TestVisualizerPipeline(t *testing.T) {
	const (
		sampleRate = 44100.0
		freq       = 440.0 // A4 — should spike in a mid bar
		numSamples = 1024
	)

	samples := generateSine(freq, sampleRate, numSamples)
	windowed := applyHann(samples)
	mags := computeFFTMags(windowed)
	bins := getBins(mags)

	expectedBin := int(math.Round(freq * numSamples / sampleRate))
	fmt.Printf("Expected spike near bin index: %d\n", expectedBin)
	fmt.Printf("Bins: %v\n", bins)

	maxVal := 0.0
	maxIdx := 0
	for i, v := range bins {
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}
	fmt.Printf("Highest bar: index %d, value %.4f\n", maxIdx, maxVal)
}
