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
	user, err := s.r.AddUser(userName)
	if err != nil {
		return user, err
	}

	err = NscAddUser(user.UserName)
	if err != nil {
		// the logic here sucks (in production we would use a transactional system)
		s.r.DeleteUser(user.UserName)
		return User{}, err
	}
	return user, nil
}

func (s *authService) AddForbiddenDevice(userId string, deviceId int) error {
	return s.r.AddForbiddenDevice(userId, deviceId)
}

func (s *authService) Dump() error {
	return s.r.Dump()
}
