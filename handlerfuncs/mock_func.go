package handlerfuncs

import (
	"bytes"
	"io/ioutil"
	"mservice/config"
	"mservice/util"
	"net/http"
)

func Mock(rw http.ResponseWriter, r *http.Request) {
	var err error
	var code int
	defer util.HttpErrWriter(rw, err, code)

	mockReq, err := http.NewRequest(
		"GET",
		config.ProtoHttp+config.LocalPublicAddr+"/tasks/",
		bytes.NewBuffer([]byte("")),
	)
	if err != nil {
		return
	}
	mockReq.Header.Set("tasks", config.MockTasksSet)

	resp, err := http.DefaultClient.Do(mockReq)
	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	rw.Write(content)
}
