package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeConverter_RangeToInt(t *testing.T) {
	errInvalidRangeMinMax := "range.InvalidRange: Minimum value should be less than maximum: %v"
	errInvalidRange := "range.InvalidRange: Provided: '%v'. Correct value looks like '1:99'"
	tests := []struct {
		sourceRange   string
		expected      []int
		expectedError error
	}{
		{sourceRange: "1", expected: []int{1}, expectedError: nil},
		{sourceRange: "1:3", expected: []int{1, 2, 3}, expectedError: nil},
		{sourceRange: "9:11", expected: []int{9, 10, 11}, expectedError: nil},
		{sourceRange: "11:9", expected: nil, expectedError: fmt.Errorf(errInvalidRangeMinMax, "11:9")},
		{sourceRange: "-5:-3", expected: []int{-5, -4, -3}, expectedError: nil},
		{sourceRange: "-3:-5", expected: nil, expectedError: fmt.Errorf(errInvalidRangeMinMax, "-3:-5")},
		{sourceRange: "1:1", expected: []int{1}, expectedError: nil},
		{sourceRange: "a:b", expected: nil, expectedError: fmt.Errorf(errInvalidRange, "a:b")},
	}
	for _, tt := range tests {
		t.Run(
			tt.sourceRange, func(t *testing.T) {
				converter := NewRangeConverter()
				actual, err := converter.RangeToInt([]string{tt.sourceRange})
				if tt.expectedError != nil {
					if assert.Error(t, err) {
						assert.EqualError(t, err, tt.expectedError.Error())
					}
					return
				} else if !assert.NoError(t, err) {
					return
				}

				assert.Equal(t, tt.expected, actual)
			},
		)
	}
}
