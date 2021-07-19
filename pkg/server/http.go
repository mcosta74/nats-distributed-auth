package server

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(svc AuthService) http.Handler {
	r := mux.NewRouter()

	addUserHandler := kithttp.NewServer(
		makeAddUserEndpoint(svc),
		decodeAddUserRequest,
		encodeResponse,
	)

	addForbiddenDeviceHandler := kithttp.NewServer(
		makeAddForbiddenDeviceEndpoint(svc),
		decodeAddForbiddenDeviceRequest,
		encodeResponse,
	)

	r.Methods("POST").Path("/add-user").Handler(addUserHandler)
	r.Methods("POST").Path("/add-forbidden-device").Handler(addForbiddenDeviceHandler)

	return r
}

func SetupHTTPHandlers(svc AuthService) {
	addUserHandler := kithttp.NewServer(
		makeAddUserEndpoint(svc),
		decodeAddUserRequest,
		encodeResponse,
	)

	addForbiddenDeviceHandler := kithttp.NewServer(
		makeAddForbiddenDeviceEndpoint(svc),
		decodeAddForbiddenDeviceRequest,
		encodeResponse,
	)

	http.Handle("/add-user", addUserHandler)
	http.Handle("/add-forbidden-device", addForbiddenDeviceHandler)
}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAddForbiddenDeviceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addForbiddenDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
