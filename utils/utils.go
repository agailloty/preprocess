package utils

import (
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/dataset"
	"github.com/pelletier/go-toml/v2"
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

func SerializeStruct(content interface{}, filename string) {
	configFile := content
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error generating %s : %v", filename, err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(configFile); err != nil {
		log.Fatalf("An error occured during TOML enconding : %v", err)
	}
	log.Printf("%s successfully generated.\n", filename)
}

func ExtractNonNullInts(data []dataset.Nullable[int]) []int {
	result := make([]int, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		}
	}
	return result
}

func ExtractNonNullFloats(data []dataset.Nullable[float64]) []float64 {
	result := make([]float64, len(data))
	for i, item := range data {
		if item.IsValid {
			result[i] = item.Value
		}
	}
	return result
}

func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ExtractUniqueValues(data []dataset.Nullable[string]) []common.ValueKeyCount {
	summary := make(map[string]common.ValueKeyCount)
	for _, value := range data {
		var modality common.ValueKeyCount
		var ok bool
		modality, ok = summary[value.Value]
		if !ok {
			summary[value.Value] = common.ValueKeyCount{Key: value.Value, Count: 1}
		} else {
			modality.Count++
			summary[value.Value] = modality
		}
	}
	unique := slices.Collect(maps.Values(summary))

	return unique
}
