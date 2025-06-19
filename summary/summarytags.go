package summary

import "github.com/agailloty/preprocess/common"

type SummaryFile struct {
	Data        common.DataSpecs `toml:"data"`
	DataSummary DatasetSummary   `toml:"data_summary"`
	Columns     []ColumnSummary  `toml:"columns"`
}

type DatasetSummary struct {
	RowCount       int `toml:"rows_count"`
	ColumnCount    int `toml:"columns_count"`
	NumericColumns int `toml:"numeric_columns"`
	StringColumns  int `toml:"string_columns"`
}

type ColumnSummary struct {
	Name                string                 `toml:"name,omitempty"`
	Type                string                 `toml:"type"`
	RowCount            int                    `toml:"rows_count,omitempty"`
	UniqueValueCount    int                    `toml:"unique_count,omitempty"`
	UniqueValues        []string               `toml:"unique_values,omitempty"`
	UniqueValuesSummary []common.ValueKeyCount `toml:"summary,inline,omitempty"`
	IsDeleted           bool
	IsAdded             bool
	IsAltered           bool
	IsIdentical         bool
	AddedStringValues   []string
	RemovedStringValues []string

	Min     float64 `toml:"min,omitempty"`
	Mean    float64 `toml:"mean,omitempty"`
	Median  float64 `toml:"median,omitempty"`
	Max     float64 `toml:"max,omitempty"`
	Missing int     `toml:"missing"`

	OldStats *NumericStats
	NewStats *NumericStats
}

type NumericStats struct {
	RowCount int
	Missing  int
	Mean     float64
	Median   float64
	Min      float64
	Max      float64
}
