package audio

import (
	"io"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/valentino7504/sona/internal/audio/decoding"
)

type OtoBackend struct {
	context *oto.Context
	player  *oto.Player
}

func (ob *OtoBackend) Play() {
	ob.player.Play()
	for ob.player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}

func convertOptions(ctxOpts contextOptions) *oto.NewContextOptions {
	var otoCtxOpts oto.NewContextOptions
	otoCtxOpts.ChannelCount = ctxOpts.ChannelCount
	otoCtxOpts.SampleRate = ctxOpts.SampleRate
	switch ctxOpts.BitDepth {
	case decoding.Format16BitInt:
		otoCtxOpts.Format = oto.FormatSignedInt16LE
	case decoding.Format32BitFloat:
		otoCtxOpts.Format = oto.FormatFloat32LE
	case decoding.FormatUnsigned8BitInt:
		otoCtxOpts.Format = oto.FormatUnsignedInt8
	default:
		panic("unrecognised or unsupported audio format")
	}
	return &otoCtxOpts
}

func NewOtoBackend(decoded io.Reader, opts contextOptions) (*OtoBackend, error) {
	ctx, ready, err := oto.NewContext(convertOptions(opts))
	if err != nil {
		return nil, err
	}
	<-ready
	player := ctx.NewPlayer(decoded)
	return &OtoBackend{context: ctx, player: player}, nil
}
