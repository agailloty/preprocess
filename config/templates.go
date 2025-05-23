package config

import "github.com/agailloty/preprocess/common"

var dataDefautConfig = DataConfig{
	DataSpecs: common.DataSpecs{Filename: "data.csv",
		CsvSeparator:     ",",
		DecimalSeparator: ".",
		Encoding:         "utf8"},
	Columns: []ColumnConfig{
		{Name: "id", Type: "int"},
		{Name: "name", Type: "string"},
		{Name: "price", Type: "float"},
	},
}
