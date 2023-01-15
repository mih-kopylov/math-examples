package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const configFileName = "math-examples.yaml"

func main() {
	rand.Seed(time.Now().UnixNano())

	params, err := readParams()
	if err != nil {
		fmt.Println("Не удалось прочитать конфигурацию, похоже, она повреждена.")
		waitForEnter()
		panic(err)
	}

	st := newStat()
	for i := 0; i < params.ExamplesCount; i++ {
		ex, err := generateExample(params, st)
		if err != nil {
			if errors.Is(err, errUnableToGenerateExample) {
				fmt.Println("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию.")
				waitForEnter()
				os.Exit(1)
			}
			panic(err)
		}

		userAnswer := readAnswer(ex)
		ans := st.add(ex, userAnswer)

		if params.ShowCorrectAnswerAfter == afterEach {
			ans.printCorrectAnswer()
		}
	}

	fmt.Println("================")
	if params.ShowCorrectAnswerAfter == afterAll {
		for _, a := range st.answers {
			a.ex.printExercise()
			a.printAnswer()
			a.printCorrectAnswer()
		}
	}
	fmt.Printf("Правильных ответов: %v из %v\n", st.getCorrectAnswersCount(), params.ExamplesCount)
	fmt.Printf("Затраченное время: %v\n", st.getTotalTime().Format("04:05"))
	waitForEnter()
}

func waitForEnter() {
	fmt.Println("Нажмите Enter")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func readAnswer(ex *example) int {
	for {
		ex.printExercise()

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

func generateExample(params *exampleParams, st *stat) (*example, error) {
	for i := 0; ; i++ {
		result, err := tryGenerateExample(params, st)
		if err == nil {
			return result, nil
		}
		if i > 1000 {
			return nil, errUnableToGenerateExample
		}
	}
}

func tryGenerateExample(params *exampleParams, st *stat) (*example, error) {
	result := example{}
	result.initialValue = generateOperand(params.AvailableOperands)
	for i := 0; i < params.OperandsCount-1; i++ {
		op, err := generateOperationWithinBounds(result, params)
		if err != nil {
			return nil, err
		}
		result.operations = append(result.operations, op)
	}
	if st.tooFrequentAnswer(result.answer()) {
		return nil, errTooFrequentExampleAnswer
	}
	return &result, nil
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
