package server

type AuthService interface {
	AddUser(userName string) (User, error)
	AddForbiddenDevice(userId string, deviceId int) error
	Dump() error
}

func NewAuthService(r Repository) AuthService {
	return &authService{
		r: r,
	}
}

type authService struct {
	r Repository
}

func (s *authService) AddUser(userName string) (User, error) {
	return s.r.AddUser(userName)
}

func (s *authService) AddForbiddenDevice(userId string, deviceId int) error {
	return s.r.AddForbiddenDevice(userId, deviceId)
}

func (s *authService) Dump() error {
	return s.r.Dump()
}
