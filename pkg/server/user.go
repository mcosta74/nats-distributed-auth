package server

type User struct {
	UserName         string `json:"user_name,omitempty"`
	ForbiddenDevices []int  `json:"forbidden_devices,omitempty"`
}
