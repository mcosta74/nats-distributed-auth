package main

import (
	"log"
	"net/http"
	"os"

	kitlog "github.com/go-kit/log"
	"github.com/mcosta74/nats-distributed-auth/pkg/server"
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
		rep = server.NewRepository()
		s   = server.NewAuthService(rep)
	)

	handler := server.MakeHandler(s)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
