package main

type OperandExample struct {
	root Operand
}

func NewOperandExample(root Operand) Example {
	return &OperandExample{
		root: root,
	}
}

func (e *OperandExample) ExerciseString() string {
	return e.root.String()
}

func (e *OperandExample) Answer() int {
	return e.root.Value()
}
