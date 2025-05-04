package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func ExportCsv(df dataset.DataFrame, exportConf config.ExportConfig) {
	fileName := df.Name + "_EXPORT.csv"
	if exportConf.Path != "" {
		fileName = exportConf.Path
	}

	df.SaveToCSV(fileName, ",")
}
