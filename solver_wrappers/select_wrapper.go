package solver_wrappers

import (
	"context"
	"errors"
	"mservice/config"
)

func Solve(data [][]interface{}, ctx context.Context) error {

	switch ctx.Value("name") {
	case config.CycliclShift:
		err := SolveForCyclicRotation(data)
		if err != nil {
			return errors.New("failed during task solving")
		}
	case config.Warrentries:

	case "":
		return errors.New("no task name were specified")
	}

	return nil
}