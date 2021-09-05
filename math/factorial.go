package factorial

import "math/big"

func Factorial(n int64) *big.Int {
	return big.NewInt(1).MulRange(2, n)
}
