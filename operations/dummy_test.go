package operations

import (
	"strings"
	"testing"

	"github.com/agailloty/preprocess/dataset"
	"github.com/stretchr/testify/assert"
)

func TestMakeDummy(t *testing.T) {
	countries := []dataset.Nullable[string]{
		{Value: "Japan", IsValid: true},
		{Value: "France", IsValid: true},
		{Value: "France", IsValid: true},
		{Value: "Italy", IsValid: true},
		{Value: "Germany", IsValid: true},
		{Value: "Germany", IsValid: true},
		{Value: "France", IsValid: true},
	}

	countryCol := dataset.String{Name: "Countries", Data: countries}

	df := dataset.DataFrame{Columns: []dataset.DataSetColumn{&countryCol}}

	dummyOp := dummyOperation{df: &df, col: &countryCol, dropLast: false, prefixColName: false}

	t.Run("Dummy without drop last returns 4 elements", func(t *testing.T) {
		res, err := makeDummy(dummyOp)
		assert.NoError(t, err)
		assert.Equal(t, len(res), 4)
	})

	dummyOp = dummyOperation{df: &df, col: &countryCol, dropLast: true, prefixColName: false}

	t.Run("Dummy with drop last returns  elements", func(t *testing.T) {
		res, err := makeDummy(dummyOp)
		assert.NoError(t, err)
		assert.Equal(t, len(res), 3)
	})

	dummyOp = dummyOperation{df: &df, col: &countryCol, dropLast: true, prefixColName: false}

	t.Run("Dummy without prefix does not affect new column names", func(t *testing.T) {
		prefix := "Countries"
		var results []bool
		res, err := makeDummy(dummyOp)
		assert.NoError(t, err)
		for _, col := range res {
			isPrefixed := !strings.HasPrefix(col.Name, prefix)
			results = append(results, isPrefixed)
		}
		for _, v := range results {
			assert.True(t, v)
		}
	})

	dummyOp = dummyOperation{df: &df, col: &countryCol, dropLast: true, prefixColName: true}

	t.Run("Dummy with prefix affect new column names", func(t *testing.T) {
		prefix := "Countries"
		var results []bool
		res, err := makeDummy(dummyOp)
		assert.NoError(t, err)
		for _, col := range res {
			isPrefixed := strings.HasPrefix(col.Name, prefix)
			results = append(results, isPrefixed)
		}
		for _, v := range results {
			assert.True(t, v)
		}
	})

	dummyOp = dummyOperation{df: &df, col: &countryCol, dropLast: false, prefixColName: true}

	t.Run("Number of ones equals number of rows", func(t *testing.T) {
		res, err := makeDummy(dummyOp)
		assert.NoError(t, err)
		onesCount := 0

		for _, col := range res {
			for _, data := range col.Data {
				if data.Value == 1 {
					onesCount++
				}
			}
		}

		assert.Equal(t, onesCount, countryCol.Length())
	})

}
