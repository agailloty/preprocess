package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/agailloty/preprocess/common"
	"github.com/agailloty/preprocess/utils"
)

func SetConfigFile() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error opening config.toml : %v", err)
	}
	defer file.Close()

	var config Prepfile
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("error decoding config.toml : %v", err)
	}
}

type Prepfile struct {
	Data        common.DataSpecs  `toml:"data"`
	Preprocess  PreprocessConfig  `toml:"preprocess"`
	PostProcess PostProcessConfig `toml:"postprocess"`
}

func InitDefaultPrepfile() Prepfile {
	return Prepfile{
		Data:        defaultDataSpec,
		Preprocess:  defaultPreprocessOps,
		PostProcess: postProcessDefaultConfig,
	}
}

func LoadConfigFromPrepfile(path string) (*Prepfile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening Prepfile : %w", err)
	}
	defer file.Close()

	var config Prepfile
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding Prepfile : %w", err)
	}

	return &config, nil
}

func MakeConfigFromCommandsArgs(dfSepc common.DataSpecs, columnList []string, operationList []string) *Prepfile {

	// I assume there is strict association between --column and --op
	nCol := len(columnList)
	if len(columnList) != len(operationList) {
		nCol = min(len(columnList), len(operationList))
	}

	//var config Config

	prepColumns := make([]ColumnConfig, nCol)

	for i := range nCol {
		prepColumns = append(prepColumns, makeColumnConfigFromArgs(columnList[i], operationList[i]))
	}
	newName := utils.AppendPrefixOrSuffix(dfSepc.Filename, "", "_cleaned")
	return &Prepfile{
		Data: dfSepc,
		Preprocess: PreprocessConfig{
			Columns: prepColumns},
		PostProcess: PostProcessConfig{Format: "csv", FileName: newName},
	}

}

func parseOperationFromArgs(operation string) (PreprocessOp, error) {
	providedOp := strings.Split(operation, ":")
	if len(providedOp) < 2 {
		return PreprocessOp{}, errors.New("not enough parameters")
	}
	opName, paramKvp := providedOp[0], providedOp[1]
	args := strings.Split(paramKvp, "=")

	if len(args) < 2 {
		return PreprocessOp{}, errors.New("no argument provided")
	}

	result := PreprocessOp{Op: opName}

	qualifier, value := string(args[0]), string(args[1])
	if qualifier == "method" {
		result.Method = value
	} else {
		result.Value = value
	}

	return result, nil
}

func makeColumnConfigFromArgs(columnName string, operation string) ColumnConfig {

	columnConfig := ColumnConfig{
		Name: columnName,
	}

	preprocessOp := make([]PreprocessOp, 1)
	op, err := parseOperationFromArgs(operation)
	if err == nil {
		preprocessOp = append(preprocessOp, op)
		columnConfig.Operations = &preprocessOp
	}

	return columnConfig
}
