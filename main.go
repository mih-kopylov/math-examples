package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	app, err := readParams()
	if err != nil {
		fmt.Println("Не удалось прочитать конфигурацию, похоже, она повреждена.")
		waitForEnter()
		panic(err)
	}

	profileName := flag.String("p", "", "Profile to use")
	flag.Parse()

	if *profileName == "" {
		fmt.Println("Не передано имя профиля. Используйте -p аргумент")
		waitForEnter()
		os.Exit(1)
	}

	profile, exists := app.Profiles[*profileName]
	if !exists {
		fmt.Println(
			fmt.Sprintf(
				"Не удалось найти профиль %s. Известные профили: %s", *profileName, maps.Keys(app.Profiles),
			),
		)
		waitForEnter()
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Добро пожаловать, %s!", *profileName))

	st := newStat()
	fmt.Printf("Начали решать %v\n", time.Now().Format(time.DateTime))
	for i := 0; i < profile.ExamplesCount; i++ {
		ex, err := generateExample(&profile, st, i+1)
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

		if profile.ShowCorrectAnswerAfter == afterEach {
			ans.printCorrectAnswer()
		}
	}

	fmt.Println("================")
	if profile.ShowCorrectAnswerAfter == afterAll {
		for _, a := range st.answers {
			a.ex.printExercise()
			a.printAnswer()
			a.printCorrectAnswer()
		}
	}
	fmt.Printf("Правильных ответов: %v из %v\n", st.getCorrectAnswersCount(), profile.ExamplesCount)
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

func generateExample(params *profileParams, st *stat, number int) (*example, error) {
	for i := 0; ; i++ {
		result, err := tryGenerateExample(params, st)
		if err == nil {
			result.number = number
			return result, nil
		}
		if i > 1000 {
			return nil, errUnableToGenerateExample
		}
	}
}

func tryGenerateExample(params *profileParams, st *stat) (*example, error) {
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

func generateOperationWithinBounds(result example, params *profileParams) (operation, error) {
	for i := 0; ; i++ {
		op := generateOperation(params)
		temporaryOperations := slices.Clone(result.operations)
		temporaryOperations = append(temporaryOperations, op)
		temporaryResult := example{
			initialValue: result.initialValue,
			operations:   temporaryOperations,
		}
		if withinBounds(temporaryResult.answer(), params) {
			return op, nil
		}
		if i > 100 {
			return nil, errUnableToGenerateOperation
		}
	}
}

func withinBounds(answer int, params *profileParams) bool {
	return answer >= params.MinBoundary && answer <= params.MaxBoundary
}

func generateOperation(params *profileParams) operation {
	operationTypeIndex := r.Intn(len(params.AvailableOperationTypes))
	opType := params.AvailableOperationTypes[operationTypeIndex]

	operandIndex := r.Intn(len(params.AvailableOperands))
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
	index := r.Intn(len(operands))
	return operands[index]
}
