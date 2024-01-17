package main

type Operation interface {
	apply(initialValue int) int
	operationString() string
	operand() int
}

type PlusOperation struct {
	valueToAdd int
}

func (p *PlusOperation) apply(initialValue int) int {
	return initialValue + p.valueToAdd
}

func (p *PlusOperation) operationString() string {
	return "+"
}

func (p *PlusOperation) operand() int {
	return p.valueToAdd
}

type MinusOperation struct {
	valueToSubtract int
}

func (p *MinusOperation) apply(initialValue int) int {
	return initialValue - p.valueToSubtract
}

func (p *MinusOperation) operationString() string {
	return "-"
}

func (p *MinusOperation) operand() int {
	return p.valueToSubtract
}
