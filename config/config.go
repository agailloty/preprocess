package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/agailloty/preprocess/utils"
	"github.com/spf13/viper"
)

func SetConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config Prepfile : %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshalling Prepfile : %v", err)
	}
}

type Config struct {
	Data        DataConfig        `toml:"data"`
	PostProcess PostProcessConfig `toml:"postprocess"`
}

func InitDefaultConfig() Config {
	return Config{
		Data:        dataDefautConfig,
		PostProcess: postProcessDefaultConfig,
	}
}

func LoadConfigFromPrepfile(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading Prepfile : %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("a mapping error occured : %w", err)
	}

	return &config, nil
}

func MakeConfigFromCommandsArgs(datapath string, sep string, columnList []string, operationList []string) *Config {

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
	newName := utils.AppendPrefixOrSuffix(datapath, "", "_cleaned")
	return &Config{
		Data:        DataConfig{File: datapath, Separator: sep, Columns: prepColumns},
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
		columnConfig.Preprocess = &preprocessOp
	}

	return columnConfig
}
