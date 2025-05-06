package cmd

import (
	"github.com/agailloty/preprocess/config"
	"github.com/spf13/cobra"
)

var datafilename string
var separator string
var output string
var templateOnly bool

// Commande Cobra pour générer le fichier config.toml
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a Prepfile",
	Args:  cobra.NoArgs,
}

func initConfig(cmd *cobra.Command, args []string) {
	if initCmd.Flags().NFlag() == 0 {
		initCmd.Help()
		return
	}
	config.InitializePrepfile(datafilename, separator, output, templateOnly)
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(initCmd)
	initCmd.Run = initConfig
	initCmd.Flags().StringVarP(&datafilename, "file", "f", "", "Path to the dataset")
	initCmd.Flags().StringVarP(&separator, "sep", "s", ",", "Separator for csv file")
	initCmd.Flags().StringVarP(&output, "output", "o", "Prepfile.toml", "Output name for Prepfile")
	initCmd.Flags().BoolVarP(&templateOnly, "template", "t", false, "Generate example Prepfile.toml template")
}
