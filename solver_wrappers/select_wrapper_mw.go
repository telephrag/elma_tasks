package solver_wrappers

import (
	"context"
	"errors"
	"mservice/config"
	"mservice/models"
	"mservice/solvers"
)

func SelectWrapperMw(ctx context.Context) (models.ResultJson, error) {
	var solution models.ResultJson
	var err error

	taskName, ok := ctx.Value("taskName").(string)
	if !ok {
		return models.ResultJson{}, errors.New("invalid context received")
	}

	data, ok := ctx.Value("inputData").([][]interface{})
	if !ok {
		return models.ResultJson{}, errors.New("invalid context received")
	}

	switch taskName {
	case config.CycliclShift:
		solution, err = SolveForCyclicRotation(data)
	case config.Warrentries:
		solution, err = SolveForOthers(data, solvers.Warrentries, taskName)
	case config.AbscentElem:
		solution, err = SolveForOthers(data, solvers.AbscentElem, taskName)
	case config.SequenceCheck:
		solution, err = SolveForOthers(data, solvers.SequenceCheck, taskName)
	case "":
		return models.ResultJson{}, errors.New("no task name were specified")
	}

	if err != nil {
		return models.ResultJson{}, errors.New("failed during task solving")
	}

	return solution, nil
}
