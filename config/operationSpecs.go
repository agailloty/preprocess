package config

type BinningSpecs struct {
	Bins      []BinningOperation `toml:"bins"`
	NewColumn string             `toml:"newColumn"`
}

type BinningOperation struct {
	Lower float32 `toml:"lower"`
	Upper float32 `toml:"upper"`
	Label string  `toml:"label"`
}
