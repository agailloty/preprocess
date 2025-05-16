package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func applyOperationsOnNumericColumns(df *dataset.DataFrame, operations *[]config.PreprocessOp) {
	for _, column := range df.Columns {
		if column.GetType() == "int" || column.GetType() == "float" {
			applyNumericOperationsOnColumn(operations, column, df)
		}
	}
}
