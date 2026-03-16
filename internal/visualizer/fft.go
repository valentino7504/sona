package visualizer

import (
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
)

const NO_BARS = 16

// applyHann applies the Hann smoothing function to an array of samples to reduce
// spectral leakage and ensure that the visualizer does not jitter. frameSize is
// always > 1 so it can never result in a zero division error.
func applyHann(samples []float64) []float64 {
	n := len(samples)
	windowed := make([]float64, n)
	for i, sample := range samples {
		wi := 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(n-1)))
		windowed[i] = wi * sample
	}
	return windowed
}

// computeFFTMags computes magnitude of fft for an array of samples. Should be used after hann
// smoothing and overlapping (probably when I am able to get this working).
// It returns the magnitudes of the FFT values, since that is all I need for visualization.
func computeFFTMags(samples []float64) []float64 {
	n := len(samples)
	fft := fourier.NewFFT(n)
	coeffs := fft.Coefficients(nil, samples)
	mags := make([]float64, len(coeffs))
	for i, coeff := range coeffs {
		mags[i] = cmplx.Abs(coeff)
	}
	return mags
}

// getBins is used to group the magnitude values into bins for the visualizer. I
// am only using 16 bars to start out with.
func getBins(vals []float64) []float64 {
	boundaries := make([]float64, NO_BARS+1)
	start := 1.0
	end := float64(len(vals) - 1)
	ratio := math.Pow((end / start), (1.0 / float64(NO_BARS)))
	for i := range boundaries {
		boundaries[i] = start * math.Pow(ratio, float64(i))
	}
	bins := make([]float64, NO_BARS)
	for i := range bins {
		binStart, binEnd := int(math.Floor(boundaries[i])), int(math.Floor(boundaries[i+1]))
		sum := 0.0
		for _, val := range vals[binStart:binEnd] {
			sum += val
		}
		if binEnd != binStart {
			bins[i] = sum / float64(binEnd-binStart)
		} else {
			bins[i] = 0
		}
	}
	return bins
}
