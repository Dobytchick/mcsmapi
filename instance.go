package mcsmapi

// InstanceListResponse represents the response structure for a list of instances.
type InstanceListResponse struct {
	Status int `json:"status"`
	Data   struct {
		MaxPage  int               `json:"maxPage"`
		PageSize int               `json:"pageSize"`
		AllTags  []string          `json:"allTags"`
		Data     []*InstanceDetail `json:"data"`
	} `json:"data"`
	Time int64 `json:"time"`
}

// InstanceDetailResponse represents the response structure for a single instance.
type InstanceDetailResponse struct {
	Status int             `json:"status"`
	Data   *InstanceDetail `json:"data"`
	Time   int64           `json:"time"`
}

// CreateInstanceResponse represents the response structure after creating an instance.
type CreateInstanceResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string          `json:"instanceUuid"`
		Config       *InstanceConfig `json:"config"`
	} `json:"data"`
	Time int64 `json:"time"`
}

type UpdateConfigResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string `json:"instanceUuid"`
	} `json:"data"`
	Time int64 `json:"time"`
}

// DeleteInstancesResponse represents the response for deleting instances.
type DeleteInstancesResponse struct {
	Status int      `json:"status"`
	Data   []string `json:"data"`
	Time   int64    `json:"time"`
}

// StartInstanceResponse represents the response for starting an instance.
type StartInstanceResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string `json:"instanceUuid"`
	} `json:"data"`
	Time int64 `json:"time"`
}

// StopInstanceResponse represents the response for stopping an instance.
type StopInstanceResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string `json:"instanceUuid"`
	} `json:"data"`
	Time int64 `json:"time"`
}

// RestartInstanceResponse represents the response for restarting an instance.
type RestartInstanceResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string `json:"instanceUuid"`
	} `json:"data"`
	Time int64 `json:"time"`
}

// KillInstanceResponse represents the response for killing an instance.
type KillInstanceResponse struct {
	Status int `json:"status"`
	Data   struct {
		InstanceUUID string `json:"instanceUuid"`
	} `json:"data"`
	Time int64 `json:"time"`
}
