package models

type Task struct {
	TaskName string          `json:"taskName"`
	Data     [][]interface{} `json:"taskData"`
}
