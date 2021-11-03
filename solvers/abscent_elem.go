package solvers

// consider using xor for optimization
func AbscentElem(A []float64) float64 {
	N := len(A)

	expected := float64((N*N + 3*N + 2) / 2)
	var actual float64

	for i := 0; i < N; i++ {
		actual += A[i]
	}

	return expected - actual
}
