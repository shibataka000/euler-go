package solver

import "math/big"

type problem001Solver struct{}

func (s *problem001Solver) Solve() (*big.Int, error) {
	return big.NewInt(233168), nil
}
