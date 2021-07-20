package server

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type addUserRequest struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type addUserResponse struct {
	User
}

type getUserRequest struct {
	UserName string `json:"user_name,omitempty"`
}

type getUserResponse struct {
	User
}
type addForbiddenDeviceRequest struct {
	UserName string `json:"user_name,omitempty"`
	DeviceId int    `json:"device_id,omitempty"`
}

type addForbiddenDeviceResponse struct{}

var (
	ErrInvalidRequest = errors.New("invalid request")
)

func makeAddUserEndpoint(s AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addUserRequest)

		if req.UserName == "" || req.Password == "" {
			return addUserResponse{}, ErrDuplicatedUserName
		}

		v, err := s.AddUser(req.UserName, req.Password)
		if err != nil {
			return addUserResponse{}, err
		}
		return addUserResponse{v}, nil
	}
}

func makeAddForbiddenDeviceEndpoint(s AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addForbiddenDeviceRequest)
		err := s.AddForbiddenDevice(req.UserName, req.DeviceId)
		return addForbiddenDeviceResponse{}, err
	}
}

func makeGetUserEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		user, err := s.GetUser(req.UserName)
		return getUserResponse{user}, err
	}
}
