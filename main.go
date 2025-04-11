package main

import (
	"fmt"

	"github.com/agailloty/preprocess/cmd"
	"github.com/agailloty/preprocess/dataset"
)

func main() {
	cmd.Execute()
	data := dataset.ReadAllLines("fifa_players_22.csv", ",")
	guessedTypes := dataset.ReadDatasetColumns(data)

	for _, dt := range guessedTypes {
		fmt.Printf("%T \n", dt)
	}
}
