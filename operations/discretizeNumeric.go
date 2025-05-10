package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func makeBinsFromNumericColumns(column dataset.DataSetColumn, binSpec config.BinningSpecs, df *dataset.DataFrame) {
	binData := make([]dataset.Nullable[string], column.Length())

	switch v := column.(type) {
	case *dataset.Float:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range binSpec.Bins {
				if val.Value >= bin.Lower && val.Value <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
		columnName := binSpec.NewColumn
		overrideExistingColumn := false
		if binSpec.NewColumn == "" {
			columnName = column.GetName()
			overrideExistingColumn = true
		}
		binnedColumn := dataset.String{
			Name: columnName,
			Data: binData,
		}

		if overrideExistingColumn {
			for i := range df.Columns {
				if df.Columns[i].GetName() == columnName {
					df.Columns[i] = &binnedColumn
				}
			}
		} else {
			df.Columns = append(df.Columns, &binnedColumn)
		}
	}
}
