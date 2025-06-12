package operations

import (
	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
	"github.com/agailloty/preprocess/utils"
)

type groupOperation struct {
	df      *dataset.DataFrame
	col     *dataset.String
	options []config.GroupOption
}

func (o groupOperation) run() []operationResult {
	results := []operationResult{}
	results = append(results, op_group(o))
	return results
}

func op_group(specs groupOperation) operationResult {
	for i, data := range specs.col.Data {
		specs.col.Data[i].Value = findReplace(data.Value, specs.options)
	}

	return operationResult{isSuccess: true}
}

func findReplace(value string, options []config.GroupOption) string {
	result := value
	for _, replacement := range options {
		if utils.Contains(replacement.Values, value) {
			result = replacement.Name
		}
	}
	return result
}
