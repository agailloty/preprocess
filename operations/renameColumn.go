package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func RenameColumn(column dataset.DataSetColumn, conf []config.ColumnConfig) {

	for _, colConfig := range conf {
		if colConfig.Name == column.GetName() {
			if colConfig.NewName != "" {
				column.SetName(colConfig.NewName)
			}
			break
		}
	}
}
