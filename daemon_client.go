package mcsmapi

import (
	"encoding/json"
	"fmt"
)

type daemonClient MethodClient

func newDaemonClient(client *Client) *daemonClient {
	return &daemonClient{client: client, endpoint: "service"}
}

// Add sends a POST request to add a new daemon using the given request payload.
func (dc *daemonClient) Add(req *AddDaemonRequest) (*BaseResponse, error) {
	return dc.doJSONRequest("POST", "remote_service", req)
}

// Delete sends a DELETE request to remove a daemon based on query parameters.
func (dc *daemonClient) Delete(query *DaemonConfigActionQuery) (*BaseResponse, error) {
	endpoint := fmt.Sprintf("remote_service?%s", query.BuildQueryString())
	return dc.doRawRequest("DELETE", endpoint)
}

// TryConnect attempts to establish a connection to a remote daemon using query parameters.
func (dc *daemonClient) TryConnect(query *TryConnectDaemonQuery) (*BaseResponse, error) {
	endpoint := fmt.Sprintf("link_remote_service?%s", query.BuildQueryString())
	return dc.doRawRequest("GET", endpoint)
}

// UpdateConnectConfig updates the connection configuration for a daemon.
func (dc *daemonClient) UpdateConnectConfig(query *UpdateConfigQuery) (*BaseResponse, error) {
	endpoint := fmt.Sprintf("remote_service?%s", query.BuildQueryString())
	return dc.doRawRequest("PUT", endpoint)
}

func (dc *daemonClient) doJSONRequest(method, endpoint string, body any) (*BaseResponse, error) {
	var result BaseResponse
	err := dc.client.doRequestAndDecode(method, dc.endpoint+"/"+endpoint, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (dc *daemonClient) doRawRequest(method, endpoint string) (*BaseResponse, error) {
	resp, err := dc.client.sendRequest(method, dc.endpoint+"/"+endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result BaseResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
