package audio

import (
	"path/filepath"
	"strings"
)

func extractFileFormat(fileName string) string {
	ext := filepath.Ext(fileName)
	if ext == "" {
		return ""
	}
	if len(ext) > 1 {
		return strings.ToLower(ext[1:])
	}
	return ""
}
