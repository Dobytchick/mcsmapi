package mcsmapi

type Query interface {
	BuildQueryString() string
}
