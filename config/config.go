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
	Data DataConfig `toml:"data"`
}

type DataConfig struct {
	File      string         `toml:"file"`
	Separator string         `toml:"separator"`
	Columns   []ColumnConfig `toml:"columns"`
}

type ColumnConfig struct {
	Name string `toml:"name"`
	Type string `toml:"type"`
}

func InitDefaultConfig() Config {
	return Config{
		Data: DataConfig{
			File:      "data.csv",
			Separator: ",",
			Columns: []ColumnConfig{
				{Name: "id", Type: "int"},
				{Name: "name", Type: "string"},
				{Name: "price", Type: "float"},
				{Name: "available", Type: "bool"},
			},
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading Prepfile : %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("A mapping error occured : %w", err)
	}

	return &config, nil
}
