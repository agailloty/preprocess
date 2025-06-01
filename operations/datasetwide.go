package operations

import (
	"cmp"
	"log"
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

func ExportCsv(df dataset.DataFrame, filename string) {
	fileName := df.Name + "_EXPORT.csv"
	if filename != "" {
		fileName = filename
	}

	df.SaveToCSV(fileName, ",")
	log.Printf("Successfully exported : %s", fileName)
}
