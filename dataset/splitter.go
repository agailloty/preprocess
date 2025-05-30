package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

type SplitSpec struct {
	Name string
	Rows []int
}

type SplitterFunc func(df *DataFrame, args ...any) []SplitSpec

func (df *DataFrame) SaveSplittedDataframeToCSV(splitF SplitterFunc) {
	splitSpecs := splitF(df)

	for _, spec := range splitSpecs {
		go writeDataframeSubset(df, spec)
	}
}

func writeDataframeSubset(df *DataFrame, spec SplitSpec) {
	if spec.Name == "" {
		spec.Name = df.Name + time.Now().Format("mmssnn")
	}
	spec.Name = spec.Name + ".csv"
	if len(spec.Rows) == 0 {
		return
	}

	if df.writeCsv(spec.Name, spec) != nil {
		log.Printf("Error while writing split %s \n", spec.Name)
	}
}

func (d *DataFrame) writeCsv(filepath string, spec SplitSpec) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	var headers []string
	for _, col := range d.Columns {
		headers = append(headers, col.GetName())
	}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write rows
	for _, i := range spec.Rows {
		var row []string
		for _, col := range d.Columns {
			row = append(row, col.ValueAt(i))
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
