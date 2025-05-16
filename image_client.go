package mcsmapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// imageClient provides methods for working with Docker images, containers, and networks via the "environment" endpoint.
type imageClient MethodClient

func newImageClient(client *Client) *imageClient {
	return &imageClient{client: client, endpoint: "environment"}
}

// sendRequest delegates the request to the base client, prefixing the endpoint.
func (ic *imageClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	fullEndpoint := ic.endpoint
	if endpoint != "" {
		fullEndpoint += "/" + endpoint
	}
	return ic.client.sendRequest(method, fullEndpoint, body)
}

// doRequestAndDecode performs the request and decodes JSON response into the provided output.
func (ic *imageClient) doRequestAndDecode(method, endpoint string, body, out any) error {
	resp, err := ic.sendRequest(method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(out)
}

// GetImageList retrieves the list of Docker images for a daemon.
// GET /environment/image?daemonId=...
func (ic *imageClient) GetImageList(daemonId string) (*GetImageListResponse, error) {
	var res GetImageListResponse
	endpoint := fmt.Sprintf("image?daemonId=%s", daemonId)
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// GetContainerList retrieves the list of Docker containers for a daemon.
// GET /environment/containers?daemonId=...
func (ic *imageClient) GetContainerList(daemonId string) (*GetContainerListResponse, error) {
	var res GetContainerListResponse
	endpoint := fmt.Sprintf("containers?daemonId=%s", daemonId)
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// GetNetworkList retrieves the list of Docker networks for a daemon.
// GET /environment/network?daemonId=...
func (ic *imageClient) GetNetworkList(daemonId string) (*GetNetworkListResponse, error) {
	var res GetNetworkListResponse
	endpoint := fmt.Sprintf("network?daemonId=%s", daemonId)
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}

// CreateImage builds a new Docker image on the daemon.
// POST /environment/image?daemonId=...
func (ic *imageClient) CreateImage(daemonId string, req *CreateImageRequest) (*BaseSuccessResponse, error) {
	var res BaseSuccessResponse
	endpoint := fmt.Sprintf("image?daemonId=%s", daemonId)
	err := ic.doRequestAndDecode("POST", endpoint, req, &res)
	return &res, err
}

// GetBuildProgress retrieves the build progress for images on the daemon.
// GET /environment/progress?daemonId=...
func (ic *imageClient) GetBuildProgress(daemonId string) (*GetBuildProgressResponse, error) {
	var res GetBuildProgressResponse
	endpoint := fmt.Sprintf("progress?daemonId=%s", daemonId)
	err := ic.doRequestAndDecode("GET", endpoint, nil, &res)
	return &res, err
}
