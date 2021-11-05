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

func PostBack(rw http.ResponseWriter, r *http.Request) {

	var errMsg string = ""
	var err error = nil
	defer util.Ordinary500(rw, err, errMsg)

	res, ok := r.Context().Value("results").([]models.Result)
	if !ok {
		http.Error(rw, "Wrong type arrived with request.", http.StatusInternalServerError)
	}

	resp, err := PostTasks(res, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Write([]byte(fmt.Sprint(resp)))
}

func PostTask(res models.Result, rw http.ResponseWriter) (*http.Response, error) {

	req, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	rw.Header().Set("Content-type", "application/json")

	resp, err := http.Post("http://"+config.DataProviderAddr+"/tasks/solution", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		rw.WriteHeader(resp.StatusCode)
		return nil, errors.New(fmt.Sprint(resp.StatusCode))
	}

	return resp, nil
}

func PostTasks(res []models.Result, rw http.ResponseWriter) ([]models.Response, error) {

	rc := make([]models.Response, len(res))
	for i := range res {
		resp, err := PostTask(res[i], rw)
		if err != nil {
			return nil, err
		}

		content, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		var r models.Response
		err = json.Unmarshal(content, &r)
		if err != nil {
			return nil, err
		}

		rc[i] = r
	}

	return rc, nil
}
