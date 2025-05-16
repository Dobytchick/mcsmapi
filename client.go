package mcsmapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	token   string

	Dashboard *dashboardClient
	Daemon    *daemonClient
	Instance  *instanceClient
	File      *fileClient
	User      *userClient
	Image     *imageClient

	httpClient *http.Client
}

const HTTPTimeout = 10

func NewClient(token string, baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: HTTPTimeout * time.Second,
		}
	}

	client := &Client{
		token:      token,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	client.Daemon = newDaemonClient(client)
	client.Dashboard = newDashboardClient(client)
	client.File = newFileClient(client)
	client.Instance = newInstanceClient(client)
	client.User = newUserClient(client)
	client.Image = newImageClient(client)

	return client
}

func (c *Client) createRequest(endpoint string, body any, method string) (*http.Request, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	sep := "?"
	if strings.Contains(endpoint, "?") {
		sep = "&"
	}
	apiQuery := sep + "apikey=" + c.token

	fmt.Println(c.baseURL + "/api/" + endpoint + apiQuery)

	req, err := http.NewRequest(method, c.baseURL+"/api/"+endpoint+apiQuery, bytes.NewBuffer(bodyBytes))

	return req, err
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	return c.httpClient.Do(req)
}

func (c *Client) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	req, err := c.createRequest(endpoint, body, method)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req)
}

func (c *Client) doRequestAndDecode(method, endpoint string, body, out any) error {
	resp, err := c.sendRequest(method, endpoint, body)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response failed: %w", err)
	}

	return nil
}
