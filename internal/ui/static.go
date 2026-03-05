package ui

import (
	"os"

	"golang.org/x/term"
)

var (
	barStartRow  int
	trackInfoRow int
)

func drawStatic() {
	barStartRow = 8
	_, termHeight, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	// height of terminal - height of header rows - number of rows at the bottom
	barHeight := termHeight - 7 - 8
	// get location of the row where the play icon + duration bar is
	trackInfoRow = barStartRow + barHeight + 1
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
