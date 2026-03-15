package visualizer

import (
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/valentino7504/sona/internal/audio/decoding"
)

type VisualizerData struct {
	PcmReader  io.Reader
	SampleRate int
	BitDepth   decoding.BitDepth
	Channels   int
}

// NewVisualizerData creates a new [visualizer.VisualizerData] based on the decoded audio
//
// It takes in the pcm reader to be assigned to the new visualizer and the decoded audio.
func NewVisualizerData(pcmReader io.Reader, da *decoding.DecodedAudio) *VisualizerData {
	return &VisualizerData{
		SampleRate: da.SampleRate,
		BitDepth:   da.BitDepth,
		Channels:   da.Channels,
		PcmReader:  pcmReader,
	}
}

// byteDepth just converts the number of bits in the bit depth to bytes
func (data *VisualizerData) byteDepth() int {
	switch data.BitDepth {
	case decoding.Format16BitInt:
		return 2
	case decoding.Format32BitFloat:
		return 4
	case decoding.FormatUnsigned8BitInt:
		return 1
	default:
		panic("cannot determine bit depth")
	}
}

// readFrame reads a portion of pcm data for the fast fourier transform.
//
// the size of this portion is determined by the number of audio channels, the
// byte depth of the pcm data and the size of a frame of pcm data, usually 1024.
func (data *VisualizerData) readFrame(frameSize int) ([]byte, error) {
	readBytes := make([]byte, data.Channels*frameSize*data.byteDepth())
	if _, err := io.ReadFull(data.PcmReader, readBytes); err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return readBytes, err
		}
		return nil, err
	}
	return readBytes, nil
}

// bytesToSamples converts the pcm data to float64 data.
//
// This is needed as FFT can only be performed on 64 bit float numbers.
func (data *VisualizerData) bytesToSamples(raw []byte) []float64 {
	bd := data.byteDepth()
	bytesPerStep := bd * data.Channels
	output := make([]float64, len(raw)/bytesPerStep)
	for i := 0; i < len(raw); i += bytesPerStep {
		var sample float64 = 0
		for j := 0; j < data.Channels; j++ {
			channelBytes := raw[i+j*bd : i+j*bd+bd]
			switch data.BitDepth {
			case decoding.Format16BitInt:
				signed := int16(binary.LittleEndian.Uint16(channelBytes))
				sample += float64(signed) / 32768.0
			case decoding.Format32BitFloat:
				uint32Val := binary.LittleEndian.Uint32(channelBytes)
				sample += float64(math.Float32frombits(uint32Val))
			case decoding.FormatUnsigned8BitInt:
				sample += (float64(channelBytes[0]) - 128) / 128
			}
		}
		sample /= float64(data.Channels)
		output[i/bytesPerStep] = sample
	}
	return output
}
