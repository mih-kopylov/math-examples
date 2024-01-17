package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperandGenerator_GenerateExample(t *testing.T) {
	tests := []struct {
		name        string
		profile     *ProfileParams
		stat        *Stat
		expectedErr error
	}{
		{
			name: "plus", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				AvailableMultiplicationOperands: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{PlusOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "minus", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				AvailableMultiplicationOperands: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{MinusOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "multiply", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     100,
				OperandsCount:                   2,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				AvailableMultiplicationOperands: []int{2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{MultiplyOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "divide", profile: &ProfileParams{
				ExamplesCount:          10,
				MinBoundary:            2,
				MaxBoundary:            100,
				OperandsCount:          2,
				ShowCorrectAnswerAfter: "each",
				AvailableOperands: []int{
					2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
					30,
				},
				AvailableMultiplicationOperands: []int{2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{DivideOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "parenthesis for plus and multiply", profile: &ProfileParams{
				ExamplesCount:                   10,
				MinBoundary:                     1,
				MaxBoundary:                     200,
				OperandsCount:                   3,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
				AvailableMultiplicationOperands: []int{2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{PlusOperationType, MultiplyOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "parenthesis for minus and divide", profile: &ProfileParams{
				ExamplesCount:          10,
				MinBoundary:            1,
				MaxBoundary:            200,
				OperandsCount:          3,
				ShowCorrectAnswerAfter: "each",
				AvailableOperands: []int{
					2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
					30,
				},
				AvailableMultiplicationOperands: []int{2, 3, 4, 5, 6, 7, 8, 9},
				AvailableOperationTypes:         []OperationType{MinusOperationType, DivideOperationType},
			}, stat: NewStat(), expectedErr: nil,
		},
		{
			name: "out of bounds", profile: &ProfileParams{
				ExamplesCount:                   1,
				MinBoundary:                     1,
				MaxBoundary:                     10,
				OperandsCount:                   2,
				ShowCorrectAnswerAfter:          "each",
				AvailableOperands:               []int{9},
				AvailableMultiplicationOperands: []int{9},
				AvailableOperationTypes:         []OperationType{PlusOperationType},
			}, stat: NewStat(),
			expectedErr: ErrUnableToGenerateExample.New("Не удалось придумать пример с заданной конфигурацией. Проверьте конфигурацию."),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				generator := NewOperandGenerator(tt.profile)
				for i := 0; i < tt.profile.ExamplesCount; i++ {
					actual, err := generator.GenerateExample(tt.profile, tt.stat)

					if tt.expectedErr != nil {
						if assert.Error(t, err) {
							assert.EqualError(t, err, tt.expectedErr.Error())
						}
						return
					} else if !assert.NoError(t, err) {
						return
					}

					t.Logf("%v) %v = %v", i, actual.ExerciseString(), actual.Answer())
					tt.stat.AddAnswer(&Answer{actual, actual.Answer()})
					assert.GreaterOrEqual(t, actual.Answer(), tt.profile.MinBoundary)
					assert.LessOrEqual(t, actual.Answer(), tt.profile.MaxBoundary)
				}
			},
		)
	}
}

func TestOperandGenerator_GenerateExample_MultiplyUsesOwnOperands(t *testing.T) {

}
