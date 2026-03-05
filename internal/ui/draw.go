package ui

import (
	"os"
	"strings"
)

func draw(items ...string) {
	var builder strings.Builder
	for _, str := range items {
		builder.WriteString(str)
	}
	_, _ = os.Stdout.Write([]byte(builder.String()))
}
