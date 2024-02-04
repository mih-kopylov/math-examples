package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperandGenerator_GenerateExample(t *testing.T) {
	tests := []struct {
		name         string
		profile      *ProfileParams
		distribution *Distribution[int]
		expectedErr  error
	}{
		{
			name: "plus", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:9"},
				AvailableMultiplicationOperands: []string{"1:9"},
				AvailableOperationTypes:         []OperationType{PlusOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "minus", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:9"},
				AvailableMultiplicationOperands: []string{"1:9"},
				AvailableOperationTypes:         []OperationType{MinusOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "plus and minus with parenthesis", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   5,
				Parenthesis:                     true,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:50"},
				AvailableMultiplicationOperands: []string{"1:9"},
				AvailableOperationTypes:         []OperationType{PlusOperationType, MinusOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "plus and minus without parenthesis", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   5,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:50"},
				AvailableMultiplicationOperands: []string{"1:9"},
				AvailableOperationTypes:         []OperationType{PlusOperationType, MinusOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "multiply", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:11"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes:         []OperationType{MultiplyOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "divide", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     2,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"2:30"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes:         []OperationType{DivideOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "parenthesis for plus and multiply", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     200,
				OperandsCount:                   3,
				Parenthesis:                     true,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:15"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes:         []OperationType{PlusOperationType, MultiplyOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "parenthesis for plus, minus and multiply", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   4,
				Parenthesis:                     true,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"1:50"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes: []OperationType{
					PlusOperationType, MinusOperationType, MultiplyOperationType,
				},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "parenthesis for minus and divide", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     200,
				OperandsCount:                   3,
				Parenthesis:                     true,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"2:30"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes:         []OperationType{MinusOperationType, DivideOperationType},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "all operations with parenthesis", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     200,
				OperandsCount:                   5,
				Parenthesis:                     true,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"2:30"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes: []OperationType{
					PlusOperationType, MinusOperationType, MultiplyOperationType, DivideOperationType,
				},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "all operations without parenthesis", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     200,
				OperandsCount:                   5,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"2:30"},
				AvailableMultiplicationOperands: []string{"2:9"},
				AvailableOperationTypes: []OperationType{
					PlusOperationType, MinusOperationType, MultiplyOperationType, DivideOperationType,
				},
			}, distribution: NewDistribution[int](), expectedErr: nil,
		},
		{
			name: "out of bounds", profile: &ProfileParams{
				ExamplesCount:                   1,
				MinBoundary:                     1,
				MaxBoundary:                     10,
				OperandsCount:                   2,
				Parenthesis:                     false,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []string{"9"},
				AvailableMultiplicationOperands: []string{"9"},
				AvailableOperationTypes:         []OperationType{PlusOperationType},
			}, distribution: NewDistribution[int](),
			expectedErr: ErrUnableToGenerateExample.New("The configuration looks wrong"),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				generator, err := NewOperandGenerator(tt.profile)
				if !assert.NoError(t, err) {
					return
				}

				for i := 0; i < tt.profile.ExamplesCount; i++ {
					actual, err := generator.GenerateExample(tt.profile, tt.distribution)

					if tt.expectedErr != nil {
						if assert.Error(t, err) {
							assert.EqualError(t, err, tt.expectedErr.Error())
						}
						return
					} else if !assert.NoError(t, err) {
						return
					}

					t.Logf("%v) %v = %v", i, actual.ExerciseString(), actual.Answer())
					tt.distribution.Add(actual.Answer())
					assert.GreaterOrEqual(t, actual.Answer(), tt.profile.MinBoundary)
					assert.LessOrEqual(t, actual.Answer(), tt.profile.MaxBoundary)
				}
			},
		)
	}
}
