package config

import (
	"fmt"
	"log"
	"os"

	"github.com/agailloty/preprocess/dataset"
	"github.com/pelletier/go-toml/v2"
)

func InitializePrepfile(filename string, sep string, output string, templateOnly bool) {
	if filename == "" || templateOnly {
		InitializeDefaultPrepfile(output)
		return
	}
	dataset := dataset.ReadDataFrame(filename, sep)
	var configColumns []ColumnConfig
	for _, col := range dataset.Columns {
		configColumns = append(configColumns,
			ColumnConfig{Name: col.GetName(), Type: col.GetType()})
	}

	configFile := Config{
		Data: DataConfig{
			File:      filename,
			Separator: sep,
			Columns:   configColumns,
		},
	}

	file, err := os.Create(output)
	if err != nil {
		log.Fatalf("Error generating Prefile.toml : %v", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(configFile); err != nil {
		log.Fatalf("An error occured during TOML enconding : %v", err)
	}

	fmt.Printf("%s successfully generated.", output)
}

func InitializeDefaultPrepfile(output string) {
	configFile := InitDefaultConfig()
	file, err := os.Create(output)
	if err != nil {
		log.Fatalf("Error generating Prefile.toml : %v", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(configFile); err != nil {
		log.Fatalf("An error occured during TOML enconding : %v", err)
	}
	fmt.Printf("Template %s successfully generated.", output)
}
