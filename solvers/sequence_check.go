package solvers

import "sort"

func SequenceCheck(A []float64) float64 {
	sort.Float64s(A)

	for i := 0; i < len(A)-1; i++ {
		if A[i] != A[i+1]+1 {
			return 0
		}
	}

	return 1
}
