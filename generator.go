package main

import (
	"fmt"

	"github.com/joomcode/errorx"
	"golang.org/x/exp/slices"
)

var (
	ErrGeneratorNamespace        = errorx.NewNamespace("generator")
	ErrUnableToGenerateOperation = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateOperation")
	ErrUnableToGenerateExample   = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateExample")
	ErrTooFrequentExampleAnswer  = errorx.NewType(ErrGeneratorNamespace, "TooFrequentExampleAnswer")
)

type Generator struct {
	profile *ProfileParams
}

func NewGenerator(profile *ProfileParams) *Generator {
	return &Generator{
		profile: profile,
	}
}

func (g *Generator) GenerateExample(params *ProfileParams, st *Stat, index int) (*Example, error) {
	for i := 0; ; i++ {
		result, err := g.tryGenerateExample(params, st)
		if err == nil {
			result.index = index
			return result, nil
		}
		if i > 1000 {
			return nil, ErrUnableToGenerateExample.New("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию.")
		}
	}
}

func (g *Generator) tryGenerateExample(params *ProfileParams, st *Stat) (*Example, error) {
	result := Example{}
	result.initialValue = g.generateOperand(params.AvailableOperands)
	for i := 0; i < params.OperandsCount-1; i++ {
		op, err := g.generateOperationWithinBounds(result, params)
		if err != nil {
			return nil, err
		}
		result.operations = append(result.operations, op)
	}
	if st.TooFrequentAnswer(result.Answer()) {
		return nil, ErrTooFrequentExampleAnswer.NewWithNoMessage()
	}
	return &result, nil
}

func (g *Generator) generateOperationWithinBounds(result Example, params *ProfileParams) (operation, error) {
	for i := 0; ; i++ {
		op := g.generateOperation(params)
		temporaryOperations := slices.Clone(result.operations)
		temporaryOperations = append(temporaryOperations, op)
		temporaryResult := Example{
			initialValue: result.initialValue,
			operations:   temporaryOperations,
		}
		if g.withinBounds(temporaryResult.Answer(), params) {
			return op, nil
		}
		if i > 100 {
			return nil, ErrUnableToGenerateOperation.NewWithNoMessage()
		}
	}
}

func (g *Generator) withinBounds(answer int, params *ProfileParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}

func (g *Generator) generateOperation(params *ProfileParams) operation {
	operationTypeIndex := r.Intn(len(params.AvailableOperationTypes))
	opType := params.AvailableOperationTypes[operationTypeIndex]

	operandIndex := r.Intn(len(params.AvailableOperands))
	operand := params.AvailableOperands[operandIndex]

	switch opType {
	case PlusOperationType:
		return &plusOperation{operand}
	case MinusOperationType:
		return &minusOperation{operand}
	default:
		panic(fmt.Sprintf("unsupported operation type: %v", opType))
	}
}

func (g *Generator) generateOperand(operands []int) int {
	index := r.Intn(len(operands))
	return operands[index]
}
