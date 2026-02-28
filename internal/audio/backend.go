package audio

import "github.com/valentino7504/sona/internal/audio/decoding"

type AudioPlayer interface {
	Play()
}

// NewAudioPlayer accepts filename and an optional format parameter for creating an audio player
// backend for use. For now it defaults to oto. Will extend later.
func NewAudioPlayer(fileName string, fileFormat string) (AudioPlayer, error) {
	var fileFmt string
	if fileFormat != "" {
		fileFmt = fileFormat
	} else {
		fileFmt = extractFileFormat(fileName)
	}
	decoder, err := decoding.NewAudioDecoder(fileFmt)
	if err != nil {
		return nil, err
	}
	decodedAudio, err := decoder.Decode(fileName)
	if err != nil {
		return nil, err
	}
	ctxOpts := newContextOptions(
		withSampleRate(decodedAudio.SampleRate),
		withChannelCount(decodedAudio.Channels),
		withBitDepth(decodedAudio.BitDepth),
	)
	otoBackend, err := NewOtoBackend(decodedAudio.Data, ctxOpts)
	if err != nil {
		return nil, err
	}
	return otoBackend, nil
}
