package operations

import (
	"fmt"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

type opSummary struct {
	totalOperations   int
	numericOperations int
	stringOperations  int
}

// Count the number of all operations that are about to be applied on the dataset.
// It first count the number of operations applied on whole dataset plus the numer of operations
// applied on each column
func summarizeOperations(prepfile *config.Prepfile, df *dataset.DataFrame) opSummary {
	numericOperations := 0
	stringOperations := 0

	nbNumericColums := 0
	nbTextColumns := 0

	for _, col := range df.Columns {
		if col.GetType() == "string" {
			nbTextColumns++
		} else {
			nbNumericColums++
		}
	}

	// When operations are applied on whole numeric columns then we consider
	// they are applied to each column individually except the columns marked as
	// excluded
	if prepfile.Preprocess.NumericOperations.Operations != nil {
		// Make sure excluded numeric columns really exist
		excluded := []string{}
		for _, colName := range prepfile.Preprocess.NumericOperations.ExcludeCols {
			if exists(df.Columns, colName, "int") || exists(df.Columns, colName, "float") {
				excluded = append(excluded, colName)
			}
		}
		nbDatasetNumericOps := len(*prepfile.Preprocess.NumericOperations.Operations)

		numericOperations = (nbNumericColums - len(excluded)) * nbDatasetNumericOps
	}

	if prepfile.Preprocess.TextOperations.Operations != nil {
		excluded := []string{}
		for _, colName := range prepfile.Preprocess.TextOperations.ExcludeCols {
			if exists(df.Columns, colName, "string") {
				excluded = append(excluded, colName)
			}
		}
		nbTextOps := len(*prepfile.Preprocess.TextOperations.Operations)
		stringOperations = (nbTextColumns - len(excluded)) * nbTextOps
	}

	for _, col := range prepfile.Preprocess.Columns {
		if col.Operations != nil {
			if col.Type == "string" {
				stringOperations += len(*col.Operations)
			} else {
				numericOperations += len(*col.Operations)
			}
		}
	}

	return opSummary{
		numericOperations: numericOperations,
		stringOperations:  stringOperations,
		totalOperations:   numericOperations + stringOperations,
	}

}

func (s *opSummary) logOperationStats() {
	fmt.Printf("Total operations applied : %d \n", s.totalOperations)
	fmt.Printf("Operations on numeric columns : %d \n", s.numericOperations)
	fmt.Printf("Operations on string columns : %d \n", s.stringOperations)
}

func exists(columns []dataset.DataSetColumn, colName string, colType string) bool {
	result := false
	for _, col := range columns {
		if col.GetName() == colName && col.GetType() == colType {
			result = true
			break
		}
	}

	return result
}
