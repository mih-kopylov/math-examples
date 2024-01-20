package main

import (
	"math/rand"
	"time"

	"github.com/joomcode/errorx"
	"golang.org/x/exp/slices"
)

var (
	ErrOperandOutOfBounds            = errorx.NewType(ErrGeneratorNamespace, "OperandOutOfBounds")
	ErrUnableToGenerateOperand       = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateOperand")
	ErrUnableToGenerateOperationType = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateOperationType")
	ErrUnableToGenerateDirection     = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateDirection")
)

type Direction string

const (
	LeftDirection  Direction = "left"
	RightDirection Direction = "right"
)

var AvailableDirections = []Direction{LeftDirection, RightDirection}

type OperandGenerator struct {
	random                          *rand.Rand
	profile                         *ProfileParams
	availableOperands               []int
	availableMultiplicationOperands []int
}

func NewOperandGenerator(profile *ProfileParams) (Generator, error) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	rangeConverter := NewRangeConverter()
	availableOperands, err := rangeConverter.RangeToInt(profile.AvailableOperands)
	if err != nil {
		return nil, err
	}

	availableMultiplicationOperands, err := rangeConverter.RangeToInt(profile.AvailableMultiplicationOperands)
	if err != nil {
		return nil, err
	}

	return &OperandGenerator{
		random:                          random,
		profile:                         profile,
		availableOperands:               availableOperands,
		availableMultiplicationOperands: availableMultiplicationOperands,
	}, nil
}

func (g *OperandGenerator) GenerateExample(params *ProfileParams, distribution *Distribution[int]) (Example, error) {
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

func (g *OperandGenerator) tryGenerateExample(params *ProfileParams, distribution *Distribution[int]) (Example, error) {
	result := NewSingleValueOperand(g.randomValue(g.availableOperands))

	operationTypeDistribution := NewDistributionWithKnownKeys[OperationType](params.AvailableOperationTypes)
	directionDistribution := NewDistributionWithKnownKeys[Direction](AvailableDirections)

	for i := 1; i < params.OperandsCount; i++ {
		tempResult, err := g.generateOperandBasedOn(result, params, operationTypeDistribution, directionDistribution)
		if err != nil {
			return nil, err
		}

		if !g.withinBounds(tempResult.Value(), params) {
			return nil, ErrOperandOutOfBounds.NewWithNoMessage()
		}

		result = tempResult
	}

	if distribution.IsTooFrequent(result.Value()) {
		return nil, ErrTooFrequentExampleAnswer.NewWithNoMessage()
	}

	return NewOperandExample(result), nil
}

func (g *OperandGenerator) randomValue(availableOperands []int) int {
	index := g.random.Intn(len(availableOperands))
	return availableOperands[index]
}

func (g *OperandGenerator) randomDirection() Direction {
	index := g.random.Intn(len(AvailableDirections))
	return AvailableDirections[index]
}

func (g *OperandGenerator) generateDirection(directionDistribution *Distribution[Direction]) (Direction, error) {
	for i := 0; i < 100; i++ {
		direction := g.randomDirection()
		if directionDistribution.IsTooFrequent(direction) {
			continue
		}
		directionDistribution.Add(direction)
		return direction, nil
	}
	return LeftDirection, ErrUnableToGenerateDirection.NewWithNoMessage()
}

func (g *OperandGenerator) generateOperandBasedOn(
	operand Operand, params *ProfileParams, operationTypeDistribution *Distribution[OperationType],
	directionDistribution *Distribution[Direction],
) (Operand, error) {
	originalOperand := operand
	operationType, err := g.generateOperationType(params, operationTypeDistribution)
	if err != nil {
		return nil, err
	}

	direction, err := g.generateDirection(directionDistribution)
	if err != nil {
		return nil, err
	}

	switch operationType {
	case PlusOperationType:
		return g.generateSumOperand(direction, originalOperand)
	case MinusOperationType:
		return g.generateSubtractOperand(direction, originalOperand)
	case MultiplyOperationType:
		return g.generateMultiplyOperand(originalOperand, direction)
	case DivideOperationType:
		return g.generateDivideOperand(originalOperand)
	default:
		return nil, ErrUnsupportedOperationType.New("type: %v", operationType)
	}
}

func (g *OperandGenerator) generateDivideOperand(originalOperand Operand) (Operand, error) {
	if originalOperand.NeedsParenthesis() {
		originalOperand = NewParenthesisOperand(originalOperand)
	}
	var availableDivideOperands []int
	for _, op := range g.availableMultiplicationOperands {
		if op == 0 {
			continue
		}
		if originalOperand.Value()%op != 0 {
			continue
		}
		result := originalOperand.Value() / op
		if !slices.Contains(g.availableOperands, result) {
			continue
		}
		availableDivideOperands = append(availableDivideOperands, op)
	}

	if len(availableDivideOperands) == 0 {
		return nil, ErrUnableToGenerateOperand.NewWithNoMessage()
	}

	newOperand := NewSingleValueOperand(g.randomValue(availableDivideOperands))
	return NewDivideOperand(originalOperand, newOperand), nil
}

func (g *OperandGenerator) generateMultiplyOperand(originalOperand Operand, direction Direction) (Operand, error) {
	if !slices.Contains(g.availableMultiplicationOperands, originalOperand.Value()) {
		return nil, ErrUnableToGenerateOperand.NewWithNoMessage()
	}
	if originalOperand.NeedsParenthesis() {
		originalOperand = NewParenthesisOperand(originalOperand)
	}
	newOperand := NewSingleValueOperand(g.randomValue(g.availableMultiplicationOperands))
	if direction == RightDirection {
		return NewMultiplyOperand(originalOperand, newOperand), nil
	}
	return NewMultiplyOperand(newOperand, originalOperand), nil
}

func (g *OperandGenerator) generateSubtractOperand(direction Direction, originalOperand Operand) (Operand, error) {
	newOperand := NewSingleValueOperand(g.randomValue(g.availableOperands))
	if direction == RightDirection {
		return NewSubtractOperand(originalOperand, newOperand), nil
	}
	if originalOperand.NeedsParenthesis() {
		originalOperand = NewParenthesisOperand(originalOperand)
	}
	return NewSubtractOperand(newOperand, originalOperand), nil
}

func (g *OperandGenerator) generateSumOperand(direction Direction, originalOperand Operand) (Operand, error) {
	newOperand := NewSingleValueOperand(g.randomValue(g.availableOperands))
	if direction == RightDirection {
		return NewSumOperand(originalOperand, newOperand), nil
	}
	return NewSumOperand(newOperand, originalOperand), nil
}

func (g *OperandGenerator) generateOperationType(
	params *ProfileParams, operationTypeDistribution *Distribution[OperationType],
) (OperationType, error) {
	if len(params.AvailableOperationTypes) > 1 {
		for tryNumber := 0; tryNumber < 100; tryNumber++ {
			operationType := g.getRandomOperationType(params)
			if operationTypeDistribution.IsTooFrequent(operationType) {
				continue
			}
			operationTypeDistribution.Add(operationType)
			return operationType, nil
		}
		return PlusOperationType, ErrUnableToGenerateOperationType.NewWithNoMessage()
	}
	return g.getRandomOperationType(params), nil

}

func (g *OperandGenerator) getRandomOperationType(params *ProfileParams) OperationType {
	operationTypeIndex := g.random.Intn(len(params.AvailableOperationTypes))
	return params.AvailableOperationTypes[operationTypeIndex]
}

func (g *OperandGenerator) withinBounds(answer int, params *ProfileParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}
