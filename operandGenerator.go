package main

import (
	"math/rand"
	"time"

	"github.com/joomcode/errorx"
	"golang.org/x/exp/slices"
)

var (
	ErrOperandOutOfBounds      = errorx.NewType(ErrGeneratorNamespace, "OperandOutOfBounds")
	ErrUnableToGenerateOperand = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateOperand")
)

type OperandGenerator struct {
	random  *rand.Rand
	profile *ProfileParams
}

func NewOperandGenerator(profile *ProfileParams) Generator {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &OperandGenerator{
		random:  random,
		profile: profile,
	}
}

func (g *OperandGenerator) GenerateExample(params *ProfileParams, stat *Stat) (Example, error) {
	for tryNumber := 0; ; tryNumber++ {
		result, err := g.tryGenerateExample(params, stat)
		if err == nil {
			return result, nil
		}

		if tryNumber > 1000 {
			return nil, ErrUnableToGenerateExample.New("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию.")
		}
	}
}

func (g *OperandGenerator) tryGenerateExample(params *ProfileParams, stat *Stat) (Example, error) {
	result := NewSingleValueOperand(g.randomValue(params.AvailableOperands))

	for i := 1; i < params.OperandsCount; i++ {
		tempResult, err := g.generateOperandBasedOn(result, params)
		if err != nil {
			return nil, err
		}

		if !g.withinBounds(tempResult.Value(), params) {
			return nil, ErrOperandOutOfBounds.NewWithNoMessage()
		}

		result = tempResult
	}

	if stat.TooFrequentAnswer(result.Value()) {
		return nil, ErrTooFrequentExampleAnswer.NewWithNoMessage()
	}

	return NewOperandExample(result), nil
}

func (g *OperandGenerator) randomValue(availableOperands []int) int {
	index := g.random.Intn(len(availableOperands))
	return availableOperands[index]
}

func (g *OperandGenerator) generateOperandBasedOn(operand Operand, params *ProfileParams) (Operand, error) {
	operationTypeIndex := g.random.Intn(len(params.AvailableOperationTypes))
	opType := params.AvailableOperationTypes[operationTypeIndex]
	originalOperand := operand

	switch opType {
	case PlusOperationType:
		newOperand := NewSingleValueOperand(g.randomValue(params.AvailableOperands))
		return NewSumOperand(originalOperand, newOperand), nil
	case MinusOperationType:
		newOperand := NewSingleValueOperand(g.randomValue(params.AvailableOperands))
		return NewSubtractOperand(originalOperand, newOperand), nil
	case MultiplyOperationType:
		if !slices.Contains(params.AvailableMultiplicationOperands, originalOperand.Value()) {
			return nil, ErrUnableToGenerateOperand.NewWithNoMessage()
		}
		if originalOperand.NeedsParenthesis() {
			originalOperand = NewParenthesisOperand(originalOperand)
		}
		newOperand := NewSingleValueOperand(g.randomValue(params.AvailableMultiplicationOperands))
		return NewMultiplyOperand(originalOperand, newOperand), nil
	case DivideOperationType:
		if originalOperand.NeedsParenthesis() {
			originalOperand = NewParenthesisOperand(originalOperand)
		}
		var availableDivideOperands []int
		for _, op := range params.AvailableMultiplicationOperands {
			if op == 0 {
				continue
			}
			if originalOperand.Value()%op != 0 {
				continue
			}
			result := originalOperand.Value() / op
			if !slices.Contains(params.AvailableOperands, result) {
				continue
			}
			availableDivideOperands = append(availableDivideOperands, op)
		}

		if len(availableDivideOperands) == 0 {
			return nil, ErrUnableToGenerateOperand.NewWithNoMessage()
		}

		newOperand := NewSingleValueOperand(g.randomValue(availableDivideOperands))
		return NewDivideOperand(originalOperand, newOperand), nil
	default:
		return nil, ErrUnsupportedOperationType.New("type: %v", opType)
	}

}

func (g *OperandGenerator) withinBounds(answer int, params *ProfileParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}
