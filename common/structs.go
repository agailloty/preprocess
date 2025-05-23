package common

type DataSpecs struct {
	Filename         string `toml:"filename"`
	CsvSeparator     string `toml:"separator"`
	DecimalSeparator string `toml:"decimalseparator"`
	Encoding         string `toml:"encoding"`
}
