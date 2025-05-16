package mcsmapi

import (
	"github.com/google/go-querystring/query"
	"log"
)

// BuildQueryString constructs the URL-encoded query string
func BuildQueryString(q interface{}) string {
	v, err := query.Values(q)
	if err != nil {
		log.Fatal(err)
	}
	return v.Encode()
}

// BaseRequest includes common query parameters used across many request types
type BaseRequest struct {
	DaemonID string `url:"daemonId"`
	UUID     string `url:"uuid"`
}

type BaseSuccessResponse struct {
	Status int   `json:"status"`
	Data   bool  `json:"data"`
	Time   int64 `json:"time"`
}

func (q *BaseRequest) BuildQueryString() string {
	return BuildQueryString(q)
}

type BoolResponse struct {
	Data bool `json:"data"`
	CommonResponse
}

type MethodClient struct {
	endpoint string
	client   *Client
}
