package utils

import (
	"path/filepath"
	"strings"

	"github.com/agailloty/preprocess/dataset"
)

func AppendPrefixOrSuffix(filename string, prefix string, suffix string) string {
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filepath.Base(filename), ext)
	newName := prefix + base + suffix + ext

	return newName
}

func OverrideDataFrameColumn(df *dataset.DataFrame, columnName string, newColumn dataset.DataSetColumn) {
	for i := range df.Columns {
		if df.Columns[i].GetName() == columnName {
			df.Columns[i] = newColumn
		}
	}
}
