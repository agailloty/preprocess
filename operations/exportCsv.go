package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func ExportCsv(df dataset.DataFrame, prepfile *config.Config) {
	fileName := df.Name + "_EXPORT.csv"
	df.SaveToCSV(fileName, ",")
}
