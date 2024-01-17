package main

import (
	"time"
)

type Stat struct {
	answers []Answer
	start   time.Time
}

func NewStat() *Stat {
	return &Stat{
		answers: nil,
		start:   time.Now(),
	}
}

func (s *Stat) getTotalTime() time.Time {
	return time.Time{}.Add(time.Since(s.start))
}

func (s *Stat) TooFrequentAnswer(value int) bool {
	statMap := make(map[int]int)
	for _, v := range s.answers {
		exAnswer := v.example.Answer()
		statMap[exAnswer] = statMap[exAnswer] + 1
	}
	return s.tooFrequentAnswer(value, statMap)
}

func (s *Stat) tooFrequentAnswer(value int, statMap map[int]int) bool {
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

func (s *Stat) AddAnswer(answer *Answer) {
	s.answers = append(s.answers, *answer)
}

func (s *Stat) getCorrectAnswersCount() int {
	result := 0
	for _, a := range s.answers {
		if a.example.Answer() == a.value {
			result++
		}
	}
	return result
}

func (s *Stat) PrintStartMessage(printer Printer) {
	printer.Println("Начали решать %v", s.start.Format(time.DateTime))
}

func (s *Stat) PrintAllAnswers(printer Printer) {
	for i, a := range s.answers {
		printer.Println("%v) %v = %v %v", i+1, a.example.ExerciseString(), a.value, a.GetAnalysis())
	}
}

func (s *Stat) PrintStatistics(printer Printer) {
	printer.Println("Правильных ответов: %v из %v", s.getCorrectAnswersCount(), len(s.answers))
	printer.Println("Затраченное время: %v", s.getTotalTime().Format("04:05"))
}
