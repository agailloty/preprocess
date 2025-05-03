package config

type DataConfig struct {
	File      string         `toml:"file"`
	Separator string         `toml:"separator"`
	Columns   []ColumnConfig `toml:"columns"`
}

type ColumnConfig struct {
	Name    string `toml:"name"`
	Type    string `toml:"type"`
	NewName string `toml:"newName"`
}

var dataDefautConfig = DataConfig{
	File:      "data.csv",
	Separator: ",",
	Columns: []ColumnConfig{
		{Name: "id", Type: "int"},
		{Name: "name", Type: "string"},
		{Name: "price", Type: "float"},
		{Name: "available", Type: "bool"},
	},
}
