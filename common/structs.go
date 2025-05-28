package common

type DataSpecs struct {
	Filename          string `toml:"filename"`
	CsvSeparator      string `toml:"csv_separator"`
	DecimalSeparator  string `toml:"decimal_separator"`
	Encoding          string `toml:"encoding"`
	MissingIdentifier string `toml:"missing_identifier"`
}

type ValueKeyCount struct {
	Key   string `toml:"modality"`
	Count int    `toml:"count"`
}
