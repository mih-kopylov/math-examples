package main

import (
	"fmt"
)

type Operand interface {
	// Value returns the value of the operand
	Value() int

	// String returns string representation of the operand
	String() string

	// NeedsParenthesis indicates if operand needs parenthesis when multiplying
	NeedsParenthesis() bool

	// OperandsCount returns number of operands inside the one
	OperandsCount() int
}

type SingleValueOperand struct {
	value int
}

func NewSingleValueOperand(value int) Operand {
	return &SingleValueOperand{
		value: value,
	}
}

func (o *SingleValueOperand) Value() int {
	return o.value
}

func (o *SingleValueOperand) String() string {
	return fmt.Sprintf("%v", o.value)
}

func (o *SingleValueOperand) NeedsParenthesis() bool {
	return false
}

func (o *SingleValueOperand) OperandsCount() int {
	return 1
}

type ParenthesisOperand struct {
	inner Operand
}

func NewParenthesisOperand(inner Operand) Operand {
	return &ParenthesisOperand{
		inner: inner,
	}
}

func (o *ParenthesisOperand) Value() int {
	return o.inner.Value()
}

func (o *ParenthesisOperand) String() string {
	return fmt.Sprintf("(%v)", o.inner.String())
}

func (o *ParenthesisOperand) NeedsParenthesis() bool {
	return false
}

func (o *ParenthesisOperand) OperandsCount() int {
	return o.inner.OperandsCount()
}

type SumOperand struct {
	left  Operand
	right Operand
}

func NewSumOperand(left Operand, right Operand) Operand {
	return &SumOperand{
		left:  left,
		right: right,
	}
}

func (o *SumOperand) Value() int {
	return o.left.Value() + o.right.Value()
}

func (o *SumOperand) String() string {
	return fmt.Sprintf("%v + %v", o.left, o.right)
}

func (o *SumOperand) NeedsParenthesis() bool {
	return true
}

func (o *SumOperand) OperandsCount() int {
	return o.left.OperandsCount() + o.right.OperandsCount()
}

type SubtractOperand struct {
	left  Operand
	right Operand
}

func NewSubtractOperand(left Operand, right Operand) Operand {
	return &SubtractOperand{
		left:  left,
		right: right,
	}
}

func (o *SubtractOperand) Value() int {
	return o.left.Value() - o.right.Value()
}

func (o *SubtractOperand) String() string {
	return fmt.Sprintf("%v - %v", o.left, o.right)
}

func (o *SubtractOperand) NeedsParenthesis() bool {
	return true
}

func (o *SubtractOperand) OperandsCount() int {
	return o.left.OperandsCount() + o.right.OperandsCount()
}

type MultiplyOperand struct {
	left  Operand
	right Operand
}

func NewMultiplyOperand(left Operand, right Operand) Operand {
	return &MultiplyOperand{
		left:  left,
		right: right,
	}
}

func (o *MultiplyOperand) Value() int {
	return o.left.Value() * o.right.Value()
}

func (o *MultiplyOperand) String() string {
	return fmt.Sprintf("%v * %v", o.left, o.right)
}

func (o *MultiplyOperand) NeedsParenthesis() bool {
	return false
}

func (o *MultiplyOperand) OperandsCount() int {
	return o.left.OperandsCount() + o.right.OperandsCount()
}

type DivideOperand struct {
	left  Operand
	right Operand
}

func NewDivideOperand(left Operand, right Operand) Operand {
	return &DivideOperand{
		left:  left,
		right: right,
	}
}

func (o *DivideOperand) Value() int {
	return o.left.Value() / o.right.Value()
}

func (o *DivideOperand) String() string {
	return fmt.Sprintf("%v / %v", o.left, o.right)
}

func (o *DivideOperand) NeedsParenthesis() bool {
	return false
}

func (o *DivideOperand) OperandsCount() int {
	return o.left.OperandsCount() + o.right.OperandsCount()
}
