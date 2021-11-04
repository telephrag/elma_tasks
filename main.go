package main

import (
	"mservice/config"
	"mservice/internal"
	"mservice/public"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go func() {
		err := http.ListenAndServe(
			config.LocalPublicAddr,
			public.GetDataAndSolve(),
		)

		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := http.ListenAndServe(
			config.LocalInternalAddr,
			internal.InternalService(),
		)

		if err != nil {
			panic(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-interrupt
}
