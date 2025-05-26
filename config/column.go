package config

type ColumnConfig struct {
	Name       string          `toml:"name"`
	Type       string          `toml:"type"`
	NewName    string          `toml:"newName,omitempty"`
	Operations *[]PreprocessOp `toml:"operations,omitempty"`
}
