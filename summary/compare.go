package summary

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

type DatasetDiff struct {
	AddedColumns     []dataset.DataSetColumn
	RemovedColumns   []dataset.DataSetColumn
	AlteredColumns   []AlteredColumn
	IdenticalColumns []dataset.DataSetColumn
}

type AlteredColumn struct {
	SourceColumn       dataset.DataSetColumn
	StringColumnDiff   StringColumnDiff
	NumericColumnDiffs NumericColumnDiff
}

type StringColumnDiff struct {
	RemovedValues []string
	AddedValues   []string
}

type NumericColumnDiff struct {
}

func ComputeDiffs(baseDf *dataset.DataFrame, newDf *dataset.DataFrame) DatasetDiff {
	var removedColumns []dataset.DataSetColumn
	for _, column := range baseDf.Columns {
		if !newDf.ColumnExists(column) {
			removedColumns = append(removedColumns, column)
		}
	}

	var addedColumns []dataset.DataSetColumn

	for _, column := range newDf.Columns {
		if !baseDf.ColumnExists(column) {
			addedColumns = append(addedColumns, column)
		}
	}

	var alteredColumns []AlteredColumn
	var identicalColumns []dataset.DataSetColumn
	for _, column := range newDf.Columns {
		baseCol := baseDf.GetColumn(column)
		if baseCol == nil {
			continue
		}
		if baseCol.GetType() == "string" && isColumnDifferent(baseCol, column) {
			newDfColumn, ok := column.(*dataset.String)
			if ok {
				baseDfColumn, _ := baseCol.(*dataset.String)
				alteredColumn := AlteredColumn{SourceColumn: baseDfColumn, StringColumnDiff: getStringColumnDiff(baseDfColumn, newDfColumn)}
				alteredColumns = append(alteredColumns, alteredColumn)
			}
		} else if baseCol.GetType() == "string" && !isColumnDifferent(baseCol, column) {
			identicalColumns = append(identicalColumns, baseCol)
		} else {
			if isColumnDifferent(baseCol, column) {
				alteredColumn := AlteredColumn{SourceColumn: baseCol}
				alteredColumns = append(alteredColumns, alteredColumn)
			} else if !isColumnDifferent(baseCol, column) {
				identicalColumns = append(identicalColumns, baseCol)
			}
		}
	}

	return DatasetDiff{
		AddedColumns:     addedColumns,
		RemovedColumns:   removedColumns,
		AlteredColumns:   alteredColumns,
		IdenticalColumns: identicalColumns,
	}
}

func getStringColumnDiff(base *dataset.String, new *dataset.String) StringColumnDiff {
	added, removed := utils.GetDiff(base.Data, new.Data)
	return StringColumnDiff{
		RemovedValues: removed,
		AddedValues:   added,
	}
}

func isColumnDifferent(source dataset.DataSetColumn, target dataset.DataSetColumn) bool {
	if source.Length() != target.Length() {
		return true
	}

	for i := 0; i < source.Length()-1; i++ {
		if source.ValueAt(i) != target.ValueAt(i) {
			return true
		}
	}

	return false
}
