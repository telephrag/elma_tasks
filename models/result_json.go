package models

import (
	"context"
	"errors"
	"fmt"
	"mservice/config"
)

type ResultJson struct {
	UserName string            `json:"user_name"`
	TaskName string            `json:"task"`
	Results  payloadsToResults `json:"results"`
}

type payloadsToResults struct {
	Payloads [][]interface{} `json:"payload"`
	Results  [][]float64     `json:"results"`
}

func NewResult(payloads [][]interface{}, results [][]float64, taskName string) (ResultJson, error) {
	rj := ResultJson{}

	switch taskName {
	case config.CycliclShift:
		rj = ResultJson{
			UserName: config.UserName,
			TaskName: taskName,
			Results: payloadsToResults{
				Payloads: payloads,
				Results:  results,
			},
		}
	}

	return rj, nil
}

func NewResultFromRaw(rr ResultRaw, data [][]interface{}, ctx context.Context) (ResultJson, error) {
	var rj ResultJson
	switch ctx.Value("name") {
	case config.CycliclShift:
		rj = ResultJson{
			UserName: config.UserName,
			TaskName: ctx.Value("name").(string),
			Results: payloadsToResults{
				Payloads: data,
				Results:  rr.ResultArrs,
			},
		}
	default:
		return ResultJson{}, errors.New("no such task name")
	}

	return rj, nil
}

func (rj ResultJson) Print() {
	fmt.Printf("User name: %s\n", rj.UserName)
	fmt.Printf("Task:      %s\n", rj.TaskName)
	fmt.Printf("Results:\n")
	fmt.Println("	Payload: \n		", rj.Results.Payloads)
	fmt.Println("	Results:  \n		", rj.Results.Results)
}

func (rj ResultJson) Empty() (res bool) {
	res = res && (rj.UserName == "")
	res = res && (rj.TaskName == "")
	res = res && (len(rj.Results.Payloads) == 0)
	res = res && (len(rj.Results.Results) == 0)
	return res
}
