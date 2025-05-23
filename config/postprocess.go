package config

type PostProcessConfig struct {
	DropColumns *[]string           `toml:"dropcolumns,omitempty"`
	SortDataset *SortDatasetColumns `toml:"sortdataset,omitempty"`
	Format      string              `toml:"format"`
	FileName    string              `toml:"filename"`
}

type SortDatasetColumns struct {
	Descending bool `toml:"descending"`
}

var postProcessDefaultConfig = PostProcessConfig{
	Format:      "csv",
	FileName:    "data_cleaned.csv",
	DropColumns: nil,
}
