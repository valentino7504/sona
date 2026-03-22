package ui

import "strings"

func drawBars(binVals []float64) {
	draw(
		moveCursorTo(0, barBottomRow),
		bars(binVals),
	)
}

func bars(bins []float64) string {
	var builder strings.Builder
	barHeight := barBottomRow - barTopRow
	maxBinVal := getMaxVal(bins)
	heights := make([]int, len(bins))
	for i, binVal := range bins {
		heights[i] = int((binVal / maxBinVal) * float64(barHeight))
	}
	fillChar, emptyChar := '█', ' '
	for i := range barHeight {
		currLineNo := barBottomRow - i
		builder.WriteString(moveCursorTo(0, currLineNo))
		for _, h := range heights {
			if i < h {
				builder.WriteRune(fillChar)
			} else {
				builder.WriteRune(emptyChar)
			}
			builder.WriteString(moveCursorRight(2))
		}
	}
	return builder.String()
}

func getMaxVal[T int | float64](vals []T) T {
	maxVal := vals[0]
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
