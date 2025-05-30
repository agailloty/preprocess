package config

type PostProcessConfig struct {
	DropColumns  *[]string           `toml:"drop_columns,omitempty"`
	SortDataset  *SortDatasetColumns `toml:"sort_dataset,omitempty"`
	Format       string              `toml:"format"`
	FileName     string              `toml:"filename"`
	DataSetSplit *DataSetSplit       `toml:"dataset_split"`
}

type SortDatasetColumns struct {
	Descending bool `toml:"descending"`
}

var postProcessDefaultConfig = PostProcessConfig{
	Format:      "csv",
	FileName:    "data_cleaned.csv",
	DropColumns: nil,
}
