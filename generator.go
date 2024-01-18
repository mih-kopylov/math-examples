package main

import (
	"github.com/joomcode/errorx"
)

var (
	ErrGeneratorNamespace       = errorx.NewNamespace("generator")
	ErrUnableToGenerateExample  = errorx.NewType(ErrGeneratorNamespace, "UnableToGenerateExample")
	ErrTooFrequentExampleAnswer = errorx.NewType(ErrGeneratorNamespace, "TooFrequentExampleAnswer")
	ErrUnsupportedOperationType = errorx.NewType(ErrGeneratorNamespace, "UnsupportedOperationType")
)

type Generator interface {
	GenerateExample(params *ProfileParams, distribution *Distribution[int]) (Example, error)
}
