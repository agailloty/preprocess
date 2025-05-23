package config

import "github.com/agailloty/preprocess/common"

var defaultDataSpec = common.DataSpecs{Filename: "data.csv",
	CsvSeparator:     ",",
	DecimalSeparator: ".",
	Encoding:         "utf-8"}

var defaultPreprocessOps = PreprocessConfig{
	Columns: []ColumnConfig{
		{Name: "id", Type: "int"},
		{Name: "name", Type: "string"},
		{Name: "price", Type: "float"},
	},
}

var defaultPrepfile = Prepfile{
	Data:       defaultDataSpec,
	Preprocess: defaultPreprocessOps,
}
