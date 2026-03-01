package audio

import (
	"fmt"

	"github.com/valentino7504/sona/internal/audio/decoding"
)

// AudioPlayer is the interface exposed to the frontend for basic audio controlling.
type AudioPlayer interface {
	Play() // plays the requested audio
	// TODO - Pause(), Seek(time), Stop()
}

// NewAudioPlayer accepts filename and an optional format parameter for creating an audio player
// backend for use. For now it defaults to oto. Will extend later.
func NewAudioPlayer(fileName string, fileFormat string) (AudioPlayer, error) {
	fileFmt := fileFormat
	if fileFmt == "" {
		fileFmt = extractFileFormat(fileName)
	}
	decoder, err := decoding.NewAudioDecoder(fileFmt)
	if err != nil {
		return nil, fmt.Errorf("create audio decoder: %w", err)
	}
	decodedAudio, err := decoder.Decode(fileName)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}
	ctxOpts := newContextOptions(
		withSampleRate(decodedAudio.SampleRate),
		withChannelCount(decodedAudio.Channels),
		withBitDepth(decodedAudio.BitDepth),
	)
	otoBackend, err := NewOtoBackend(decodedAudio.Data, ctxOpts)
	if err != nil {
		return nil, fmt.Errorf("initialize oto backend: %w", err)
	}
	return otoBackend, nil
}
