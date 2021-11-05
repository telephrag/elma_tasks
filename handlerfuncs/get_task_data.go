package handlerfuncs

import (
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
