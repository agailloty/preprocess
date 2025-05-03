package dataset

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DataFrame struct {
	Name         string
	Columns      []DataSetColumn
	RowsCount    int
	ColumnsCount int
}

func DisplayColumn(column DataSetColumn, n int) {
	switch v := column.(type) {
	case Float:
		fmt.Printf("%s (float) \n", v.Name)
		for i := range n {
			fmt.Printf("%.2f ", v.Data[i])
		}
	case String:
		fmt.Printf("%s (string) \n", v.Name)
		for i := range n {
			fmt.Printf("%s ", v.Data[i])
		}
	case Integer:
		fmt.Printf("%s (int) \n", v.Name)
		for i := range n {
			fmt.Printf("%d ", v.Data[i])
		}
	}
}

func ReadDataFrame(filename string, sep string) DataFrame {
	columns := readDatasetColumns(filename, sep)
	ext := filepath.Ext(filename)
	dfName := strings.TrimSuffix(filename, ext)
	return DataFrame{
		Name:         dfName,
		Columns:      columns,
		RowsCount:    columns[0].Length(),
		ColumnsCount: len(columns),
	}
}

func (d DataFrame) SaveToCSV(filepath string, sep string) error {
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
	for i := 0; i < d.RowsCount; i++ {
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
