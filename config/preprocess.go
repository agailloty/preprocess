package config

type PreprocessConfig struct {
	Columns           []ColumnConfig     `toml:"columns"`
	NumericOperations *DatasetOperations `mapstructure:"numerics,omitempty"`
	TextOperations    *DatasetOperations `mapstructure:"texts,omitempty"`
}
