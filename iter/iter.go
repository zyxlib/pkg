package iter

func PermutationInts(xs []int) (perm [][]int) {
	n := len(xs)
	if n > 0 {
		perm = solvePermutation(xs, 0, n-1, perm)
	}
	return perm
}

func solvePermutation(xs []int, lo int, hi int, acc [][]int) [][]int {
	if lo == hi {
		ys := make([]int, len(xs))
		copy(ys, xs)
		return append(acc, ys)
	}
	for i := lo; i <= hi; i++ {
		xs[lo], xs[i] = xs[i], xs[lo]
		acc = solvePermutation(xs, lo+1, hi, acc)
		xs[lo], xs[i] = xs[i], xs[lo]
	}
	return acc
}
