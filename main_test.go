package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	e := NewOperationExample(
		2, []Operation{
			&PlusOperation{valueToAdd: 3},
			&MinusOperation{valueToSubtract: 1},
		},
	)
	assert.Equal(t, 4, e.Answer())
}
