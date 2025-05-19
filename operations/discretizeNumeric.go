package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func makeBinsFromNumericColumns(column dataset.DataSetColumn, bins []config.BinningOperation, df *dataset.DataFrame, overrideColumn bool) {
	binData := make([]dataset.Nullable[string], column.Length())

	switch v := column.(type) {
	case *dataset.Float:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
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
	case *dataset.Integer:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range bins {
				if float32(val.Value) >= bin.Lower && float32(val.Value) <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: column.ValueAt(i)}
				}
			}
		}
	}

	columnName := column.GetName()
	if !overrideColumn {
		columnName = columnName + "_C"
	}
	binnedColumn := dataset.String{
		Name: columnName,
		Data: binData,
	}

	if overrideColumn {
		for i := range df.Columns {
			if df.Columns[i].GetName() == columnName {
				df.Columns[i] = &binnedColumn
			}
		}
	} else {
		df.Columns = append(df.Columns, &binnedColumn)
	}
}
