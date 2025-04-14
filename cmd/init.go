package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/agailloty/preprocess/dataset"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a pre-processing project in current folder",
	Run:   initialize,
}

func initialize(cmd *cobra.Command, args []string) {
	argLength := len(args)
	if argLength > 0 {
		sep := ","
		if argLength >= 2 {
			sep = args[1]
		}
		data := dataset.ReadDatasetColumns(args[0], sep)
		file, err := os.Create("Prepfile")

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s \n", args[0]))
		file.WriteString(fmt.Sprintf("%s \n", sep))

		for _, col := range data {
			switch v := col.(type) {
			case dataset.Float:
			case dataset.String:
			case dataset.Integer:
				content := fmt.Sprintf("%s : %T \n", v.Name, v)
				file.WriteString(content)
			}

		}

	}
}
