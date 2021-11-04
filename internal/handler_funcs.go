package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mservice/config"
	"mservice/util"
	"net/http"

	"github.com/go-chi/chi"
)

func GetTaskData(rw http.ResponseWriter, r *http.Request) {
	var errMsg string = ""
	var err error = nil
	defer util.Ordinary500(rw, err, errMsg)

	taskName := chi.URLParam(r, "name")

	resp, err := http.Get("http://" + config.DataProviderAddr + "/tasks/" + taskName)
	if err != nil {
		errMsg = "Failed to get data from provider.\n"
		return
	}

	if resp.StatusCode != http.StatusOK {
		rw.WriteHeader(resp.StatusCode)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		errMsg = "Failed to read response body.\n"
		return
	}

	rw.Write(content)
}

func PostBack(rw http.ResponseWriter, r *http.Request) {

	var errMsg string = ""
	var err error = nil
	defer util.Ordinary500(rw, err, errMsg)

	rw.Header().Set("Content-type", "application/json")

	rj := r.Context().Value("result")

	req, err := json.Marshal(rj)
	if err != nil {
		errMsg = "Failed to marshal result.\n"
		return
	}

	resp, err := http.Post("http://"+config.DataProviderAddr+"/tasks/solution", "application/json", bytes.NewBuffer(req))
	if err != nil {
		errMsg = "Failed to post data back to the provider.\n"
	}

	if resp.StatusCode != http.StatusOK {
		rw.WriteHeader(resp.StatusCode)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		errMsg = "Failed to read response body."
		return
	}
	rw.Write(content)
}
