package mcsmapi

type AddDaemonRequest struct {
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	Prefix    string `json:"prefix"`
	Title     string `json:"remarks"`
	AccessKey string `json:"apiKey"`
}
