package summary

import (
	"maps"
	"slices"

	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
	"github.com/agailloty/preprocess/utils"
)

func Summarize(df dataset.DataFrame, filename string) {
	colSummaries := make([]ColumnSummary, df.ColumnsCount)
	for i, col := range df.Columns {
		switch v := col.(type) {
		case *dataset.String:
			colSummaries[i] = summarizeStringColumn(v)
		case *dataset.Float:
			colSummaries[i] = summarizeFloatColumn(v)
		case *dataset.Integer:
			colSummaries[i] = summarizeIntColumn(v)
		}
	}

	summaryFile := SummaryFile{Columns: colSummaries}

	utils.SerializeStruct(summaryFile, filename)
}

func summarizeStringColumn(column *dataset.String) ColumnSummary {
	summary := make(map[string]ValueKeyCount)
	for _, value := range column.Data {
		var modality ValueKeyCount
		var ok bool
		if !value.IsValid {
			modality = summary["NA"]
			modality.Count++
			summary[value.Value] = modality
		} else {
			modality, ok = summary[value.Value]
			if !ok {
				summary[value.Value] = ValueKeyCount{Key: value.Value, Count: 1}
			} else {
				modality.Count++
				summary[value.Value] = modality
			}
		}
	}

	uniqueValueCount := len(summary)
	keys := make([]string, 0, len(summary))

	for k := range summary {
		keys = append(keys, k)
	}

	uniqueValuesSummary := slices.Collect(maps.Values(summary))
	missingCount := utils.CountMissing(&column.Data)

	return ColumnSummary{
		Name:                column.Name,
		RowCount:            column.Length(),
		UniqueValueCount:    uniqueValueCount,
		UniqueValues:        keys,
		UniqueValuesSummary: uniqueValuesSummary,
		MissingCount:        missingCount}
}

func summarizeFloatColumn(column *dataset.Float) ColumnSummary {
	validData := utils.ExtractNonNullFloats(column.Data)
	mean := statistics.Mean(validData)
	median := statistics.Median(validData)
	missingCount := utils.CountMissing(&column.Data)

	return ColumnSummary{
		Name:         column.Name,
		RowCount:     column.Length(),
		Mean:         mean,
		Median:       median,
		MissingCount: missingCount}
}

func summarizeIntColumn(column *dataset.Integer) ColumnSummary {
	validData := utils.ExtractNonNullInts(column.Data)
	mean := statistics.Mean(validData)
	median := statistics.Median(validData)
	missingCount := utils.CountMissing(&column.Data)

	return ColumnSummary{
		Name:         column.Name,
		RowCount:     column.Length(),
		Mean:         mean,
		Median:       median,
		MissingCount: missingCount}
}
