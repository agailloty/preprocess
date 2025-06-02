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

	t.Run("Dummy without drop last returns 4 elements", func(t *testing.T) {
		res := makeDummy(&countryCol, false, false, true)
		assert.Equal(t, len(res), 4)
	})

	t.Run("Dummy with drop last returns  elements", func(t *testing.T) {
		res := makeDummy(&countryCol, true, false, true)
		assert.Equal(t, len(res), 3)
	})

	t.Run("Dummy without prefix does not affect new column names", func(t *testing.T) {
		prefix := "Countries"
		var results []bool
		res := makeDummy(&countryCol, true, false, true)

		for _, col := range res {
			isPrefixed := !strings.HasPrefix(prefix, col.Name)
			results = append(results, isPrefixed)
		}
		for _, v := range results {
			assert.True(t, v)
		}
	})
}
