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

	targetColumnsSummary := targetSummary.Columns

	for _, colSummary := range targetColumnsSummary {
		for _, diffColumn := range dfDiff.AddedColumns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsAdded = true
			}
		}

		for _, diffColumn := range dfDiff.RemovedColumns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsDeleted = true
			}
		}

		for _, diffColumn := range dfDiff.AlteredColumns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsAltered = true
			}
		}

		for _, diffColumn := range dfDiff.IdenticalColumns {
			if diffColumn.GetName() == colSummary.Name && diffColumn.GetType() == colSummary.Type {
				colSummary.IsIdentical = true
			}
		}
	}

	return DiffSummary{Data: target.DataSpecs, DataSummary: targetSummary.DataSummary, Columns: targetColumnsSummary}

}
