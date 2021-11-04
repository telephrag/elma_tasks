package util

import (
	"net/http"
)

func Ordinary500(rw http.ResponseWriter, err error, msg string) {
	if err == nil {
		return
	}

	rw.Write([]byte(err.Error() + ": " + msg))
	rw.WriteHeader(http.StatusInternalServerError)
}
