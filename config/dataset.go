package config

type DatasetOperations struct {
	Operations *[]PreprocessOp `toml:"operations,omitempty"`
}
