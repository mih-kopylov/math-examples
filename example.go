package main

import (
	"fmt"
	"strings"
)

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

func (e *example) printExercise() {
	fmt.Printf("%v = ", e.exerciseString())
}

func (e *example) answer() int {
	result := e.initialValue
	for _, o := range e.operations {
		result = o.apply(result)
	}
	return result
}

func (e *example) isCorrectAnswer(value int) bool {
	return e.answer() == value
}
