package operations

import (
	"errors"
	"log"

	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func makeDummy(specs dummyOperation) (error, []dataset.Integer) {
	uniqueValues := utils.ExtractUniqueValues(specs.col.Data)

	if specs.isExcluded {
		return errors.New("EXCLUDED_COLUMN"), []dataset.Integer{}
	}

	if specs.dropLast {
		uniqueValues = uniqueValues[:len(uniqueValues)-1]
	}

	if len(uniqueValues) >= 500 && !specs.continueWithTooMany {
		log.Fatalf(`[Dummy operation] : There are too many values for %s. Total count : %d. Use exclude_columns = ["%s"] to exclude it.`, specs.col.Name, len(uniqueValues), specs.col.Name)
		return errors.New("DUMMY_TOO_MANY_VALUES"), []dataset.Integer{}
	}

	dummyCols := make([]dataset.Integer, len(uniqueValues))

	for i, uniqueVal := range uniqueValues {
		dummyCol := makeIntegerColumn(specs.col.Name, uniqueVal.Key, specs.col.Length(), specs.prefixColName)
		for idx, value := range specs.col.Data {
			dummyValue := 0
			if value.Value == uniqueVal.Key {
				dummyValue = 1
			}
			dummyCol.Data[idx].Value = dummyValue
			dummyCol.Data[idx].IsValid = true
		}
		dummyCols[i] = dummyCol
	}

	return nil, dummyCols
}

func makeIntegerColumn(colName string, modalityName string, length int, prefixName bool) dataset.Integer {
	dummyName := modalityName
	if prefixName {
		dummyName = colName + "_" + modalityName
	}
	column := dataset.Integer{Name: dummyName, Data: make([]dataset.Nullable[int], length)}
	return column
}
