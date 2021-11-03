package solver_wrappers

import (
	"mservice/converters"
	"mservice/models"
	"sync"
)

func SolveForOthers(data [][]interface{}, solver func([]float64) float64, taskName string) (models.ResultJson, error) {
	var wg sync.WaitGroup

	var res []float64 = make([]float64, len(data))

	wg.Add(len(data))
	for i := range data {
		go func(index int) {
			arr, err := converters.GetF64ArrAt(data, index)
			if err != nil {
				panic(err)
			}

			res[index] = solver(arr)

			wg.Done()
		}(i)
	}
	wg.Wait()

	solution := models.NewResultWith1DArr(data, res, taskName)

	return solution, nil
}
