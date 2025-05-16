package mcsmapi

import (
	"github.com/google/go-querystring/query"
	"log"
)

type TryConnectDaemonQuery struct {
	UUID string `url:"uuid"`
}

func (q *TryConnectDaemonQuery) BuildQueryString() string {
	v, err := query.Values(q)
	if err != nil {
		log.Fatal(err)
	}
	return v.Encode()
}

type DaemonConfigActionQuery struct {
	UUID string `url:"uuid"`
}

func (q *DaemonConfigActionQuery) BuildQueryString() string { return BuildQueryString(q) }
