package solver_wrappers

import (
	"errors"
	"mservice/converters"
	"mservice/models"
	"mservice/solvers"
	"reflect"
	"sync"
)

func SolveForCyclicRotation(data [][]interface{}) (models.ResultRaw, error) {
	var wg sync.WaitGroup

	var rotated [][][]float64 = make([][][]float64, len(data))
	for i := range rotated {
		ra := reflect.ValueOf(data[i][0])
		switch ra.Kind() {
		case reflect.Slice:
			rotated[i] = make([][]float64, 1)
			rotated[i][0] = make([]float64, ra.Len())
		default:
			return models.ResultRaw{}, errors.New("slice not found in the input data, slice was expected")
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

			solvers.CyclicRotation(arr, k)

			copy(rotated[index][0], arr)

			wg.Done()
		}(i)
	}
	wg.Wait()

	solution := models.ResultRaw{
		ResultArrs: rotated,
	}

	return solution, nil
}
