package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistribution_IsTooFrequent(t *testing.T) {
	type testCase struct {
		values       []int
		valueToCheck int
		expected     bool
	}
	tests := []testCase{
		{
			values:       nil,
			valueToCheck: 1,
			expected:     false,
		},
		{
			values:       []int{1},
			valueToCheck: 1,
			expected:     true,
		},
		{
			values:       []int{1, 1},
			valueToCheck: 1,
			expected:     true,
		},
		{
			values:       []int{1, 2},
			valueToCheck: 1,
			expected:     false,
		},
		{
			values:       []int{1, 1, 2},
			valueToCheck: 1,
			expected:     true,
		},
		{
			values:       []int{1, 1, 1, 2},
			valueToCheck: 1,
			expected:     true,
		},
		{
			values:       []int{1, 1, 2, 2},
			valueToCheck: 1,
			expected:     false,
		},
		{
			values:       []int{1, 1, 2, 2, 3},
			valueToCheck: 1,
			expected:     true,
		},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%v", tt.values), func(t *testing.T) {
				distribution := NewDistribution[int]()
				for _, value := range tt.values {
					distribution.Add(value)
				}
				actual := distribution.IsTooFrequent(tt.valueToCheck)
				assert.Equal(t, tt.expected, actual)
			},
		)
	}
}
