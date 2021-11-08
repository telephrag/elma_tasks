package handlerfuncs

import (
	"errors"
	"io/ioutil"
	"mservice/config"
	"mservice/util"
	"net/http"

	"github.com/go-chi/chi"
)

func GetTaskData(rw http.ResponseWriter, r *http.Request) {
	var err error
	var code int = 0
	defer util.HttpErrWriter(rw, err, code)

	taskName := chi.URLParam(r, "name")

	resp, err := http.Get("http://" + config.DataProviderAddr + "/tasks/" + taskName)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New("")
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
