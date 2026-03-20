package audio

import (
	"fmt"
	"io"

	"github.com/valentino7504/sona/internal/audio/decoding"
	"github.com/valentino7504/sona/internal/visualizer"
)

// AudioPlayer is the interface exposed to the frontend for basic audio controlling.
type AudioPlayer interface {
	Play()  // plays the requested audio
	Pause() // pauses the audio playing
	Close()
}

// NewAudioPlayer accepts filename and an optional format parameter for creating an audio player
// backend for use and a visualizer data struct for use by the visualizer. For now it defaults to
// oto for player. Will extend later.
func NewAudioPlayer(fileName string, fileFormat string) (AudioPlayer, *visualizer.Visualizer, error) {
	fileFmt := fileFormat
	if fileFmt == "" {
		fileFmt = extractFileFormat(fileName)
	}
	decoder, err := decoding.NewAudioDecoder(fileFmt)
	if err != nil {
		return nil, nil, fmt.Errorf("create audio decoder: %w", err)
	}
	decodedAudio, err := decoder.Decode(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding: %w", err)
	}
	ctxOpts := newContextOptions(
		withSampleRate(decodedAudio.SampleRate),
		withChannelCount(decodedAudio.Channels),
		withBitDepth(decodedAudio.BitDepth),
	)
	pipeReader, pipeWriter := io.Pipe()
	otoBackend, err := NewOtoBackend(io.TeeReader(decodedAudio.Data, pipeWriter), ctxOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize oto backend: %w", err)
	}
	return otoBackend, visualizer.NewVisualizer(pipeReader, decodedAudio), nil
}
