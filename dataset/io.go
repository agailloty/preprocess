package dataset

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GuessTypes(data [][]string) []DataSetColumn {
	rowLength := len(data)
	colLength := len(data[0])
	var columns []DataSetColumn
	for i := 0; i < colLength; i++ {
		var column []string
		for j := 0; j < rowLength; j++ {
			column = append(column, data[j][i])
		}
		columns = append(columns, String{data: column, name: "user"})

	}

	return columns
}

func ReadAllLines(filepath string, sep string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), sep))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
