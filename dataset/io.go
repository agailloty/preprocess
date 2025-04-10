package dataset

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
