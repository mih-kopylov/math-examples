package main

import (
	"errors"
	"os"

	"github.com/joomcode/errorx"
	"gopkg.in/yaml.v3"
)

var (
	ErrParamsNamespace         = errorx.NewNamespace("params")
	ErrFailCreateDefaultParams = errorx.NewType(ErrParamsNamespace, "FailCreateDefaultParams")
	ErrFailReadParams          = errorx.NewType(ErrParamsNamespace, "FailReadParams")
)

const configFileName = "math-examples.yaml"

type AppParams struct {
	Profiles map[string]ProfileParams
}

type ProfileParams struct {
	ExamplesCount                   int               `yaml:"examplesCount"`
	MinBoundary                     int               `yaml:"minBoundary"`
	MaxBoundary                     int               `yaml:"maxBoundary"`
	OperandsCount                   int               `yaml:"operandsCount"`
	ShowCorrectAnswerAfter          CorrectAnswerMode `yaml:"showCorrectAnswerAfter"`
	AvailableOperands               []string          `yaml:"availableOperands"`
	AvailableMultiplicationOperands []string          `yaml:"availableMultiplicationOperands"`
	AvailableOperationTypes         []OperationType   `yaml:"availableOperationTypes"`
}

type OperationType string

const (
	PlusOperationType     OperationType = "plus"
	MinusOperationType    OperationType = "minus"
	MultiplyOperationType OperationType = "multiply"
	DivideOperationType   OperationType = "divide"
)

type CorrectAnswerMode string

const (
	AfterEach CorrectAnswerMode = "each"
	AfterAll  CorrectAnswerMode = "all"
)

func ReadParams() (*AppParams, error) {
	_, err := os.Stat(configFileName)
	if errors.Is(err, os.ErrNotExist) {
		defaultParams := AppParams{
			Profiles: map[string]ProfileParams{
				"Имя": {
					ExamplesCount:          10,
					MinBoundary:            0,
					MaxBoundary:            100,
					OperandsCount:          2,
					ShowCorrectAnswerAfter: AfterEach,
					AvailableOperationTypes: []OperationType{
						PlusOperationType,
						MinusOperationType,
						MultiplyOperationType,
						DivideOperationType,
					},
					AvailableOperands:               []string{"1-100"},
					AvailableMultiplicationOperands: []string{"1-9"},
				},
			},
		}

		bytes, err := yaml.Marshal(defaultParams)
		if err != nil {
			return nil, ErrFailCreateDefaultParams.Wrap(err, "failed to marshall default params")
		}

		err = os.WriteFile(configFileName, bytes, os.ModePerm)
		if err != nil {
			return nil, ErrFailCreateDefaultParams.Wrap(err, "failed to write default params to file")
		}
	}

	bytes, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, ErrFailReadParams.Wrap(err, "failed to read params from file")
	}

	var result AppParams

	err = yaml.Unmarshal(bytes, &result)
	if err != nil {
		return nil, ErrFailReadParams.Wrap(err, "failed to unmarshall params")
	}

	return &result, nil
}
