package config

type PreprocessOp struct {
	Op     string             `toml:"op"`
	Value  any                `toml:"value,omitempty"`
	Method string             `toml:"method,omitempty"`
	Bins   []BinningOperation `toml:"bins"`
}
