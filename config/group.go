package config

type GroupOption struct {
	Values []string `toml:"values"`
	Name   string   `toml:"name"`
}
