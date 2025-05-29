package cmd

import (
	"time"

	"log"

	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/config"
	"github.com/spf13/cobra"
)

var output string
var templateOnly bool

// Commande Cobra pour générer le fichier config.toml
var initCmd = &cobra.Command{
	Use:   "init",
	Short: INIT_SHORT_DESCRIPTION,
	Long:  INIT_LONG_DESCRIPTION,
	Args:  cobra.NoArgs,
}

func initConfig(cmd *cobra.Command, args []string) {
	start := time.Now()
	if initCmd.Flags().NFlag() == 0 {
		initCmd.Help()
		return
	}
	if decimalSeparator == "" {
		decimalSeparator = "."
	}

	if encoding == "" {
		encoding = "utf-8"
	}
	dfSpec := common.DataSpecs{
		Filename:         datafilename,
		CsvSeparator:     csvseparator,
		DecimalSeparator: decimalSeparator,
		Encoding:         encoding,
	}
	config.InitializePrepfile(dfSpec, output, templateOnly)

	elapsed := time.Since(start)
	log.Printf("Finished in : %s\n", elapsed)
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Run = initConfig
	setDataSpecFlags(initCmd)
	initCmd.Flags().StringVarP(&output, "output", "o", "Prepfile.toml", "Output name for Prepfile")
	initCmd.Flags().BoolVarP(&templateOnly, "template", "t", false, "Generate example Prepfile.toml template")
}
