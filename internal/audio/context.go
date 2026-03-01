package audio

import "github.com/valentino7504/sona/internal/audio/decoding"

// contextOptions are used to define the parameters for the oto Context,
// ensuring DAC is done properly.
type contextOptions struct {
	SampleRate   int
	BitDepth     decoding.BitDepth
	ChannelCount int
}

// withSampleRate provides option for customising the sample rate, affecting how fast the audio plays.
// Usually around 44100 or 48000
func withSampleRate(rate int) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.SampleRate = rate
	}
}

// withChannelCount provides option for customising the output channels - stereo or mono.
func withChannelCount(count int) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.ChannelCount = count
	}
}

// withBitDepth determines the bit depth for customising the number of bytes per sample.
func withBitDepth(format decoding.BitDepth) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.BitDepth = format
	}
}

// newContextOptions takes in a variable number of the option functions and generates a new contextOptions
// with the specified values or the sane defaults.
func newContextOptions(customOpts ...func(*contextOptions)) contextOptions {
	opts := &contextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		BitDepth:     decoding.Format16BitInt,
	}
	for _, customOpt := range customOpts {
		customOpt(opts)
	}
	return *opts
}
