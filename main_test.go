package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	e := example{
		initialValue: 2,
		operations: []operation{
			&plusOperation{valueToAdd: 3},
			&minusOperation{valueToSubtract: 1},
		},
	}
	assert.Equal(t, 4, e.answer())
}

func TestTooFrequentAnswer(t *testing.T) {
	assert.True(t, tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 3}))
	assert.True(t, tooFrequentAnswer(5, map[int]int{5: 1}))
	assert.False(t, tooFrequentAnswer(5, map[int]int{}))
	assert.False(t, tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 1}))
}
