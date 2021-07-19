package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mcosta74/nats-distributed-auth/pkg/server"
)

func main() {
	fmt.Println("Hello World")

	var (
		rep = server.NewRepository()
		s   = server.NewAuthService(rep)
	)
	s.AddUser("massimo")
	s.AddForbiddenDevice("massimo", 12)

	s.AddUser("foo")

	handler := server.MakeHandler(s)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
