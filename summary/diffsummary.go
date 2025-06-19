package summary

import (
	"github.com/agailloty/preprocess/dataset"
)

type DiffSummary struct {
	SourceDataSummary DatasetSummary  `toml:"source_data_summary"`
	TargetDataSummary DatasetSummary  `toml:"target_data_summary"`
	Columns           []ColumnSummary `toml:"columns"`
}

// This function aims to produce a diff summary of two datasets
// it returns a SummaryFile
func GenerateDiffSummary(source *dataset.DataFrame, target *dataset.DataFrame) DiffSummary {
	dfDiff := ComputeDiffs(source, target)
	targetSummary := GetSummaryFile(*target, []string{})
	sourceSummary := GetSummaryFile(*source, []string{})
	var colSummaries []ColumnSummary

	for _, diffColumn := range dfDiff.AddedColumns {
		for _, colSummary := range targetSummary.Columns {
			columnType := diffColumn.GetType()
			if columnType == "int" || columnType == "float" {
				columnType = "numeric"
			}
			if diffColumn.GetName() == colSummary.Name && columnType == colSummary.Type {
				colSummary.IsAdded = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.RemovedColumns {
		for _, colSummary := range sourceSummary.Columns {
			columnType := diffColumn.GetType()
			if columnType == "int" || columnType == "float" {
				columnType = "numeric"
			}
			if diffColumn.GetName() == colSummary.Name && columnType == colSummary.Type {
				colSummary.IsDeleted = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.AlteredColumns {
		for _, colSummary := range targetSummary.Columns {
			columnType := diffColumn.SourceColumn.GetType()
			if columnType == "int" || columnType == "float" {
				columnType = "numeric"
			}
			if diffColumn.SourceColumn.GetName() == colSummary.Name && columnType == colSummary.Type {
				colSummary.IsAltered = true
				colSummary.AddedStringValues = diffColumn.StringColumnDiff.AddedValues
				colSummary.RemovedStringValues = diffColumn.StringColumnDiff.RemovedValues
				colSummary.OldStats = sourceSummary.extractNumericColumStats(colSummary.Name, columnType)
				colSummary.NewStats = targetSummary.extractNumericColumStats(colSummary.Name, columnType)
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.IdenticalColumns {
		for _, colSummary := range targetSummary.Columns {
			columnType := diffColumn.GetType()
			if columnType == "int" || columnType == "float" {
				columnType = "numeric"
			}
			if diffColumn.GetName() == colSummary.Name && columnType == colSummary.Type {
				colSummary.IsIdentical = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	return DiffSummary{SourceDataSummary: sourceSummary.DataSummary,
		TargetDataSummary: targetSummary.DataSummary,
		Columns:           colSummaries}

}

func (s *SummaryFile) extractNumericColumStats(colName string, colType string) *NumericStats {
	var numericSummary NumericStats
	for _, col := range s.Columns {
		if col.Name == colName && col.Type == colType {
			numericSummary = NumericStats{
				RowCount: col.RowCount,
				Missing:  col.Missing,
				Mean:     col.Mean,
				Median:   col.Median,
				Min:      col.Min,
				Max:      col.Max,
			}
			break
		}
	}
	return &numericSummary
}
