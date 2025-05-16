package mcsmapi

type DeleteInstanceBody struct {
	UUIDS      []string `json:"uuids"`
	DeleteFile bool     `json:"deleteFile"`
}

// ListInstancesQuery defines the query parameters for listing remote service instances.
type ListInstancesQuery struct {
	DaemonID     string         `json:"daemonId" url:"daemonId"`
	Page         int            `json:"page" url:"page"`
	PageSize     int            `json:"page_size" url:"page_size"`
	InstanceName string         `json:"instance_name" url:"instance_name"`
	Status       InstanceStatus `json:"status" url:"status"`
}

func (q *ListInstancesQuery) BuildQueryString() string { return BuildQueryString(q) }

// GetInstanceQuery defines the query parameters for retrieving a single instance by UUID.
type GetInstanceQuery struct {
	UUID     string `json:"uuid" url:"uuid"`
	DaemonID string `json:"daemonId" url:"daemonId"`
}

func (q *GetInstanceQuery) BuildQueryString() string { return BuildQueryString(q) }

// UpdateConfigQuery defines the query parameters for updating an instance config.
type UpdateConfigQuery struct {
	UUID     string `url:"uuid"`
	DaemonID string `url:"daemonId"`
}

func (q *UpdateConfigQuery) BuildQueryString() string { return BuildQueryString(q) }

// StartInstanceQuery defines the query parameters for starting an instance.
type StartInstanceQuery struct {
	UUID     string `url:"uuid"`
	DaemonID string `url:"daemonId"`
}

func (q *StartInstanceQuery) BuildQueryString() string { return BuildQueryString(q) }

// StopInstanceQuery defines the query parameters for stopping an instance.
type StopInstanceQuery struct {
	UUID     string `url:"uuid"`
	DaemonID string `url:"daemonId"`
}

func (q *StopInstanceQuery) BuildQueryString() string { return BuildQueryString(q) }

// RestartInstanceQuery defines the query parameters for restarting an instance.
type RestartInstanceQuery struct {
	UUID     string `url:"uuid"`
	DaemonID string `url:"daemonId"`
}

func (q *RestartInstanceQuery) BuildQueryString() string { return BuildQueryString(q) }

// KillInstanceQuery defines the query parameters for killing an instance.
type KillInstanceQuery struct {
	UUID     string `url:"uuid"`
	DaemonID string `url:"daemonId"`
}

func (q *KillInstanceQuery) BuildQueryString() string { return BuildQueryString(q) }

// BatchOperationQuery defines the query parameters for batch operations (start, stop, restart, kill).
type BatchOperationQuery struct {
	InstanceUUID string `url:"instanceUuid"`
	DaemonID     string `url:"daemonId"`
}

func (q *BatchOperationQuery) BuildQueryString() string { return BuildQueryString(q) }
