package cyclicrotation

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Save (n - k) to (n - 1) elements
// 2. Move (0) to (n - k - 1) elements to the right. Iterate right to left.
// 3. Move k elements from 1" to the left. Iterate left to right.
// 4. 1" and 2" probably should be reversed when k < n/2

func measureTime(f func([]int, int) []int, A []int, k int) time.Duration {
	var start = time.Now()
	f(A, k)
	return time.Since(start)
}

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

func makeArray() []int {
	var A [1000]int
	for i := 0; i < 1000; i++ {
		A[i] = rand.Intn(2000) - 1000
	}

	return A[:100]
}

func main() {
	var arr []int
	var k int

	for i := 1; i < 1000; i++ {
		k = rand.Intn(1000)
		arr = makeArray()
		fmt.Println(measureTime(Solution, arr, k))
	}
}
