package decoding

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

// Mp3Decoder is a decoder for mp3 format audio files with a pointer Decode method,
// implementing [AudioDecoder]
type Mp3Decoder struct{}

// Decode decodes mp3 format audio files.
//
// It takes the file name as a string parameter and returns an [io.Reader] and any error
// encountered while decoding.
func (decoder *Mp3Decoder) Decode(fileName string) (io.Reader, error) {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err.Error())
	}
	bytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(bytesReader)
	if err != nil {
		return nil, fmt.Errorf("failed to decode mp3 file %s: %v", fileName, err.Error())
	}
	return decodedMp3, nil
}
