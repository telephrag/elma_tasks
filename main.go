package main

import (
	"mservice/config"
	"mservice/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go func() {
		err := http.ListenAndServe(
			config.LocalAddr,
			services.GetTaskDataAndSolve(),
		)

		if err != nil {
			panic(err)
		}

	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-interrupt
}
