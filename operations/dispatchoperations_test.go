package operations

import (
	"testing"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/stretchr/testify/assert"
)

const (
	fifaCsvFile     = "../testdata/fifa_players.csv"
	numericPrepFile = "../testdata/fifa_text_ops.toml"
)

func TestNumericOperations(t *testing.T) {
	dfSpec := common.DataSpecs{
		Filename:          fifaCsvFile,
		CsvSeparator:      ",",
		DecimalSeparator:  ".",
		Encoding:          "utf-8",
		MissingIdentifier: "",
	}
	df := dataset.ReadDataFrame(dfSpec)
	prepfile, _ := config.LoadConfigFromPrepfile(numericPrepFile)

	t.Run("Numeric operations are correctly counted", func(t *testing.T) {
		summary := summarizeOperations(prepfile, &df)
		assert.Equal(t, summary.numericOperations, 88)
	})
}
