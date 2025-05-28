package summary

import (
	"maps"
	"slices"
	"sort"

	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/statistics"
	"github.com/agailloty/preprocess/utils"
)

func Summarize(df dataset.DataFrame, excluded []string, filename string) {
	summaryFile := GetSummaryFile(df, excluded)
	utils.SerializeStruct(summaryFile, filename)
}

func GetSummaryFile(df dataset.DataFrame, excluded []string) SummaryFile {
	numericColumnCount := 0
	stringColumnCount := 0
	for _, col := range excluded {
		df.DeleteColumnByName(col)
	}
	colSummaries := make([]ColumnSummary, df.ColumnsCount)
	for i, col := range df.Columns {
		if utils.Contains(excluded, col.GetName()) {
			continue
		}
		switch v := col.(type) {
		case *dataset.String:
			colSummaries[i] = summarizeStringColumn(v)
			stringColumnCount++
		case *dataset.Float:
			colSummaries[i] = summarizeFloatColumn(v)
			numericColumnCount++
		case *dataset.Integer:
			colSummaries[i] = summarizeIntColumn(v)
			numericColumnCount++
		}
	}

	summaryFile := SummaryFile{Data: df.DataSpecs,
		DataSummary: DatasetSummary{
			RowCount:       df.RowsCount,
			NumericColumns: numericColumnCount,
			StringColumns:  stringColumnCount,
			ColumnCount:    numericColumnCount + stringColumnCount,
		},
		Columns: colSummaries}

	return summaryFile
}

func summarizeStringColumn(column *dataset.String) ColumnSummary {
	summary := make(map[string]ValueKeyCount)
	for _, value := range column.Data {
		var modality ValueKeyCount
		var ok bool
		modality, ok = summary[value.Value]
		if !ok {
			summary[value.Value] = ValueKeyCount{Key: value.Value, Count: 1}
		} else {
			modality.Count++
			summary[value.Value] = modality
		}
	}

	uniqueValueCount := len(summary)

	uniqueValuesSummary := slices.Collect(maps.Values(summary))
	keys := slices.Collect(maps.Keys(summary))
	missingCount := column.CountMissing()

	sort.Slice(uniqueValuesSummary, func(i, j int) bool {
		return uniqueValuesSummary[i].Count > uniqueValuesSummary[j].Count
	})

	return ColumnSummary{
		Name:                column.Name,
		Type:                column.GetType(),
		RowCount:            column.Length(),
		UniqueValueCount:    uniqueValueCount,
		UniqueValues:        keys,
		UniqueValuesSummary: uniqueValuesSummary,
		Missing:             missingCount}
}

func summarizeFloatColumn(column *dataset.Float) ColumnSummary {
	validData := utils.ExtractNonNullFloats(column.Data)
	mean := statistics.Mean(validData)
	median := statistics.Median(validData)
	missingCount := column.CountMissing()

	min, max := statistics.MinMax(validData)

	return ColumnSummary{
		Name:     column.Name,
		Type:     "numeric",
		RowCount: column.Length(),
		Min:      min,
		Max:      max,
		Mean:     mean,
		Median:   median,
		Missing:  missingCount}
}

func summarizeIntColumn(column *dataset.Integer) ColumnSummary {
	validData := utils.ExtractNonNullInts(column.Data)
	mean := statistics.Mean(validData)
	median := statistics.Median(validData)
	missingCount := column.CountMissing()

	min, max := statistics.MinMax(validData)

	return ColumnSummary{
		Name:     column.Name,
		Type:     "numeric",
		RowCount: column.Length(),
		Min:      float64(min),
		Max:      float64(max),
		Mean:     mean,
		Median:   median,
		Missing:  missingCount}
}
