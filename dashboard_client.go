package mcsmapi

import "net/http"

type dashboardClient MethodClient

func newDashboardClient(client *Client) *dashboardClient {
	return &dashboardClient{client: client, endpoint: "overview"}
}

func (dc *dashboardClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	return dc.client.sendRequest(method, dc.endpoint+"/"+endpoint, body)
}

func (dc *dashboardClient) GetOverview() (*PanelData, error) {
	var overviewData PanelData
	if err := dc.client.doRequestAndDecode("GET", dc.endpoint+"/", nil, &overviewData); err != nil {
		return nil, err
	}

	return &overviewData, nil
}
