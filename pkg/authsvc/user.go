package authsvc

type User struct {
	UserName         string `json:"username"`
	Password         string `json:"-"`
	ForbiddenDevices []int  `json:"forbidden_devices"`
}
