package config

type PreprocessOp struct {
	Op                             string             `toml:"op"`
	Value                          any                `toml:"value,omitempty"`
	Method                         string             `toml:"method,omitempty"`
	Bins                           []BinningOperation `toml:"bins"`
	DummyPrefixColName             bool               `toml:"dummy_prefix"`
	DummyDropLast                  bool               `toml:"dummy_droplast"`
	DummyContinueWithTooManyValues bool               `toml:"continue_with_toomany"`
	ExcludeCols                    []string           `toml:"exclude_columns"`
}
