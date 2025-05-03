package config

type PostProcessConfig struct {
	Export         ExportConfig      `toml:"export"`
	SummaryStats   bool              `toml:"summary_stats"`
	CountMissing   bool              `toml:"count_missing"`
	ValidateSchema bool              `toml:"validate_schema"`
	SortBy         *SortConfig       `toml:"sort_by,omitempty"`
	DropColumns    []DropColumnEntry `toml:"drop_columns,omitempty"`
}

type ExportConfig struct {
	Format string `toml:"format"`
	Path   string `toml:"path"`
}

type SortConfig struct {
	Columns    []string `toml:"columns"`
	Descending bool     `toml:"descending"`
}

type DropColumnEntry struct {
	Name string `toml:"name"`
}

var postProcessDefaultConfig = PostProcessConfig{
	Export:       ExportConfig{Format: "csv", Path: "data.csv"},
	SummaryStats: false,
	CountMissing: false,
	DropColumns:  nil,
}
