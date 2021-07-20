package authsvc

import (
	"context"
	"errors"
)

var (
	ErrAuthenticationFailed = errors.New("authentication failed")
)

type AuthService interface {
	AddUser(ctx context.Context, userName, password string) (User, error)
	GetUser(ctx context.Context, userName string) (User, error)
	AddForbiddenDevice(ctx context.Context, userId string, deviceId int) error
	Login(ctx context.Context, userName, password string) (User, error)
}

func NewAuthService(r Repository) AuthService {
	return &authService{
		r: r,
	}
}

type authService struct {
	r Repository
}

func (s *authService) AddUser(_ context.Context, userName, password string) (User, error) {
	user, err := s.r.AddUser(userName, password)
	if err != nil {
		return user, err
	}

	// err = NscAddUser(user.UserName)
	// if err != nil {
	// 	// the logic here sucks (in production we would use a transactional system)
	// 	s.r.DeleteUser(user.UserName)
	// 	return User{}, err
	// }
	return user, nil
}

func (s authService) GetUser(_ context.Context, userName string) (User, error) {
	return s.r.GetUser(userName)
}

func (s *authService) AddForbiddenDevice(_ context.Context, userId string, deviceId int) error {
	return s.r.AddForbiddenDevice(userId, deviceId)
}

func (s authService) Login(_ context.Context, userName, password string) (User, error) {
	user, err := s.r.FindUser(userName, password)
	if err != nil {
		return User{}, ErrAuthenticationFailed
	}
	return user, nil
}
