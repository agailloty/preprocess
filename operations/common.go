package operations

import (
	"github.com/agailloty/preprocess/dataset"
)

func replaceMissingValues(column dataset.DataSetColumn, value any) {
	switch v := column.(type) {
	case *dataset.String:
		fillMissingStringWithValue(v, value.(string))
	case *dataset.Integer:
		val, ok := value.(int64)
		if ok {
			fillMissingIntegerWithValue(v, int(val))
		}
	case *dataset.Float:
		val, ok := value.(float64)
		if ok {
			fillMissingFloatWithValue(v, float64(val))
		}
	}
}

func fillMissingStringWithValue(column *dataset.String, newValue string) {
	for i := range column.Data {
		if !column.Data[i].IsValid {
			column.Data[i].Value = newValue
			column.Data[i].IsValid = true
		}
	}
}

func fillMissingIntegerWithValue(column *dataset.Integer, newValue int) {
	for i := range column.Data {
		if !column.Data[i].IsValid {
			column.Data[i].Value = newValue
			column.Data[i].IsValid = true
		}
	}
}

func fillMissingFloatWithValue(column *dataset.Float, newValue float64) {
	for i := range column.Data {
		if !column.Data[i].IsValid {
			column.Data[i].Value = newValue
			column.Data[i].IsValid = true
		}
	}
}
