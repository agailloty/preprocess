package config

type DatasetOperations struct {
	ExcludeCols []string        `toml:"exclude_columns"`
	Operations  *[]PreprocessOp `toml:"operations,omitempty"`
}
