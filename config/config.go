package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func SetConfigFile() {
	viper.SetConfigName("config") // sans extension
	viper.SetConfigType("toml")
	viper.AddConfigPath(".") // ou un chemin spécifique

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier de config : %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Erreur lors du mapping de la configuration : %v", err)
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
