package cyclicrotation

func Solution(A []int, k int) []int {
	var temp int
	var N int = len(A)

	for i := 0; i < k; i++ {
		temp = A[N-1]
		copy(A[1:], A[0:])
		A[0] = temp
	}

	return A
}
