package summary

type SummaryFile struct {
	Columns []ColumnSummary `toml:"columns"`
}

type ColumnSummary struct {
	Name                string          `toml:"name"`
	RowCount            int             `toml:"RowCount"`
	UniqueValueCount    int             `toml:"UniqueValueCount"`
	UniqueValues        []string        `toml:"UniqueValues"`
	UniqueValuesSummary []ValueKeyCount `toml:"UniqueValuesSummary,inline"`
}

type ValueKeyCount struct {
	Key   string `toml:"modality"`
	Count int    `toml:"count"`
}
