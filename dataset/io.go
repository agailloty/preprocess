package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func guessColumnType(column []string, nrow int, decimalSep string) DataSetColumn {
	stringCount := 0
	integerCount := 0
	floatCount := 0

	if nrow > len(column) {
		nrow = len(column) - 1
	}

	for i := range nrow {
		isSuccess, _ := tryParseInt(column[i])
		if isSuccess {
			integerCount += 1
			continue
		}
		isSuccess, _ = tryParseFloat(column[i], decimalSep)
		if isSuccess {
			floatCount += 1
			continue
		}
		stringCount += 1
	}

	var typedColumn DataSetColumn

	if stringCount > integerCount && stringCount > floatCount {
		var stringColumn []Nullable[string]
		for _, val := range column[1:] {
			stringColumn = append(stringColumn, Nullable[string]{IsValid: val != "", Value: val})
		}
		typedColumn = &String{Data: stringColumn, Name: column[0]}
	}

	if integerCount > stringCount && integerCount > floatCount {
		var intColumn []Nullable[int]
		for _, val := range column[1:] {
			isValid, parsedVal := tryParseInt(val)
			intColumn = append(intColumn, Nullable[int]{IsValid: isValid && val != "", Value: int(parsedVal)})
		}

		typedColumn = &Integer{Data: intColumn, Name: column[0]}
	}

	if floatCount > stringCount && floatCount > integerCount {
		var floatColumn []Nullable[float64]
		for _, val := range column[1:] {
			isValid, parsedVal := tryParseFloat(val, decimalSep)
			floatColumn = append(floatColumn, Nullable[float64]{IsValid: isValid && val != "", Value: float64(parsedVal)})
		}
		typedColumn = &Float{Data: floatColumn, Name: column[0]}
	}

	return typedColumn
}

func tryParseInt(val string) (bool, int64) {
	isSuccess := false
	value, err := strconv.ParseInt(val, 10, 32)
	if err == nil {
		isSuccess = true
	}
	return isSuccess, value
}

func tryParseFloat(val string, decimalSep string) (bool, float64) {
	isSuccess := false
	val = strings.Replace(val, decimalSep, ".", 1)
	value, err := strconv.ParseFloat(val, 64)
	if err == nil {
		isSuccess = true
	}
	return isSuccess, value
}

func convertToTypedColumns(data [][]string, decimalSep string) []DataSetColumn {
	rowLength := len(data)
	colLength := len(data[0])
	var columns []DataSetColumn
	for i := 0; i < colLength; i++ {
		var column []string
		for j := 0; j < rowLength; j++ {
			column = append(column, data[j][i])
		}
		columns = append(columns, guessColumnType(column, int(0.2*float64(rowLength)), decimalSep))
	}

	return columns
}

func readDatasetColumns(filename string, sep string, decimalSep string) []DataSetColumn {
	data := readAllLines(filename, sep)
	return convertToTypedColumns(data, decimalSep)
}

func readAllLines(filepath string, sep string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening : %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if len(sep) == 1 {
		reader.Comma = rune(sep[0])
	} else {
		log.Fatalf("Separator must be a unique character")
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv CSV : %v", err)
	}

	return records
}
