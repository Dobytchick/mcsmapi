package mcsmapi

import (
	"encoding/json"
	"net/http"
)

type dashboardClient MethodClient

func newDashboardClient(client *Client) *dashboardClient {
	return &dashboardClient{client: client, endpoint: "overview"}
}

func (dc *dashboardClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	return dc.client.sendRequest(method, dc.endpoint+"/"+endpoint, body)
}

func (dc *dashboardClient) GetOverview() (*PanelData, error) {
	var overviewData PanelData

	resp, err := dc.sendRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&overviewData)
	if err != nil {
		return nil, err
	}

	return &overviewData, nil
}
