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

	options := []kithttp.ServerOption{
		// kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	addUserHandler := kithttp.NewServer(
		makeAddUserEndpoint(svc),
		decodeAddUserRequest,
		kithttp.EncodeJSONResponse,
		options...,
	)

	addForbiddenDeviceHandler := kithttp.NewServer(
		makeAddForbiddenDeviceEndpoint(svc),
		decodeAddForbiddenDeviceRequest,
		kithttp.EncodeJSONResponse,
		options...,
	)

	getUserHandler := kithttp.NewServer(
		makeGetUserEndpoint(svc),
		decodeGetUserRequest,
		kithttp.EncodeJSONResponse,
		options...,
	)

	loginHandler := kithttp.NewServer(
		makeLoginEndpoint(svc),
		decodeLoginRequest,
		kithttp.EncodeJSONResponse,
		options...,
	)

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.Path("").Methods("POST").Handler(addUserHandler)

	userRouter := usersRouter.PathPrefix("/{userId}").Subrouter()
	userRouter.Path("").Methods("GET").Handler(getUserHandler)
	userRouter.Path("/forbidden-devices").Methods("POST").Handler(addForbiddenDeviceHandler)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Path("/login").Methods("POST").Handler(loginHandler)

	return r
}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserRequest

	request.UserName = mux.Vars(r)["userId"]
	return request, nil
}

func decodeAddForbiddenDeviceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addForbiddenDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.UserName = mux.Vars(r)["userId"]
	return request, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError called without error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFromErr(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFromErr(err error) int {
	switch err {
	case ErrUserNotFound:
		return http.StatusNotFound
	case ErrDuplicatedUserName, ErrInvalidRequest:
		return http.StatusBadRequest
	case ErrAuthenticationFailed:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func (r addUserResponse) StatusCode() int {
	return http.StatusCreated
}

func (r addForbiddenDeviceResponse) StatusCode() int {
	return http.StatusNoContent
}
