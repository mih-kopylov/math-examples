package main

import (
	"fmt"
	"strings"
)

type Example struct {
	initialValue int
	operations   []operation
	index        int
}

func (e *Example) ExerciseString() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", e.initialValue))
	for _, o := range e.operations {
		builder.WriteString(fmt.Sprintf(" %v %v", o.operationString(), o.operand()))
	}
	return builder.String()
}

func (e *Example) Answer() int {
	result := e.initialValue
	for _, o := range e.operations {
		result = o.apply(result)
	}
	return result
}

func (e *Example) isCorrectAnswer(value int) bool {
	return e.Answer() == value
}
