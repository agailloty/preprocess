package operations

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func DispatchOperations(prepfile *config.Config) {
	df := dataset.ReadDataFrame(prepfile.Data.File, prepfile.Data.Separator)
	fmt.Printf("Successfully read dataset %s \n", prepfile.Data.File)

	ExportCsv(df, prepfile)
}
