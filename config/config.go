package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func SetConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config Prepfile : %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshalling Prepfile : %v", err)
	}
}

type Config struct {
	Data        DataConfig        `toml:"data"`
	PostProcess PostProcessConfig `toml:"postprocess"`
}

func InitDefaultConfig() Config {
	return Config{
		Data:        dataDefautConfig,
		PostProcess: postProcessDefaultConfig,
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading Prepfile : %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("a mapping error occured : %w", err)
	}

	return &config, nil
}
