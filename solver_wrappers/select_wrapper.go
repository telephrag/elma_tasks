package solver_wrappers

import (
	"context"
	"errors"
	"mservice/config"
	"mservice/models"
	"mservice/solvers"
	"sync"
)

func SelectWrapper(tasks []models.Task) ([]models.Result, error) {

	var solution []models.Result = make([]models.Result, len(tasks))
	var err error

	var wg sync.WaitGroup
	var mu sync.Locker

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	wg.Add(len(tasks))
	for i := range tasks {
		go func(index int, cancelCtx context.Context) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			taskName := tasks[index].TaskName

			switch taskName {
			case config.CycliclShift:
				solution[index], err = SolveForCyclicRotation(tasks[index])
			case config.Warrentries:
				solution[index], err = SolveForOthers(tasks[index], solvers.Warrentries)
			case config.AbscentElem:
				solution[index], err = SolveForOthers(tasks[index], solvers.AbscentElem)
			case config.SequenceCheck:
				solution[index], err = SolveForOthers(tasks[index], solvers.SequenceCheck)
			case "":
				mu.Lock()
				err = errors.New(config.NoTask)
				mu.Unlock()
				cancelFunc()
			}

			if err != nil {
				mu.Lock()
				err = errors.New(config.NoTask)
				mu.Unlock()
				cancelFunc()
			}
		}(i, ctx)
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return solution, nil
}
