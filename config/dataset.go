package config

type PreprocessConfig struct {
	Columns           []ColumnConfig     `toml:"columns"`
	NumericOperations *DatasetOperations `mapstructure:"numerics,omitempty"`
	TextOperations    *DatasetOperations `mapstructure:"texts,omitempty"`
}

type PreprocessOp struct {
	Op     string             `toml:"op"`
	Value  any                `toml:"value,omitempty"`
	Method string             `toml:"method,omitempty"`
	Bins   []BinningOperation `toml:"bins"`
}

type ColumnConfig struct {
	Name       string          `toml:"name"`
	Type       string          `toml:"type"`
	NewName    string          `toml:"newName,omitempty"`
	Preprocess *[]PreprocessOp `toml:"preprocess,omitempty"`
}

type DatasetOperations struct {
	Preprocess *[]PreprocessOp `toml:"preprocess,omitempty"`
}
