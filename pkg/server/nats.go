package server

import "github.com/nats-io/nats.go"

func Connect(name string) (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)

	return nc, err
}
