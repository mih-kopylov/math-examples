package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const configFileName = "math-examples.yaml"

func main() {
	startedAt := time.Now()
	rand.Seed(time.Now().UnixNano())

	params, err := readParams()
	if err != nil {
		fmt.Println("Не удалось прочитать конфигурацию, похоже, она повреждена.")
		waitForEnter()
		panic(err)
	}

	correctAnswersCount := 0
	previousAnswers := map[int]int{}
	for i := 0; i < params.ExamplesCount; i++ {
		ex, err := generateExample(params, previousAnswers)
		if err != nil {
			if errors.Is(err, errUnableToGenerateExample) {
				fmt.Println("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию.")
				waitForEnter()
				os.Exit(1)
			}
			panic(err)
		}

		fmt.Printf("%v =\n", ex.exerciseString())

		answer := readAnswer()
		correctAnswer := ex.answer()

		previousAnswers[correctAnswer] = previousAnswers[correctAnswer] + 1
		if answer == correctAnswer {
			correctAnswersCount++
			fmt.Println("Правильно!")
		} else {
			fmt.Printf("Неправильно. Правильный ответ %v\n", correctAnswer)
		}
	}

	fmt.Println("================")
	fmt.Printf("Правильных ответов: %v из %v\n", correctAnswersCount, params.ExamplesCount)
	fmt.Printf("Затраченное время: %v\n", time.Time{}.Add(time.Since(startedAt)).Format("04:05"))
	waitForEnter()
}

func waitForEnter() {
	fmt.Println("Нажмите Enter")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func readParams() (*exampleParams, error) {
	_, err := os.Stat(configFileName)
	if errors.Is(err, os.ErrNotExist) {
		defaultParams := exampleParams{
			ExamplesCount:           10,
			MinBoundary:             0,
			MaxBoundary:             9,
			OperandsCount:           2,
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

func readAnswer() int {
	for {
		reader := bufio.NewReader(os.Stdin)
		answerString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		answerInt, err := strconv.Atoi(strings.TrimSpace(answerString))
		if err == nil {
			return answerInt
		}
		fmt.Println("Невозможно прочитать ответ, ответ должен быть числом")
	}
}

var (
	errUnableToGenerateOperation = errors.New("unable to generate operation")
	errUnableToGenerateExample   = errors.New("unable to generate example")
	errTooFrequentExampleAnswer  = errors.New("too frequent example answer")
)

func generateExample(params *exampleParams, previousAnswers map[int]int) (*example, error) {
	for i := 0; ; i++ {
		result, err := tryGenerateExample(params, previousAnswers)
		if err == nil {
			return result, nil
		}
		if i > 1000 {
			return nil, errUnableToGenerateExample
		}
	}
}

func tryGenerateExample(params *exampleParams, previousAnswers map[int]int) (*example, error) {
	result := example{}
	result.initialValue = generateOperand(params.AvailableOperands)
	for i := 0; i < params.OperandsCount-1; i++ {
		op, err := generateOperationWithinBounds(result, params)
		if err != nil {
			return nil, err
		}
		result.operations = append(result.operations, op)
	}
	if tooFrequentAnswer(result.answer(), previousAnswers) {
		return nil, errTooFrequentExampleAnswer
	}
	return &result, nil
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

func generateOperationWithinBounds(result example, params *exampleParams) (operation, error) {
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

func withinBounds(answer int, params *exampleParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}

func generateOperation(params *exampleParams) operation {
	operationTypeIndex := rand.Intn(len(params.AvailableOperationTypes))
	opType := params.AvailableOperationTypes[operationTypeIndex]

	operandIndex := rand.Intn(len(params.AvailableOperands))
	operand := params.AvailableOperands[operandIndex]

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
	ExamplesCount           int             `yaml:"examplesCount"`
	MinBoundary             int             `yaml:"minBoundary"`
	MaxBoundary             int             `yaml:"maxBoundary"`
	OperandsCount           int             `yaml:"operandsCount"`
	AvailableOperands       []int           `yaml:"availableOperands"`
	AvailableOperationTypes []operationType `yaml:"availableOperationTypes"`
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
