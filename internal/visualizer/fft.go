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

// computeFFT computes fft for an array of samples. should be used after hann
// smoothing and overlapping (probably when I am able to get this working.
func computeFFT(samples []float64) []complex128 {
	n := len(samples)
	fft := fourier.NewFFT(n)
	coeff := fft.Coefficients(nil, samples)
	return coeff
}

// getMagnitudes calculates the magnitude of the values of the FFT transform.
func getMagnitudes(coeffs []complex128) []float64 {
	mags := make([]float64, len(coeffs))
	for i, coeff := range coeffs {
		mags[i] = cmplx.Abs(coeff)
	}
	return mags
}
