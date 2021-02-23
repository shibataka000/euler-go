package solver

import (
	"fmt"
	"math/big"
)

// Solver solve specific problem in Project Euler
type Solver interface {
	Solve() (*big.Int, error)
}

// NewSolver is factory of Solver interface
func NewSolver(problemNumber int) (Solver, error) {
	switch problemNumber {
	case 1:
		return &problem001Solver{}, nil
	default:
		return nil, fmt.Errorf("No such solver")
	}
}
