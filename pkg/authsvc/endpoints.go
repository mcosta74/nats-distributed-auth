package authsvc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type addUserRequest struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type addUserResponse struct {
	User
}

type getUserRequest struct {
	UserName string `json:"username,omitempty"`
}

type getUserResponse struct {
	User
}
type addForbiddenDeviceRequest struct {
	UserName string `json:"username,omitempty"`
	DeviceId int    `json:"device_id,omitempty"`
}

type addForbiddenDeviceResponse struct{}

type loginRequest struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type loginResponse struct {
	User
}

var (
	ErrInvalidRequest = errors.New("invalid request")
)

func makeAddUserEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addUserRequest)

		if req.UserName == "" || req.Password == "" {
			return addUserResponse{}, ErrDuplicatedUserName
		}

		v, err := s.AddUser(ctx, req.UserName, req.Password)
		if err != nil {
			return addUserResponse{}, err
		}
		return addUserResponse{v}, nil
	}
}

func makeAddForbiddenDeviceEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addForbiddenDeviceRequest)
		err := s.AddForbiddenDevice(ctx, req.UserName, req.DeviceId)
		return addForbiddenDeviceResponse{}, err
	}
}

func makeGetUserEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		user, err := s.GetUser(ctx, req.UserName)
		return getUserResponse{user}, err
	}
}

func makeLoginEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		user, err := s.Login(ctx, req.UserName, req.Password)
		return loginResponse{user}, err
	}
}
