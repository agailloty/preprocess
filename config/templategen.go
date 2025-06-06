package config

import (
	"log"
	"os"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
	"github.com/pelletier/go-toml/v2"
)

func InitializePrepfile(dfSpec common.DataSpecs, output string, templateOnly bool) {
	if dfSpec.Filename == "" || templateOnly {
		InitializeDefaultPrepfile(output)
		return
	}
	dataset := dataset.ReadDataFrame(dfSpec)
	var configColumns []ColumnConfig
	for _, col := range dataset.Columns {
		configColumns = append(configColumns,
			ColumnConfig{Name: col.GetName(), Type: col.GetType()})
	}

	newName := utils.AppendPrefixOrSuffix(dfSpec.Filename, "", "_cleaned")

	configFile := Prepfile{
		Data: dfSpec,
		Preprocess: PreprocessConfig{
			Columns: configColumns,
		},
		PostProcess: PostProcessConfig{
			Format: "csv", FileName: newName,
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

	log.Printf("%s successfully generated.\n", output)
}

func InitializeDefaultPrepfile(output string) {
	prepfile := InitDefaultPrepfile()
	utils.SerializeStruct(prepfile, output)
}
