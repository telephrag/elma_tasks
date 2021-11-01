package solvers

func CyclicRotation(A []float64, k float64) []float64 {
	var temp float64
	var N int = len(A)

	for i := 0; i < int(k); i++ {
		temp = A[N-1]
		copy(A[1:], A[0:])
		A[0] = temp
	}

	return A
}
