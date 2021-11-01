package solver_wrappers

import (
	"errors"
	"fmt"
	"mservice/converters"
	"mservice/solvers"
	"reflect"
	"sync"
)

func SolveForCyclicRotation(data [][]interface{}) ([][][]float64, []float64, error) {
	var wg sync.WaitGroup

	var resK []float64 = make([]float64, len(data))
	var resA [][][]float64 = make([][][]float64, len(data))
	for i := range resA {
		resA[i] = make([][]float64, 2)
		ra := reflect.ValueOf(data[i][0])
		switch ra.Kind() {
		case reflect.Slice:
			resA[i][0] = make([]float64, ra.Len())
			resA[i][1] = make([]float64, ra.Len())
		default:
			return nil, nil, errors.New("slice not found in the input data, slice was expected")
		}
	}

	wg.Add(len(data))
	for i := range data {
		go func(index int) {
			arr, err := converters.GetArrayAt(data, index)
			if err != nil {
				panic(err)
			}

			k, err := converters.GetFloatAt(data, index)
			if err != nil {
				panic(err)
			}

			arrRot := make([]float64, len(arr))
			copy(arrRot, arr)

			solvers.CyclicRotation(arrRot, k)

			copy(resA[index][0], arr)
			copy(resA[index][1], arrRot)
			resK[index] = k

			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := range resA {
		fmt.Println("Original:")
		fmt.Println(resA[i][0])
		fmt.Println("Rotation size:")
		fmt.Println(resK[i])
		fmt.Println("Rotated:")
		fmt.Println(resA[i][1])
		fmt.Println()
	}

	return resA, resK, nil
}
