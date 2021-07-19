.PHONY: run_server
run_server: cmd/server/main.go
	go run cmd/server/main.go

.PHONY: run_nats
run_nats:
	nats-server -DV -c nats/server.conf
