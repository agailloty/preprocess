package config

type PostProcessConfig struct {
	Export       ExportConfig        `toml:"export"`
	SummaryStats bool                `toml:"summarystats"`
	CountMissing bool                `toml:"countmissing"`
	DropColumns  []DropColumnEntry   `toml:"dropcolumns,omitempty"`
	SortDataset  *SortDatasetColumns `toml:"sortdataset,omitempty"`
}

type ExportConfig struct {
	Format string `toml:"format"`
	Path   string `toml:"path"`
}

type DropColumnEntry struct {
	Name string `toml:"name"`
}

type SortDatasetColumns struct {
	Descending bool `toml:"descending"`
}

var postProcessDefaultConfig = PostProcessConfig{
	Export:       ExportConfig{Format: "csv", Path: "data.csv"},
	SummaryStats: false,
	CountMissing: false,
	DropColumns:  nil,
}
