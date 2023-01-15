package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTooFrequentAnswer(t *testing.T) {
	assert.True(t, tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 3}))
	assert.True(t, tooFrequentAnswer(5, map[int]int{5: 1}))
	assert.False(t, tooFrequentAnswer(5, map[int]int{}))
	assert.False(t, tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 1}))
}
