package solver_wrappers

import (
	"mservice/converters"
	"mservice/models"
	"sync"
)

func SolveForOthers(task models.Task, solver func([]float64) float64) (models.Result, error) {

	data := task.Data
	var res []float64 = make([]float64, len(data))

	var wg sync.WaitGroup
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

	solution := models.NewResultWith1DArr(task, res)

	return solution, nil
}
