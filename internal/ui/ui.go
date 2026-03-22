package ui

import (
	"bufio"
	"context"
	"os"

	"golang.org/x/term"
)

var prevState *term.State

const (
	CTRL_C = 3
	CTRL_D = 4
	SPC    = 32
)

func Start(cancel context.CancelFunc, input chan rune, bins chan []float64) {
	var err error
	prevState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	drawStatic()
	defer Stop()
	reader := bufio.NewReader(os.Stdin)
	var char rune

	go func() {
		for binVals := range bins {
			drawBars(binVals)
		}
	}()

	for {
		char, _, err = reader.ReadRune()
		if err != nil {
			continue
		}
		switch char {
		case 'q', CTRL_C, CTRL_D:
			// calling cancel() closes ctx.Done in main which unblocks and allows
			// shutdown
			cancel()
			return
		case 's', SPC:
			input <- char
		default:
			continue
		}
	}
}

func Stop() {
	resetCursor()
	clearScreen()
	_ = term.Restore(int(os.Stdin.Fd()), prevState)
}
