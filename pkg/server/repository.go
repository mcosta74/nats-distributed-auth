package server

import (
	"errors"
)

var (
	ErrDuplicatedUserName = errors.New("duplicated username")
	ErrUserNotFound       = errors.New("user not found")
)

type InMemoryRepository map[string]User

type Repository interface {
	AddUser(userName, password string) (User, error)
	GetUser(userName string) (User, error)
	DeleteUser(userName string) error
	AddForbiddenDevice(userId string, deviceId int) error
	FindUser(userName, password string) (User, error)
}

func NewRepository() Repository {
	return &authRepository{
		data: make(InMemoryRepository),
	}
}

type authRepository struct {
	data InMemoryRepository
}

func (r *authRepository) AddUser(userName, password string) (User, error) {
	_, already := r.data[userName]
	if already {
		return User{}, ErrDuplicatedUserName
	}

	user := User{
		UserName:         userName,
		Password:         password,
		ForbiddenDevices: make([]int, 0),
	}
	r.data[userName] = user

	user.Password = ""
	return user, nil
}

func (r authRepository) GetUser(userId string) (User, error) {
	user, prs := r.data[userId]
	if !prs {
		return User{}, ErrUserNotFound
	}
	user.Password = ""
	return user, nil
}

func (r *authRepository) DeleteUser(userId string) error {
	delete(r.data, userId)

	return nil
}

func (r *authRepository) AddForbiddenDevice(userId string, deviceId int) error {
	user, prs := r.data[userId]
	if !prs {
		return ErrUserNotFound
	}

	user.ForbiddenDevices = append(user.ForbiddenDevices, deviceId)
	r.data[userId] = user
	return nil
}

func (r authRepository) FindUser(userName, password string) (User, error) {
	user, prs := r.data[userName]
	if !prs || user.Password != password {
		return User{}, ErrUserNotFound
	}
	user.Password = ""
	return user, nil
}
