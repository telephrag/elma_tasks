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

	"github.com/go-chi/chi"
)

func GetDataAndSolve(rw http.ResponseWriter, r *http.Request) {

	var err error = nil
	var errMsg string = ""
	defer util.Ordinary500(rw, err, errMsg)

	taskName := chi.URLParam(r, "name")
	var tasks []string

	if taskName != "" { // consider moving this to middleware
		tasks = append(tasks, taskName)
	} else {
		tasks = config.TaskNames[:]
	}

	tdc, err := GetTasks(tasks, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	reqBody, err := json.Marshal(tdc)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	postReq, err := http.NewRequest(
		"POST",
		config.ProtoHttp+config.LocalInternalAddr+"/tasks/solution",
		bytes.NewBuffer(reqBody),
	)
	postReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		rw.WriteHeader(resp.StatusCode)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
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