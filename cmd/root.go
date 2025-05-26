package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var datafilename string
var csvseparator string
var decimalSeparator string
var encoding string
var prepfilePath string

var rootCmd = &cobra.Command{
	Use:   "preprocess",
	Short: PREPROCESS_SHORT_DESCRIPTION,
	Long:  PREPROCESS_LONG_DESCRIPTION,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setDataSpecFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&datafilename, "data", "d", "", "Path to the dataset")
	cmd.Flags().StringVarP(&csvseparator, "sep", "s", ",", "Separator for csv file")
	cmd.Flags().StringVarP(&decimalSeparator, "dsep", "m", ".", "Decimal separator")
	cmd.Flags().StringVarP(&encoding, "encoding", "e", "utf-8", "Character encoding")
}

func setPrepfileFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&prepfilePath, "prepfile", "f", "Prepfile.toml", "Path to the configuration file")
}
