package server

import (
	"errors"
	"log"
)

type InMemoryRepository map[string]User

type Repository interface {
	AddUser(userName string) (User, error)
	AddForbiddenDevice(userId string, deviceId int) error
	Dump() error
}

func NewRepository() Repository {
	return &authRepository{
		data: make(InMemoryRepository),
	}
}

type authRepository struct {
	data InMemoryRepository
}

func (r *authRepository) AddUser(userName string) (User, error) {
	_, already := r.data[userName]
	if already {
		return User{}, errors.New("user already present")
	}

	user := User{
		UserName:         userName,
		ForbiddenDevices: make([]int, 0),
	}

	r.data[userName] = user

	return user, nil
}

func (r *authRepository) AddForbiddenDevice(userId string, deviceId int) error {
	user, prs := r.data[userId]
	if !prs {
		return errors.New("no user found")
	}

	user.ForbiddenDevices = append(user.ForbiddenDevices, deviceId)
	r.data[userId] = user
	return nil
}

func (r *authRepository) Dump() error {
	log.Printf("DUMP:\n%v", r.data)

	return nil
}
