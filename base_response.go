package mcsmapi

type BaseResponse struct {
	Status int   `json:"status"`
	Data   any   `json:"data"`
	Time   int64 `json:"time"`
}
