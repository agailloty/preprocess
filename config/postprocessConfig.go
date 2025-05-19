package config

type PostProcessConfig struct {
	SummaryStats bool                `toml:"summarystats"`
	CountMissing bool                `toml:"countmissing"`
	DropColumns  []DropColumnEntry   `toml:"dropcolumns,omitempty"`
	SortDataset  *SortDatasetColumns `toml:"sortdataset,omitempty"`
	Format       string              `toml:"format"`
	FileName     string              `toml:"filename"`
}

type DropColumnEntry struct {
	Name string `toml:"name"`
}

type SortDatasetColumns struct {
	Descending bool `toml:"descending"`
}

var postProcessDefaultConfig = PostProcessConfig{
	SummaryStats: false,
	CountMissing: false,
	DropColumns:  nil,
}
