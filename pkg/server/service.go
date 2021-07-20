package server

import "errors"

// import "github.com/go-kit/log"

var (
	ErrAuthenticationFailed = errors.New("authentication failed")
)

type AuthService interface {
	AddUser(userName, password string) (User, error)
	GetUser(userName string) (User, error)
	AddForbiddenDevice(userId string, deviceId int) error
	Login(userName, password string) (User, error)
}

func NewAuthService(r Repository) AuthService {
	return &authService{
		r: r,
	}
}

type authService struct {
	r Repository
}

func (s *authService) AddUser(userName, password string) (User, error) {
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

func (s authService) GetUser(userName string) (User, error) {
	return s.r.GetUser(userName)
}

func (s *authService) AddForbiddenDevice(userId string, deviceId int) error {
	return s.r.AddForbiddenDevice(userId, deviceId)
}

func (s authService) Login(userName, password string) (User, error) {
	user, err := s.r.FindUser(userName, password)
	if err != nil {
		return User{}, ErrAuthenticationFailed
	}
	return user, nil
}
