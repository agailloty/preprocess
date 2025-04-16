package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/spf13/cobra"
)

// Commande Cobra pour générer le fichier config.toml
var generateConfigCmd = &cobra.Command{
	Use:   "generate-config",
	Short: "Generate a toml config file based on the provided dataset",
	Run:   generateConfig,
}

func generateConfig(cmd *cobra.Command, args []string) {
	var configFile config.Config

	if len(args) == 0 {
		configFile = config.InitDefaultConfig()
	} else {
		filename := args[0]
		sep := ","
		if len(args) >= 2 {
			sep = args[1]
		}
		dataset := dataset.ReadDatasetColumns(filename, sep)
		var configColumns []config.ColumnConfig
		for _, col := range dataset {
			configColumns = append(configColumns,
				config.ColumnConfig{Name: col.GetName(), Type: col.GetType()})
		}

		configFile = config.Config{
			Data: config.DataConfig{
				File:      filename,
				Separator: sep,
				Columns:   configColumns,
			},
		}

	}
	file, err := os.Create("Prepfile.toml")
	if err != nil {
		log.Fatalf("Error generating Prefile.toml : %v", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(configFile); err != nil {
		log.Fatalf("An error occured during TOML enconding : %v", err)
	}

	fmt.Println("File successfully generated.")
}

func init() {
	rootCmd.AddCommand(generateConfigCmd)
}
