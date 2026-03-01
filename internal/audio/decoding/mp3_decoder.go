package decoding

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

// Mp3Decoder is a decoder for mp3 format audio files with a pointer Decode method,
// implementing [AudioDecoder]
type Mp3Decoder struct{}

// Decode decodes mp3 format audio files.
//
// It takes the file name as a string parameter and returns a pointer to a [DecodedAudio]
// and any error encountered while decoding.
func (decoder *Mp3Decoder) Decode(fileName string) (*DecodedAudio, error) {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	bytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(bytesReader)
	if err != nil {
		return nil, fmt.Errorf("decode mp3: %w", err)
	}
	return &DecodedAudio{
		SampleRate: decodedMp3.SampleRate(),
		BitDepth:   Format16BitInt,
		Channels:   2,
		Data:       decodedMp3,
	}, nil
}
