package mcsmapi

import (
	"encoding/json"
	"net/http"
)

// userClient implements methods for working with users via the "auth" endpoint.
type userClient MethodClient

func newUserClient(client *Client) *userClient {
	return &userClient{client: client, endpoint: "auth"}
}

func (uc *userClient) sendRequest(method, endpoint string, body any) (*http.Response, error) {
	fullEndpoint := uc.endpoint
	if endpoint != "" {
		fullEndpoint += "/" + endpoint
	}
	return uc.client.sendRequest(method, fullEndpoint, body)
}

func (uc *userClient) doRequestAndDecode(method, endpoint string, body, out any) error {
	resp, err := uc.sendRequest(method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(out)
}

// GetList retrieves a list of users based on search parameters.
// GET /auth/search?...
func (uc *userClient) GetList(params *UserQueryParams) (*UserListResponse, error) {
	var res UserListResponse
	err := uc.doRequestAndDecode("GET", "search?"+params.BuildQueryString(), nil, &res)
	return &res, err
}

// CreateUser creates a new user.
// POST /auth
func (uc *userClient) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	var res CreateUserResponse
	err := uc.doRequestAndDecode("POST", "", req, &res)
	return &res, err
}

// UpdateUser updates user data by uuid.
// PUT /auth
type UpdateUserRequest struct {
	UUID   string   `json:"uuid"`
	Config UserData `json:"config"`
}

func (uc *userClient) UpdateUser(req *UpdateUserRequest) (*BaseSuccessResponse, error) {
	var res BaseSuccessResponse
	err := uc.doRequestAndDecode("PUT", "", req, &res)
	return &res, err
}

// DeleteUsers deletes users by a list of uuids.
// DELETE /auth
func (uc *userClient) DeleteUsers(uuids []string) (*BaseSuccessResponse, error) {
	var res BaseSuccessResponse
	err := uc.doRequestAndDecode("DELETE", "", uuids, &res)
	return &res, err
}
