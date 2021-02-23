package solver

import "math/big"

type problem001Solver struct{}

func (s *problem001Solver) Solve() (*big.Int, error) {
	var sum int64 = 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += int64(i)
		}
	}
	return big.NewInt(sum), nil
}
