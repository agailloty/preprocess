package config

type DataSetSplit struct {
	Method              string    `toml:"method"`
	SplitNames          *[]string `toml:"split_names"`
	RandomSeed          *uint64   `toml:"random_seed"`
	TrainTestSplitRatio *float64  `toml:"train_test_split_ratio"`
}
