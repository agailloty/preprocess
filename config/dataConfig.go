package config

type DataConfig struct {
	File              string             `toml:"file"`
	Separator         string             `toml:"separator"`
	Columns           []ColumnConfig     `toml:"columns"`
	NumericOperations *DatasetOperations `mapstructure:"numericColumns,omitempty"`
	TextOperations    *DatasetOperations `mapstructure:"textColumns,omitempty"`
}

type PreprocessOp struct {
	Op      string        `toml:"op"`
	Value   any           `toml:"value,omitempty"`
	Method  string        `toml:"method,omitempty"`
	BinSpec *BinningSpecs `toml:"binSpecs"`
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
