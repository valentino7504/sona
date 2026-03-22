package ui

import "fmt"

const CSI = "\033["

func clr() string {
	return (CSI + "2J\033[H")
}

func showCursor(show bool) string {
	if show {
		return (CSI + "?25h")
	}
	return (CSI + "?25l")
}

func moveCursorUp(lines int) string {
	return fmt.Sprintf(CSI+"%dA", lines)
}

func moveCursorDown(lines int) string {
	return fmt.Sprintf(CSI+"%dB", lines)
}

func moveCursorRight(columns int) string {
	return fmt.Sprintf(CSI+"%dC", columns)
}

func moveCursorLeft(columns int) string {
	return fmt.Sprintf(CSI+"%dD", columns)
}

func cursorHome() string {
	return (CSI + "H")
}

func moveCursorTo(x int, y int) string {
	return fmt.Sprintf(CSI+"%d;%dH", y, x)
}
