package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	e := Example{
		initialValue: 2,
		operations: []operation{
			&plusOperation{valueToAdd: 3},
			&minusOperation{valueToSubtract: 1},
		},
	}
	assert.Equal(t, 4, e.Answer())
}
