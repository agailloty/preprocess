package summary

import (
	"maps"
	"slices"

	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func Summarize(df dataset.DataFrame, filename string) {
	colSummaries := make([]ColumnSummary, df.ColumnsCount)
	for i, col := range df.Columns {
		switch v := col.(type) {
		case *dataset.String:
			colSummaries[i] = summarizeStringColumn(v)
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

	return ColumnSummary{
		Name:                column.Name,
		RowCount:            column.Length(),
		UniqueValueCount:    uniqueValueCount,
		UniqueValues:        keys,
		UniqueValuesSummary: uniqueValuesSummary}
}
