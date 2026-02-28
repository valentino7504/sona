// Package decoding implements functions and interfaces for decoding various audio formats
package decoding

import (
	"fmt"
	"io"
	"strings"
)

type BitDepth int

const (
	Format16BitInt = iota
	Format32BitFloat
	FormatUnsigned8BitInt
)

// AudioDecoder defines an interface for decoding various audio formats
type AudioDecoder interface {
	// Decode converts the digital bitstream/format to [io.Reader] for the oto backend.
	Decode(fileName string) (*DecodedAudio, error)
}

type DecodedAudio struct {
	SampleRate int
	BitDepth   BitDepth
	Channels   int
	Data       io.Reader
}

// NewAudioDecoder creates a new [AudioDecoder] based on the format of the audio input.
//
// It takes the format as a string parameter and returns an [AudioDecoder] or an error
// for unrecognised and unsuppported format types.
func NewAudioDecoder(fileFormat string) (AudioDecoder, error) {
	var decoder AudioDecoder
	switch strings.ToLower(fileFormat) {
	case "mp3":
		decoder = &Mp3Decoder{}
	case "":
		return nil, fmt.Errorf("no audio format provided")
	default:
		return nil, fmt.Errorf("unrecognised/unsupported audio format: %s", fileFormat)
	}
	return decoder, nil
}
