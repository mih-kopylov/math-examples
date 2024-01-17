package main

import (
	"fmt"
	"strings"
)

type OperationExample struct {
	initialValue int
	operations   []Operation
}

func NewOperationExample(initialValue int, operations []Operation) Example {
	return &OperationExample{
		initialValue: initialValue,
		operations:   operations,
	}
}

func (e *OperationExample) ExerciseString() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", e.initialValue))
	for _, o := range e.operations {
		builder.WriteString(fmt.Sprintf(" %v %v", o.operationString(), o.operand()))
	}
	return builder.String()
}

func (e *OperationExample) Answer() int {
	result := e.initialValue
	for _, o := range e.operations {
		result = o.apply(result)
	}
	return result
}
