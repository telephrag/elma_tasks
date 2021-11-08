package handlerfuncs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mservice/config"
	"mservice/models"
	"mservice/util"
	"net/http"
)

func GetDataAndSolve(rw http.ResponseWriter, r *http.Request) {

	var err error
	var code int
	defer util.HttpErrWriter(rw, err, code)

	tasks, ok := r.Context().Value("tasks").([]string)
	if !ok {
		err = errors.New(config.CtxWrongType)
		return
	}

	tc, err := GetTasks(tasks, rw)
	if err != nil {
		return
	}

	reqBody, err := json.Marshal(tc)
	if err != nil {
		return
	}

	postReq, err := http.NewRequest(
		"POST",
		config.ProtoHttp+config.LocalInternalAddr+"/tasks/solution",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return
	}
	postReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		code = resp.StatusCode
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	rw.Write(content)
}

func GetTask(taskName string, rw http.ResponseWriter) (*http.Response, error) {

	resp, err := http.Get(
		config.ProtoHttp + config.LocalInternalAddr + "/tasks/" + taskName,
	)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprint(resp.StatusCode))
	}

	return resp, nil
}

func GetTasks(tasks []string, rw http.ResponseWriter) ([]models.Task, error) {
	tc := make([]models.Task, len(tasks))
	for i := range tasks {
		resp, err := GetTask(tasks[i], rw)
		if err != nil {
			return nil, err
		}

		content, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		var data [][]interface{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			return nil, err
		}

		tc[i] = models.Task{
			TaskName: tasks[i],
			Data:     data,
		}
	}

	return tc, nil
}
