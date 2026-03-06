package ui

import (
	"os"

	"golang.org/x/term"
)

var (
	barTopRow    int
	barBottomRow int
	trackInfoRow int
)

func drawStatic() {
	barTopRow = 8
	_, termHeight, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	// height of terminal - height of header rows - number of rows at the bottom
	barHeight := termHeight - 7 - 8
	// get location of the row where the play icon + duration bar is and where the frequency bars
	// start
	trackInfoRow = barTopRow + barHeight + 1
	barBottomRow = trackInfoRow - 2
	draw(
		showCursor(false),
		clearScreen(),
		header(),
		horizontalBar(),
		moveCursorDown(barHeight),
		horizontalBar(),
		moveCursorDown(3),
		footer(),
		horizontalBar(),
	)
}

func resetCursor() {
	draw(showCursor(true))
}
