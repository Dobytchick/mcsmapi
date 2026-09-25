package mcsmapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

type requestError struct{ cause error }

func (e *requestError) Error() string { return "HTTP request failed" }
func (e *requestError) Unwrap() error { return e.cause }

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
	u, err := url.Parse(strings.TrimRight(c.baseURL, "/") + "/api/" + strings.TrimLeft(endpoint, "/"))
	if err != nil {
		return nil, fmt.Errorf("invalid request URL: %w", err)
	}
	q := u.Query()
	q.Set("apikey", c.token)
	u.RawQuery = q.Encode()
	var payload *bytes.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
		}
		payload = bytes.NewReader(bodyBytes)
	} else {
		payload = bytes.NewReader(nil)
	}
	return http.NewRequest(method, u.String(), payload)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &requestError{cause: err}
	}
	return resp, nil
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP status %d", resp.StatusCode)
	}
	var envelope struct {
		Status int `json:"status"`
	}
	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return fmt.Errorf("decode response failed: %w", err)
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return fmt.Errorf("decode response failed: %w", err)
	}
	if envelope.Status != http.StatusOK {
		return fmt.Errorf("API status %d", envelope.Status)
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode response failed: %w", err)
	}
	return nil
}
