package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/joomcode/errorx"
	"golang.org/x/exp/slices"
)

var (
	ErrUnableToGenerateOperation = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateOperation")
)

type OperationGenerator struct {
	random  *rand.Rand
	profile *ProfileParams
}

func NewOperationGenerator(profile *ProfileParams) Generator {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &OperationGenerator{
		random:  random,
		profile: profile,
	}
}

func (g *OperationGenerator) GenerateExample(params *ProfileParams, distribution *Distribution[int]) (Example, error) {
	for tryNumber := 0; ; tryNumber++ {
		result, err := g.tryGenerateExample(params, distribution)
		if err == nil {
			distribution.Add(result.Answer())
			return result, nil
		}

		if tryNumber > 1000 {
			return nil, ErrUnableToGenerateExample.New("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию.")
		}
	}
}

func (g *OperationGenerator) tryGenerateExample(params *ProfileParams, distribution *Distribution[int]) (
	Example, error,
) {
	initialValue := g.randomOperand(params.AvailableOperands)
	var operations []Operation

	for i := 0; i < params.OperandsCount-1; i++ {
		operation, err := g.generateOperationWithinBounds(initialValue, operations, params)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}
	result := NewOperationExample(initialValue, operations)

	if distribution.IsTooFrequent(result.Answer()) {
		return nil, ErrTooFrequentExampleAnswer.NewWithNoMessage()
	}

	return result, nil
}

func (g *OperationGenerator) generateOperationWithinBounds(
	initialValue int, operations []Operation, params *ProfileParams,
) (Operation, error) {
	for i := 0; ; i++ {
		operation := g.generateOperation(params)
		temporaryOperations := slices.Clone(operations)
		temporaryOperations = append(temporaryOperations, operation)
		temporaryResult := NewOperationExample(initialValue, temporaryOperations)
		if g.withinBounds(temporaryResult.Answer(), params) {
			return operation, nil
		}

		if i > 1000 {
			return nil, ErrUnableToGenerateOperation.NewWithNoMessage()
		}
	}
}

func (g *OperationGenerator) withinBounds(answer int, params *ProfileParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}

func (g *OperationGenerator) generateOperation(params *ProfileParams) Operation {
	operationTypeIndex := g.random.Intn(len(params.AvailableOperationTypes))
	opType := params.AvailableOperationTypes[operationTypeIndex]

	operand := g.randomOperand(g.profile.AvailableOperands)

	switch opType {
	case PlusOperationType:
		return &PlusOperation{operand}
	case MinusOperationType:
		return &MinusOperation{operand}
	case MultiplyOperationType:
		return &MinusOperation{operand}
	case DivideOperationType:
		return &MinusOperation{operand}
	default:
		panic(fmt.Sprintf("unsupported operation type: %v", opType))
	}
}

func (g *OperationGenerator) randomOperand(availableOperands []int) int {
	index := g.random.Intn(len(availableOperands))
	return availableOperands[index]
}
