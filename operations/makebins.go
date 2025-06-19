package operations

import (
	"errors"

	"github.com/agailloty/preprocess/dataset"
)

func makeBinsFromNumericColumns(specs binOperation) error {
	binData := make([]dataset.Nullable[string], (*specs.col).Length())
	var problem error
	switch v := (*specs.col).(type) {
	case *dataset.Float:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range specs.bins {
				if val.Value >= bin.Lower && val.Value <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					if problem == nil {
						problem = errors.New("INTERVAL_MISSING")
					}
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: v.ValueAt(i)}
				}
			}
		}
	case *dataset.Integer:
		for i, val := range v.Data {
			binFound := false
			for _, bin := range specs.bins {
				if float64(val.Value) >= bin.Lower && float64(val.Value) <= bin.Upper {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: bin.Label}
					binFound = true
					break
				}
				if !binFound {
					binData[i] = dataset.Nullable[string]{IsValid: true, Value: v.ValueAt(i)}
				}
			}
		}
	}

	columnName := (*specs.col).GetName()
	if !specs.overrideColumn {
		columnName = columnName + "_C"
	}
	binnedColumn := dataset.String{
		Name: columnName,
		Data: binData,
	}

	if specs.overrideColumn {
		for i := range specs.df.Columns {
			if specs.df.Columns[i].GetName() == columnName {
				specs.df.Columns[i] = &binnedColumn
			}
		}
	} else {
		specs.df.Columns = append(specs.df.Columns, &binnedColumn)
	}

	return problem
}
