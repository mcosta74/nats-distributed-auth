package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	"github.com/mcosta74/nats-distributed-auth/pkg/authsvc"
)

func main() {
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	}

	logger.Log("msg", "Service Started")
	defer func() {
		logger.Log("msg", "Service Stopped")
	}()

	var (
		rep = authsvc.NewRepository()
		s   = authsvc.NewAuthService(rep)
	)

	handler := authsvc.MakeHandler(s)

	done := make(chan bool)

	go func() {
		err := http.ListenAndServe(":8080", handler)
		logger.Log("msg", fmt.Sprintf("Server stopped: %s", err))
		done <- true
	}()

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		logger.Log("msg", fmt.Sprintf("System Received signal: %s", sig))
		done <- true
	}()

	<-done
}
