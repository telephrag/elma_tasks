package solver_wrappers

import (
	"errors"
	"fmt"
	"mservice/converters"
	"mservice/solvers"
	"sync"
)

func SolveForCyclicRotation(data [][]interface{}) error {
	var wg sync.WaitGroup
	var res [][][]float64 = make([][][]float64, len(data))

	for i := range res {
		res[i] = make([][]float64, 2)
		res[i][0] = make([]float64, len(data[i]))
		res[i][1] = make([]float64, len(data[i]))
	}

	for i := range data {
		wg.Add(len(data))
		go func(index int) error {
			arr, err := converters.GetArrayAt(data, index)
			if err != nil {
				return errors.New("failed conversion to slice")
			}

			k, err := converters.GetFloatAt(data, index)
			if err != nil {
				return errors.New("failed conversion to slice")
			}

			fmt.Println(solvers.CyclicRotation(arr, k))
			// TODO: Write to response here using lock

			wg.Done()

			return nil
		}(i)
	}

	return nil
}

func PrintResult(arr [][][]float64) {
	for i := range arr {
		fmt.Println("Index:  ", i)
		fmt.Println("Data:   ", arr[i][0])
		fmt.Println("Result: ", arr[i][1])
		fmt.Println()
	}
}
