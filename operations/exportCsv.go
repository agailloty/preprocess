package operations

import (
	"github.com/agailloty/preprocess/dataset"
)

func ExportCsv(df dataset.DataFrame, filename string) {
	fileName := df.Name + "_EXPORT.csv"
	if filename != "" {
		fileName = filename
	}

	df.SaveToCSV(fileName, ",")
}
