package common

type DataSpecs struct {
	Filename          string `toml:"filename"`
	CsvSeparator      string `toml:"csvseparator"`
	DecimalSeparator  string `toml:"decimalseparator"`
	Encoding          string `toml:"encoding"`
	MissingIdentifier string `toml:"missingIdentifier"`
}
