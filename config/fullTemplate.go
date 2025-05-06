package config

var dataDefautConfig = DataConfig{
	File:      "data.csv",
	Separator: ",",
	Columns: []ColumnConfig{
		{Name: "id", Type: "int"},
		{Name: "name", Type: "string"},
		{Name: "price", Type: "float"},
	},
}
