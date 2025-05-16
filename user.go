package mcsmapi

type UserRole string

const (
	RoleUser   UserRole = "1"
	RoleAdmin  UserRole = "10"
	RoleBanned UserRole = "-1"
)

type UserListResponse struct {
	Status int          `json:"status"`
	Data   UserListData `json:"data"`
	Time   int64        `json:"time"`
}

// UserListData represents a paginated list of users returned by the API.
type UserListData struct {
	Data     []UserData `json:"data"`
	MaxPage  int        `json:"maxPage"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	Total    int        `json:"total"`
}

// UserData represents user information and authentication details.
type UserData struct {
	UUID         string         `json:"uuid"`
	UserName     string         `json:"userName"`
	PassWord     string         `json:"passWord"`
	PassWordType int            `json:"passWordType"`
	Salt         string         `json:"salt"`
	Permission   int            `json:"permission"`
	RegisterTime string         `json:"registerTime"`
	LoginTime    string         `json:"loginTime"`
	Instances    []UserInstance `json:"instances"`
	ApiKey       string         `json:"apiKey"`
	IsInit       bool           `json:"isInit"`
	Secret       string         `json:"secret"`
	Open2FA      bool           `json:"open2FA"`
}

type UserInstance struct {
	InstanceUUID string `json:"instanceUuid"`
	DaemonID     string `json:"daemonId"`
}
