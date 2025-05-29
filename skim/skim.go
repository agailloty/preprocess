package skim

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/agailloty/preprocess/common"
)

const (
	maxColumns = 10
	maxRows    = 30
)

func SkimDf(dfSpec common.DataSpecs) {
	file, err := os.Open(dfSpec.Filename)
	if err != nil {
		log.Printf("error opening %s : %s", dfSpec.Filename, err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if len(dfSpec.CsvSeparator) == 1 {
		reader.Comma = rune(dfSpec.CsvSeparator[0])
	} else {
		log.Fatalf("Separator must be a unique character")
	}

	var records [][]string
	rowCount := 0

	for i := 0; i < maxRows; i++ {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading line", err)
			return
		}
		records = append(records, row)
		rowCount++
	}

	if len(records) == 0 {
		log.Println("Empty file.")
		return
	}

	totalRows := rowCount
	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading line during count", err)
			return
		}
		totalRows++
	}

	numCols := len(records[0])

	log.Printf("Columns: %d | Total rows: %d\n\n", numCols, totalRows)

	startCols := []int{}
	endCols := []int{}

	if numCols > maxColumns {
		keep := maxColumns - 1
		half := keep / 2
		for i := 0; i < half; i++ {
			startCols = append(startCols, i)
		}
		for i := numCols - (keep - half); i < numCols; i++ {
			endCols = append(endCols, i)
		}
	} else {
		for i := range records[0] {
			startCols = append(startCols, i)
		}
	}

	colIndexes := append(startCols, -1)
	colIndexes = append(colIndexes, endCols...)

	colWidths := make(map[int]int)
	for _, row := range records {
		for _, idx := range colIndexes {
			var val string
			if idx == -1 {
				val = "[...]"
			} else if idx < len(row) {
				val = row[idx]
			}
			if len(val) > colWidths[idx] {
				colWidths[idx] = len(val)
			}
		}
	}

	printRow := func(row []string) {
		for _, idx := range colIndexes {
			var val string
			if idx == -1 {
				val = "[...]"
			} else if idx < len(row) {
				val = row[idx]
			}
			log.Printf("%-*s  ", colWidths[idx], val)
		}
		log.Println()
	}

	printRow(records[0])
	for _, idx := range colIndexes {
		fmt.Print(strings.Repeat("-", colWidths[idx]), "  ")
	}
	log.Println()

	for _, row := range records[1:] {
		printRow(row)
	}
}
