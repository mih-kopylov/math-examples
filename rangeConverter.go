package main

import (
	"regexp"
	"strconv"

	"github.com/joomcode/errorx"
)

const (
	numberRegexpString = "^\\d+$"
	rangeRegexpString  = "^(-?\\d+):(-?\\d+)$"
)

var (
	ErrRangeNamespace = errorx.NewNamespace("range")
	ErrInvalidRegexp  = errorx.NewType(ErrRangeNamespace, "InvalidRegexp")
	ErrInvalidRange   = errorx.NewType(ErrRangeNamespace, "InvalidRange")
)

type RangeConverter struct {
}

func NewRangeConverter() *RangeConverter {
	return &RangeConverter{}
}

func (p *RangeConverter) RangeToInt(stringRanges []string) ([]int, error) {
	var result []int

	numberRegexp, err := regexp.Compile(numberRegexpString)
	if err != nil {
		return nil, ErrInvalidRegexp.Wrap(err, numberRegexpString)
	}

	rangeRegexp, err := regexp.Compile(rangeRegexpString)
	if err != nil {
		return nil, ErrInvalidRegexp.Wrap(err, rangeRegexpString)
	}

	for _, stringValue := range stringRanges {
		if numberRegexp.MatchString(stringValue) {
			intValue, err := strconv.Atoi(stringValue)
			if err != nil {
				return nil, ErrInvalidRange.Wrap(err, stringValue)
			}
			result = append(result, intValue)
		} else if rangeRegexp.MatchString(stringValue) {
			submatch := rangeRegexp.FindStringSubmatch(stringValue)
			minValue, err := strconv.Atoi(submatch[1])
			if err != nil {
				return nil, ErrInvalidRange.Wrap(err, "%v %v", stringValue, submatch[1])
			}

			maxValue, err := strconv.Atoi(submatch[2])
			if err != nil {
				return nil, ErrInvalidRange.Wrap(err, "%v %v", stringValue, submatch[2])
			}

			if minValue > maxValue {
				return nil, ErrInvalidRange.New(
					"Минимальное значение должно быть меньше максимального: %v", stringValue,
				)
			}

			for i := minValue; i <= maxValue; i++ {
				result = append(result, i)
			}
		} else {
			return nil, ErrInvalidRange.New("Некорректный диапазон: %v", stringValue)
		}
	}

	return result, nil
}
