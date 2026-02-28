package audio

import "github.com/valentino7504/sona/internal/audio/decoding"

type contextOptions struct {
	SampleRate   int
	BitDepth     decoding.BitDepth
	ChannelCount int
}

func withSampleRate(rate int) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.SampleRate = rate
	}
}

func withChannelCount(count int) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.ChannelCount = count
	}
}

func withBitDepth(format decoding.BitDepth) func(*contextOptions) {
	return func(opts *contextOptions) {
		opts.BitDepth = format
	}
}

func newContextOptions(customOpts ...func(*contextOptions)) contextOptions {
	opts := &contextOptions{
		SampleRate:   48000,
		ChannelCount: 2,
		BitDepth:     decoding.Format16BitInt,
	}
	for _, customOpt := range customOpts {
		customOpt(opts)
	}
	return *opts
}
