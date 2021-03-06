package util

import (
	"net/http"
)

func HttpErrWriter(rw http.ResponseWriter, err error, code int) {
	if err == nil {
		return
	}

	c := http.StatusInternalServerError
	if code != 0 {
		c = code
	}

	http.Error(rw, ": "+err.Error(), c)
}
