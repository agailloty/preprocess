package operations

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func DispatchOperations(prepfile *config.Config) {
	df := dataset.ReadDataFrame(prepfile.Data.File, prepfile.Data.Separator)
	fmt.Printf("Successfully read dataset %s \n", prepfile.Data.File)

	for _, col := range df.Columns {
		RenameColumn(col, prepfile.Data.Columns)
	}

	ExportCsv(df, prepfile)
}
