package solver_wrappers

import (
	"errors"
	"mservice/converters"
	"mservice/models"
	"mservice/solvers"
	"reflect"
	"sync"
)

func SolveForCyclicRotation(task models.Task) (models.Result, error) {

	var data [][]interface{} = task.Data

	var rotated [][]float64 = make([][]float64, len(data))
	for i := range rotated {
		ra := reflect.ValueOf(data[i][0])
		switch ra.Kind() {
		case reflect.Slice:
			rotated[i] = make([]float64, ra.Len())
		default:
			return models.Result{}, errors.New("slice not found in the input data, slice was expected")
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(data))
	for i := range data {
		go func(index int) {
			arr, err := converters.GetF64ArrAt(data, index)
			if err != nil {
				panic(err)
			}

			k, err := converters.GetF64At(data, index)
			if err != nil {
				panic(err)
			}

			solvers.CyclicRotation(arr, k)

			copy(rotated[index], arr)

			wg.Done()
		}(i)
	}
	wg.Wait()

	solution := models.NewResultWith2DArr(task, rotated)

	return solution, nil
}
