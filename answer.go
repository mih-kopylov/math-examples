package main

import (
	"fmt"
)

type Answer struct {
	example Example
	value   int
}

func (a *Answer) GetAnalysis() string {
	if a.example.Answer() == a.value {
		return "Правильно!"
	}
	return fmt.Sprintf("Неправильно. Правильный ответ %v", a.example.Answer())
}
