package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/agailloty/preprocess/common"
)

type DataFrame struct {
	common.DataSpecs
	Name         string
	Columns      []DataSetColumn
	RowsCount    int
	ColumnsCount int
}

func DisplayColumn(column DataSetColumn, n int) {
	switch v := column.(type) {
	case *Float:
		log.Printf("%s (float) \n", v.Name)
		for i := range n {
			log.Printf("%.2f ", (v.Data)[i].Value)
		}
	case *String:
		log.Printf("%s (string) \n", v.Name)
		for i := range n {
			log.Printf("%s ", (v.Data)[i].Value)
		}
	case *Integer:
		log.Printf("%s (int) \n", v.Name)
		for i := range n {
			log.Printf("%d ", (v.Data)[i].Value)
		}
	}
}

func ReadDataFrame(dfSpec common.DataSpecs) DataFrame {
	columns := readDatasetColumns(dfSpec)
	ext := filepath.Ext(dfSpec.Filename)
	dfName := strings.TrimSuffix(dfSpec.Filename, ext)
	return DataFrame{
		DataSpecs:    dfSpec,
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

func (d *DataFrame) DeleteColumn(column DataSetColumn) {
	found, index := findIndex(d.Columns, column)
	if found {
		d.Columns = slices.Delete(d.Columns, index, index+1)
	}

}

func (d *DataFrame) DeleteColumnByName(columnName string) {
	found, index := findFirstIndexByName(d.Columns, columnName)
	if found {
		d.Columns = slices.Delete(d.Columns, index, index+1)
		d.ColumnsCount = len(d.Columns)
	}

}

func findIndex(columns []DataSetColumn, colToFind DataSetColumn) (found bool, index int) {
	index = -1
	for i, column := range columns {
		if colToFind == column {
			index = i
			found = true
			break
		}
	}
	return found, index
}

func findFirstIndexByName(columns []DataSetColumn, colName string) (found bool, index int) {
	index = -1
	for i, column := range columns {
		if column.GetName() == colName {
			index = i
			found = true
			break
		}
	}
	return found, index
}
