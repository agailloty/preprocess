package summary

type SummaryFile struct {
	Columns []ColumnSummary `toml:"columns"`
}

type ColumnSummary struct {
	Name                string          `toml:"name,omitempty"`
	RowCount            int             `toml:"RowCount,omitempty"`
	UniqueValueCount    int             `toml:"UniqueValueCount,omitempty"`
	UniqueValues        []string        `toml:"UniqueValues,omitempty"`
	UniqueValuesSummary []ValueKeyCount `toml:"UniqueValuesSummary,inline,omitempty"`

	Min     float64 `toml:"Min,omitempty"`
	Mean    float64 `toml:"Mean,omitempty"`
	Median  float64 `toml:"Median,omitempty"`
	Max     float64 `toml:"Max,omitempty"`
	Missing int     `toml:"Missing"`
}

type ValueKeyCount struct {
	Key   string `toml:"modality"`
	Count int    `toml:"count"`
}
