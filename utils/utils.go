package utils

import (
	"path/filepath"
	"strings"
)

func AppendPrefixOrSuffix(filename string, prefix string, suffix string) string {
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filepath.Base(filename), ext)
	newName := prefix + base + suffix + ext

	return newName
}
