package config

type BinningSpecs struct {
	Bins      []BinningOperation `toml:"bins"`
	NewColumn string             `toml:"new_column"`
}

type BinningOperation struct {
	Lower float64 `toml:"lower"`
	Upper float64 `toml:"upper"`
	Label string  `toml:"label"`
}
