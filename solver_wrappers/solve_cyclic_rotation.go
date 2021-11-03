package solver_wrappers

import (
	"errors"
	"mservice/config"
	"mservice/converters"
	"mservice/models"
	"mservice/solvers"
	"reflect"
	"sync"
)

func SolveForCyclicRotation(data [][]interface{}) (models.ResultJson, error) {
	var wg sync.WaitGroup

	var rotated [][]float64 = make([][]float64, len(data))
	for i := range rotated {
		ra := reflect.ValueOf(data[i][0])
		switch ra.Kind() {
		case reflect.Slice:
			rotated[i] = make([]float64, ra.Len())
		default:
			return models.ResultJson{}, errors.New("slice not found in the input data, slice was expected")
		}
	}

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

	solution := models.NewResultWith2DArr(data, rotated, config.CycliclShift)

	return solution, nil
}
