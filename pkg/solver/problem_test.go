package solver

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSolver(t *testing.T) {
	tests := map[int]*big.Int{
		1: big.NewInt(233168),
	}

	for problemNumber, expectAnswer := range tests {
		name := fmt.Sprintf("problem_%03d", problemNumber)
		t.Run(name, func(t *testing.T) {
			solver, err := NewSolver(problemNumber)
			if err != nil {
				t.Error(err)
				return
			}
			actualAnswer, err := solver.Solve()
			if err != nil {
				t.Error(err)
				return
			}
			if actualAnswer.Cmp(expectAnswer) != 0 {
				t.Errorf("Expected is %v but actual is %v", expectAnswer, actualAnswer)
				return
			}
		})
	}
}
