package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/joomcode/errorx"
)

var (
	ErrAnswerReaderNamespace = errorx.NewNamespace("answerReader")
	ErrFailedReadAnswer      = errorx.NewType(ErrAnswerReaderNamespace, "FailedReadAnswer")
)

type AnswerReader interface {
	Read(printer Printer, index int, example *Example) (*Answer, error)
}

type ConsoleAnswerReader struct {
}

func (r *ConsoleAnswerReader) Read(printer Printer, index int, example *Example) (*Answer, error) {
	for {
		printer.Print("%v) %v = ", index, example.ExerciseString())

		reader := bufio.NewReader(os.Stdin)
		answerString, err := reader.ReadString('\n')
		if err != nil {
			return nil, ErrFailedReadAnswer.WrapWithNoMessage(err)
		}
		printer.PrintUserInput(answerString)

		answerInt, err := strconv.Atoi(strings.TrimSpace(answerString))
		if err != nil {
			printer.Println("Невозможно прочитать ответ, ответ должен быть числом")
			continue
		}

		answer := Answer{example, answerInt}
		return &answer, nil
	}
}

func NewConsoleAnswerReader() AnswerReader {
	return &ConsoleAnswerReader{}
}
