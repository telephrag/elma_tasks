package models

import (
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
	Results  []interface{}   `json:"results"`
}

func NewResultWith2DArr(payloads [][]interface{}, results [][]float64, taskName string) ResultJson {
	pToRes := payloadsToResults{}
	pToRes.Payloads = payloads
	pToRes.Results = make([]interface{}, len(results))
	for i := range pToRes.Results {
		pToRes.Results[i] = make([]interface{}, len(results[i]))
		pToRes.Results[i] = results[i]
	}

	rj := ResultJson{
		UserName: config.UserName,
		TaskName: taskName,
		Results:  pToRes,
	}

	return rj
}

func NewResultWith1DArr(payloads [][]interface{}, results []float64, taskName string) ResultJson {

	pToRes := payloadsToResults{}
	pToRes.Payloads = payloads
	pToRes.Results = make([]interface{}, len(results))
	for i := range pToRes.Results {
		pToRes.Results[i] = results[i]
	}

	rj := ResultJson{
		UserName: config.UserName,
		TaskName: taskName,
		Results:  pToRes,
	}

	return rj
}

func (rj ResultJson) Print() {
	fmt.Printf("User name: %s\n", rj.UserName)
	fmt.Printf("Task:      %s\n", rj.TaskName)
	fmt.Printf("Results:\n")
	fmt.Println("	Payload: \n		", rj.Results.Payloads)
	fmt.Println("	Results:  \n		", rj.Results.Results)
}

func (rj ResultJson) Empty() (res bool) {
	res = (rj.UserName == "")
	res = res && (rj.TaskName == "")
	res = res && (len(rj.Results.Payloads) == 0)
	res = res && (len(rj.Results.Results) == 0)
	return res
}
