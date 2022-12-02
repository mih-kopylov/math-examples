package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	params := exampleParams{
		examplesCount:           9,
		minBoundary:             1,
		maxBoundary:             9,
		operandsCount:           2,
		availableOperationTypes: []operationType{plusOperationType, minusOperationType},
		availableOperands:       []int{1, 2, 3, 4, 5},
	}

	correctAnswersCount := 0
	previousAnswers := map[int]int{}
	for i := 0; i < params.examplesCount; i++ {
		example := generateExample(params, previousAnswers)
		fmt.Println(fmt.Sprintf("%v =", example.exerciseString()))

		var answer int
		_, err := fmt.Scanln(&answer)
		if err != nil {
			panic(err)
		}

		correctAnswer := example.answer()
		previousAnswers[correctAnswer] = previousAnswers[correctAnswer] + 1
		if answer == correctAnswer {
			correctAnswersCount++
			fmt.Println("Правильно!")
		} else {
			fmt.Println(fmt.Sprintf("Неправильно. Правильный ответ %v", correctAnswer))
		}
	}

	fmt.Println("================")
	fmt.Println(fmt.Sprintf("Правильных ответов: %v из %v", correctAnswersCount, params.examplesCount))
}

var (
	errUnableToGenerateOperation = errors.New("unable to generate operation")
	errUnableToGenerateExample   = errors.New("unable to generate example")
	errTooFrequentExampleAnswer  = errors.New("too frequent example answer")
)

func generateExample(params exampleParams, previousAnswers map[int]int) example {
	for i := 0; ; i++ {
		result, err := tryGenerateExample(params, previousAnswers)
		if err == nil {
			return result
		}
		if i > 1000 {
			panic(errUnableToGenerateExample)
		}
	}
}

func tryGenerateExample(params exampleParams, previousAnswers map[int]int) (example, error) {
	result := example{}
	result.initialValue = generateOperand(params.availableOperands)
	for i := 0; i < params.operandsCount-1; i++ {
		op, err := generateOperationWithinBounds(result, params)
		if err != nil {
			return example{}, err
		}
		result.operations = append(result.operations, op)
	}
	if tooFrequentAnswer(result.answer(), previousAnswers) {
		return example{}, errTooFrequentExampleAnswer
	}
	return result, nil
}

func tooFrequentAnswer(answer int, previousAnswers map[int]int) bool {
	if len(previousAnswers) == 0 {
		return false
	}

	thisAnswerCount, found := previousAnswers[answer]
	thisAnswerCount++
	if found && len(previousAnswers) == 1 && thisAnswerCount > 1 {
		return true
	}
	for _, count := range previousAnswers {
		if thisAnswerCount-count > 1 {
			return true
		}
	}
	return false
}

func generateOperationWithinBounds(result example, params exampleParams) (operation, error) {
	for i := 0; ; i++ {
		op := generateOperation(params)
		temporaryOperations := slices.Clone(result.operations)
		temporaryOperations = append(temporaryOperations, op)
		temporaryResult := example{result.initialValue, temporaryOperations}
		if withinBounds(temporaryResult.answer(), params) {
			return op, nil
		}
		if i > 100 {
			return nil, errUnableToGenerateOperation
		}
	}
}

func withinBounds(answer int, params exampleParams) bool {
	return answer >= params.minBoundary && answer <= params.maxBoundary
}

func generateOperation(params exampleParams) operation {
	operationTypeIndex := rand.Intn(len(params.availableOperationTypes))
	opType := params.availableOperationTypes[operationTypeIndex]

	operandIndex := rand.Intn(len(params.availableOperands))
	operand := params.availableOperands[operandIndex]

	switch opType {
	case plusOperationType:
		return &plusOperation{operand}
	case minusOperationType:
		return &minusOperation{operand}
	default:
		panic(fmt.Sprintf("unsupported operation type: %v", opType))
	}
}

func generateOperand(operands []int) int {
	index := rand.Intn(len(operands))
	return operands[index]
}

type exampleParams struct {
	examplesCount           int
	minBoundary             int
	maxBoundary             int
	operandsCount           int
	availableOperands       []int
	availableOperationTypes []operationType
}

type operationType string

const (
	plusOperationType  operationType = "plus"
	minusOperationType operationType = "minus"
)

type operation interface {
	apply(initialValue int) int
	operationString() string
	operand() int
}

type plusOperation struct {
	valueToAdd int
}

func (p *plusOperation) apply(initialValue int) int {
	return initialValue + p.valueToAdd
}

func (p *plusOperation) operationString() string {
	return "+"
}

func (p *plusOperation) operand() int {
	return p.valueToAdd
}

type minusOperation struct {
	valueToSubtract int
}

func (p *minusOperation) apply(initialValue int) int {
	return initialValue - p.valueToSubtract
}

func (p *minusOperation) operationString() string {
	return "-"
}

func (p *minusOperation) operand() int {
	return p.valueToSubtract
}

type example struct {
	initialValue int
	operations   []operation
}

func (e *example) exerciseString() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", e.initialValue))
	for _, o := range e.operations {
		builder.WriteString(fmt.Sprintf(" %v %v", o.operationString(), o.operand()))
	}
	return builder.String()
}

func (e *example) answer() int {
	result := e.initialValue
	for _, o := range e.operations {
		result = o.apply(result)
	}
	return result
}
