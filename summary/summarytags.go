package summary

import "github.com/agailloty/preprocess/common"

type SummaryFile struct {
	Data        common.DataSpecs `toml:"data"`
	DataSummary DatasetSummary   `toml:"datasummary"`
	Columns     []ColumnSummary  `toml:"columns"`
}

type DatasetSummary struct {
	RowCount       int `toml:"rowscount"`
	ColumnCount    int `toml:"columnscount"`
	NumericColumns int `toml:"numericcolumns"`
	StringColumns  int `toml:"stringcolumns"`
}

type ColumnSummary struct {
	Name                string          `toml:"name,omitempty"`
	Type                string          `toml:"type"`
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
