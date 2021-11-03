package solver_wrappers

import (
	"context"
	"errors"
	"mservice/config"
	"mservice/models"
)

func SelectWrapper(data [][]interface{}, ctx context.Context) (models.ResultJson, error) {

	var solution models.ResultJson
	var err error

	switch ctx.Value("name") {
	case config.CycliclShift:
		solution, err = SolveForCyclicRotation(data)
		if err != nil {
			return models.ResultJson{}, errors.New("failed during task solving")
		}
	case config.Warrentries:

	case "":
		return models.ResultJson{}, errors.New("no task name were specified")
	}

	return solution, nil
}
