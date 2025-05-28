package config

type ColumnConfig struct {
	Name       string          `toml:"name"`
	Type       string          `toml:"type"`
	NewName    string          `toml:"new_name,omitempty"`
	Operations *[]PreprocessOp `toml:"operations,omitempty"`
}
