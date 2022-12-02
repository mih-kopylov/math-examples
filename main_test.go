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
