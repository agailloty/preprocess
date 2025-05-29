package config

type PreprocessConfig struct {
	Columns           []ColumnConfig     `toml:"columns"`
	NumericOperations *DatasetOperations `toml:"numerics,omitempty"`
	TextOperations    *DatasetOperations `toml:"texts,omitempty"`
}
