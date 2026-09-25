package mcsmapi

import (
	"fmt"
	"net/http"
	"net/url"
)

// instanceClient wraps the base MethodClient for working with instances.
type instanceClient MethodClient

// newInstanceClient initializes a new instanceClient.
func newInstanceClient(client *Client) *instanceClient {
	return &instanceClient{client: client, endpoint: ""}
}

// sendRequest delegates the request to the base client, prefixing the endpoint.
func (ic *instanceClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	return ic.client.sendRequest(method, endpoint, body)
}

// doRequestAndDecode performs the request and decodes JSON response into the provided output.
func (ic *instanceClient) doRequestAndDecode(method, endpoint string, body, out any) error {
	return ic.client.doRequestAndDecode(method, endpoint, body, out)
}

// GetList retrieves a list of remote service instances using query parameters.
func (ic *instanceClient) GetList(req *ListInstancesQuery) (*InstanceListResponse, error) {
	var res InstanceListResponse
	err := ic.doRequestAndDecode("GET", "service/remote_service_instances?"+req.BuildQueryString(), nil, &res)
	return &res, err
}

// GetDetail retrieves details about a specific instance.
func (ic *instanceClient) GetDetail(req *GetInstanceQuery) (*InstanceDetailResponse, error) {
	var res InstanceDetailResponse
	err := ic.doRequestAndDecode("GET", "instance?"+req.BuildQueryString(), nil, &res)
	return &res, err
}

// Create creates a new instance with a given daemon ID and configuration.
func (ic *instanceClient) Create(daemonID string, config *InstanceConfig) (*CreateInstanceResponse, error) {
	var res CreateInstanceResponse
	err := ic.doRequestAndDecode("POST", "instance?daemonId="+url.QueryEscape(daemonID), config, &res)
	return &res, err
}

// Delete removes one or more instances associated with a daemon ID.
func (ic *instanceClient) Delete(daemonID string, body *DeleteInstanceBody) (*DeleteInstancesResponse, error) {
	var res DeleteInstancesResponse
	err := ic.doRequestAndDecode("DELETE", "instance?daemonId="+url.QueryEscape(daemonID), body, &res)
	return &res, err
}

// Start launches a protected instance using its UUID and daemon ID.
func (ic *instanceClient) Start(instanceID, daemonID string) (*StartInstanceResponse, error) {
	var res StartInstanceResponse
	endpoint := fmt.Sprintf("protected_instance/open?uuid=%s&daemonId=%s", url.QueryEscape(instanceID), url.QueryEscape(daemonID))
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// Stop gracefully stops a running protected instance.
func (ic *instanceClient) Stop(instanceID, daemonID string) (*StopInstanceResponse, error) {
	var res StopInstanceResponse
	endpoint := fmt.Sprintf("protected_instance/stop?uuid=%s&daemonId=%s", url.QueryEscape(instanceID), url.QueryEscape(daemonID))
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// Restart restarts a protected instance.
func (ic *instanceClient) Restart(instanceID, daemonID string) (*RestartInstanceResponse, error) {
	var res RestartInstanceResponse
	endpoint := fmt.Sprintf("protected_instance/restart?uuid=%s&daemonId=%s", url.QueryEscape(instanceID), url.QueryEscape(daemonID))
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// Kill forcibly stops a protected instance.
func (ic *instanceClient) Kill(instanceID, daemonID string) (*KillInstanceResponse, error) {
	var res KillInstanceResponse
	endpoint := fmt.Sprintf("protected_instance/kill?uuid=%s&daemonId=%s", url.QueryEscape(instanceID), url.QueryEscape(daemonID))
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// SendCommand sends a command to a running instance.
func (ic *instanceClient) SendCommand(instanceID, daemonID, command string) (*BaseSuccessResponse, error) {
	var res BaseSuccessResponse
	params := url.Values{"uuid": {instanceID}, "daemonId": {daemonID}, "command": {command}}
	endpoint := "protected_instance/command?" + params.Encode()
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}
