package main

import (
	"fmt"
	"time"
)

type stat struct {
	answers []answer
	start   time.Time
}

func newStat() *stat {
	return &stat{
		answers: nil,
		start:   time.Now(),
	}
}

func (s *stat) getTotalTime() time.Time {
	return time.Time{}.Add(time.Since(s.start))
}

func (s *stat) tooFrequentAnswer(value int) bool {
	statMap := make(map[int]int)
	for _, v := range s.answers {
		exAnswer := v.ex.answer()
		statMap[exAnswer] = statMap[exAnswer] + 1
	}
	return tooFrequentAnswer(value, statMap)
}

func tooFrequentAnswer(value int, statMap map[int]int) bool {
	if len(statMap) == 0 {
		return false
	}

	thisAnswerCount, found := statMap[value]
	thisAnswerCount++
	if found && len(statMap) == 1 && thisAnswerCount > 1 {
		return true
	}
	for _, count := range statMap {
		if thisAnswerCount-count > 1 {
			return true
		}
	}
	return false

}

func (s *stat) add(ex *example, userAnswer int) *answer {
	result := answer{
		ex:     ex,
		answer: userAnswer,
	}
	s.answers = append(s.answers, result)
	return &result
}

func (s *stat) getCorrectAnswersCount() int {
	result := 0
	for _, a := range s.answers {
		if a.ex.isCorrectAnswer(a.answer) {
			result++
		}
	}
	return result
}

type answer struct {
	ex     *example
	answer int
}

func (a *answer) printCorrectAnswer() {
	if a.ex.isCorrectAnswer(a.answer) {
		fmt.Println("Правильно!")
	} else {
		fmt.Printf("Неправильно. Правильный ответ %v\n", a.ex.answer())
	}
}

func (a *answer) printAnswer() {
	fmt.Printf("%v ", a.answer)
}
