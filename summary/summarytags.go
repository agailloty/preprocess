package summary

type SummaryFile struct {
	Filename       string          `toml:"filename"`
	RowCount       int             `toml:"rowscount"`
	ColumnCount    int             `toml:"columnscount"`
	NumericColumns int             `toml:"numericcolumns"`
	StringColumns  int             `toml:"stringcolumns"`
	Columns        []ColumnSummary `toml:"columns"`
}

type ColumnSummary struct {
	Name                string          `toml:"name,omitempty"`
	RowCount            int             `toml:"rowscount,omitempty"`
	UniqueValueCount    int             `toml:"uniquevaluescount,omitempty"`
	UniqueValues        []string        `toml:"uniquevalues,omitempty"`
	UniqueValuesSummary []ValueKeyCount `toml:"summary,inline,omitempty"`

	Min     float64 `toml:"min,omitempty"`
	Mean    float64 `toml:"mean,omitempty"`
	Median  float64 `toml:"median,omitempty"`
	Max     float64 `toml:"max,omitempty"`
	Missing int     `toml:"missing"`
}

type ValueKeyCount struct {
	Key   string `toml:"modality"`
	Count int    `toml:"count"`
}
