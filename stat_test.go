package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTooFrequentAnswer(t *testing.T) {
	stat := NewStat()
	assert.True(t, stat.tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 3}))
	assert.True(t, stat.tooFrequentAnswer(5, map[int]int{5: 1}))
	assert.False(t, stat.tooFrequentAnswer(5, map[int]int{}))
	assert.False(t, stat.tooFrequentAnswer(5, map[int]int{1: 1, 2: 2, 5: 1}))
}
