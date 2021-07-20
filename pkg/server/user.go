package server

type User struct {
	UserName         string `json:"user_name"`
	Password         string `json:"password,omitempty"`
	ForbiddenDevices []int  `json:"forbidden_devices"`
}
