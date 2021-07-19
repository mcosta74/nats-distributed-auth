package server

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addUserRequest struct {
	UserName string `json:"user_name,omitempty"`
}

type addUserResponse struct {
	V   User   `json:"v,omitempty"`
	Err string `json:"err,omitempty"`
}

type addForbiddenDeviceRequest struct {
	UserName string `json:"user_name,omitempty"`
	DeviceId int    `json:"device_id,omitempty"`
}

type addForbiddenDeviceResponse struct {
	Err string `json:"err,omitempty"`
}

func makeAddUserEndpoint(s AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addUserRequest)
		v, err := s.AddUser(req.UserName)
		if err != nil {
			return addUserResponse{v, err.Error()}, nil
		}
		return addUserResponse{v, ""}, nil
	}
}

func makeAddForbiddenDeviceEndpoint(s AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addForbiddenDeviceRequest)
		err := s.AddForbiddenDevice(req.UserName, req.DeviceId)
		return addForbiddenDeviceResponse{err.Error()}, nil
	}
}
