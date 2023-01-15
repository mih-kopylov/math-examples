package main

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
