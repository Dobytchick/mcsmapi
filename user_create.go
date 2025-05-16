package mcsmapi

type CreateUserRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"` // 1=User, 10=Admin, -1=Banned
}

type CreateUserResponse struct {
	Status int             `json:"status"`
	Time   int64           `json:"time"`
	Data   CreatedUserData `json:"data"`
}

type CreatedUserData struct {
	UUID       string `json:"uuid"`
	UserName   string `json:"userName"`
	Permission int    `json:"permission"`
}
