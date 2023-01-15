package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type exampleParams struct {
	ExamplesCount           int               `yaml:"examplesCount"`
	MinBoundary             int               `yaml:"minBoundary"`
	MaxBoundary             int               `yaml:"maxBoundary"`
	OperandsCount           int               `yaml:"operandsCount"`
	ShowCorrectAnswerAfter  correctAnswerMode `yaml:"showCorrectAnswerAfter"`
	AvailableOperands       []int             `yaml:"availableOperands"`
	AvailableOperationTypes []operationType   `yaml:"availableOperationTypes"`
}

type operationType string

const (
	plusOperationType  operationType = "plus"
	minusOperationType operationType = "minus"
)

type correctAnswerMode string

const (
	afterEach correctAnswerMode = "each"
	afterAll  correctAnswerMode = "all"
)

func readParams() (*exampleParams, error) {
	_, err := os.Stat(configFileName)
	if errors.Is(err, os.ErrNotExist) {
		defaultParams := exampleParams{
			ExamplesCount:           10,
			MinBoundary:             0,
			MaxBoundary:             9,
			OperandsCount:           2,
			ShowCorrectAnswerAfter:  afterEach,
			AvailableOperationTypes: []operationType{plusOperationType, minusOperationType},
			AvailableOperands:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}

		bytes, err := yaml.Marshal(defaultParams)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(configFileName, bytes, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	bytes, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	var result exampleParams

	err = yaml.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
