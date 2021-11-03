package solvers

func Warrentries(A []float64) float64 {
	xorSum := uint64(A[0])
	for i := 1; i < len(A); i++ {
		xorSum = xorSum ^ uint64(A[i])
	}

	return float64(xorSum)
}
