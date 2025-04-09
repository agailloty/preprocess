package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "preprocess",
	Short: "preprocess is a fast data analysis preprocessing tool.",
	Long: `preprocess is a fast data analysis preprocessing tool built with Go.
			Complete documentation is available at 
			https://github.com/agailloty/preprocess`,
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
