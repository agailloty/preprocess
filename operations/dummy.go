package operations

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func addDummyToDataframe(df *dataset.DataFrame, col dataset.DataSetColumn, dropLast bool, dropInitialCol bool, prefixColName bool) {
	switch v := col.(type) {
	case *dataset.String:
		dummyCols := makeDummy(v, dropLast, prefixColName)
		for _, dummy := range dummyCols {
			df.Columns = append(df.Columns, &dummy)
		}

		if dropInitialCol {
			df.DeleteColumn(col)
		}
	}
}

func makeDummy(column *dataset.String, dropLast bool, prefixColName bool) []dataset.Integer {
	uniqueValues := utils.ExtractUniqueValues(column.Data)

	if dropLast {
		uniqueValues = uniqueValues[:len(uniqueValues)-1]
	}

	dummyCols := make([]dataset.Integer, len(uniqueValues))

	for i, uniqueVal := range uniqueValues {
		dummyCol := makeIntegerColumn(column.Name, uniqueVal.Key, column.Length(), prefixColName)
		for idx, value := range column.Data {
			dummyValue := 0
			if value.Value == uniqueVal.Key {
				dummyValue = 1
			}
			dummyCol.Data[idx].Value = dummyValue
		}
		dummyCols[i] = dummyCol
	}

	return dummyCols
}

func makeIntegerColumn(colName string, modalityName string, length int, prefixName bool) dataset.Integer {
	dummyName := modalityName
	if prefixName {
		dummyName = colName + "_" + modalityName
	}
	column := dataset.Integer{Name: dummyName, Data: make([]dataset.Nullable[int], length)}
	return column
}
