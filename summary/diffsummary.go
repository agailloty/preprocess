package summary

import (
	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/dataset"
)

type DiffSummary struct {
	Data        common.DataSpecs `toml:"data"`
	DataSummary DatasetSummary   `toml:"data_summary"`
	Columns     []ColumnSummary  `toml:"columns"`
}

// This function aims to produce a diff summary of two datasets
// it returns a SummaryFile
func GenerateDiffSummary(source *dataset.DataFrame, target *dataset.DataFrame) DiffSummary {
	dfDiff := ComputeDiffs(source, target)
	targetSummary := GetSummaryFile(*target, []string{})
	var colSummaries []ColumnSummary

	for _, diffColumn := range dfDiff.AddedColumns {
		for _, colSummary := range targetSummary.Columns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsAdded = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.RemovedColumns {
		for _, colSummary := range targetSummary.Columns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsDeleted = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.AlteredColumns {
		for _, colSummary := range targetSummary.Columns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsAltered = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	for _, diffColumn := range dfDiff.IdenticalColumns {
		for _, colSummary := range targetSummary.Columns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsIdentical = true
				colSummaries = append(colSummaries, colSummary)
			}
		}
	}

	return DiffSummary{Data: target.DataSpecs, DataSummary: targetSummary.DataSummary, Columns: colSummaries}

}
