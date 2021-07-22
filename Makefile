.PHONY: run_service
run_service: 
	go run cmd/authsvc/authsvc.go

.PHONY: run_nats
run_nats:
	nats-server -DV -c nats/server.conf

.PHONY: run_client
run_client:
	go run cmd/authclient/authclient.go
