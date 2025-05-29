package statistics

import (
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

func ScaleWithZscore(column dataset.DataSetColumn, df *dataset.DataFrame) {
	scaleNumericColumn(column, df, ComputeZScore)
}

func ScaleWithMinMax(column dataset.DataSetColumn, df *dataset.DataFrame) {
	scaleNumericColumn(column, df, ComputeMinMaxScore)
}

func scaleNumericColumn(column dataset.DataSetColumn, df *dataset.DataFrame, scaleFunc scaleFunc) {
	switch v := column.(type) {
	case *dataset.Integer:
		// Convert into to float64
		validData := utils.ExtractNonNullInts(v.Data)
		mu := Mean(validData)
		sigma := StdDev(validData)
		zScores := make([]dataset.Nullable[float64], column.Length())
		for i := range v.Data {
			zScore := scaleFunc(float64(v.Data[i].Value), float64(mu), float64(sigma))
			zScores[i] = dataset.Nullable[float64]{IsValid: v.Data[i].IsValid, Value: zScore}
		}
		newColumn := dataset.Float{Name: column.GetName(), Data: zScores}
		utils.OverrideDataFrameColumn(df, column.GetName(), &newColumn)

	case *dataset.Float:
		validData := utils.ExtractNonNullFloats(v.Data)
		mu := Mean(validData)
		sigma := StdDev(validData)
		for i := range v.Data {
			zScore := ComputeZScore(float64(v.Data[i].Value), float64(mu), float64(sigma))
			v.Data[i] = dataset.Nullable[float64]{IsValid: v.Data[i].IsValid, Value: zScore}
		}
	}
}
