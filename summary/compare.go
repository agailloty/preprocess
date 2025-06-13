package summary

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

type DatasetDiff struct {
	AddedColumns       []dataset.DataSetColumn
	RemovedColumns     []dataset.DataSetColumn
	AlteredColumns     []dataset.DataSetColumn
	IdenticalColumns   []dataset.DataSetColumn
	StringColumnDiffs  []StringColumnDiff
	NumericColumnDiffs []NumericColumnDiff
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

	var alteredColumns []dataset.DataSetColumn
	var stringColumnDiffs []StringColumnDiff
	var identicalColumns []dataset.DataSetColumn
	for _, column := range newDf.Columns {
		baseCol := baseDf.GetColumn(column)
		if baseCol != nil && isColumnDifferent(baseCol, column) {
			alteredColumns = append(alteredColumns, column)
			newDfColumn, ok := column.(*dataset.String)
			if ok {
				baseDfColumn, _ := column.(*dataset.String)
				stringColumnDiffs = append(stringColumnDiffs, getStringColumnDiff(baseDfColumn, newDfColumn))
			}
		} else if baseCol != nil && !isColumnDifferent(baseCol, column) {
			identicalColumns = append(identicalColumns, baseCol)
		}
	}

	return DatasetDiff{
		AddedColumns:      addedColumns,
		RemovedColumns:    removedColumns,
		AlteredColumns:    alteredColumns,
		StringColumnDiffs: stringColumnDiffs,
		IdenticalColumns:  identicalColumns,
	}
}

func getStringColumnDiff(base *dataset.String, new *dataset.String) StringColumnDiff {
	added, removed := utils.GetDiff(base.Data, new.Data)
	addedValues := make([]string, len(added))
	removedValues := make([]string, len(removed))

	for i, v := range added {
		addedValues[i] = v.Value
	}

	for i, v := range removed {
		removedValues[i] = v.Value
	}

	return StringColumnDiff{
		RemovedValues: removedValues,
		AddedValues:   addedValues,
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
