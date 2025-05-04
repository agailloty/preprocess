package operations

import (
	"cmp"
	"slices"

	"github.com/agailloty/preprocess/dataset"
)

func SortDatasetColumns(df dataset.DataFrame, descending bool) {
	if descending {
		slices.SortFunc(df.Columns, compareColumnNamesDesc)
	} else {
		slices.SortFunc(df.Columns, compareColumnNamesAsc)
	}

}

func compareColumnNamesAsc(colA dataset.DataSetColumn, colB dataset.DataSetColumn) int {
	return cmp.Compare(colA.GetName(), colB.GetName())
}

func compareColumnNamesDesc(colA dataset.DataSetColumn, colB dataset.DataSetColumn) int {
	return -1 * cmp.Compare(colA.GetName(), colB.GetName())
}
