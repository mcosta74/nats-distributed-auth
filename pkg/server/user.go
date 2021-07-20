package server

type User struct {
	UserName         string `json:"username"`
	Password         string `json:"password,omitempty"`
	ForbiddenDevices []int  `json:"forbidden_devices"`
}
