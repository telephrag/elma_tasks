package main

import (
	"context"
	"mservice/config"
	"mservice/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	rootCtx := context.Background()

	// TODO: make task selection
	var taskName string = config.CycliclShift
	taskCtx := context.WithValue(rootCtx, "name", taskName) // TODO: decide what to do with key
	//services.SolveTaskService(taskCtx)

	go func(taskCtx context.Context) {
		err := http.ListenAndServe(
			config.LocalAddr,
			services.SolveTaskService(taskCtx),
		)

		if err != nil {
			// TODO: add error handling
		}

	}(taskCtx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT) // stop with "Ctrl+C"
	<-interrupt
}
