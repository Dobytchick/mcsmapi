package mcsmapi

import (
	"encoding/json"
	"net/http"
)

// userClient реализует методы работы с пользователями через endpoint "auth".
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

// GetList получает список пользователей по параметрам поиска.
// GET /auth/search?...
func (uc *userClient) GetList(params *UserQueryParams) (*UserListResponse, error) {
	var res UserListResponse
	err := uc.doRequestAndDecode("GET", "search?"+params.BuildQueryString(), nil, &res)
	return &res, err
}

// CreateUser создает нового пользователя.
// POST /auth
func (uc *userClient) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	var res CreateUserResponse
	err := uc.doRequestAndDecode("POST", "", req, &res)
	return &res, err
}

// UpdateUser обновляет данные пользователя по uuid.
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

// DeleteUsers удаляет пользователей по списку uuid.
// DELETE /auth
func (uc *userClient) DeleteUsers(uuids []string) (*BaseSuccessResponse, error) {
	var res BaseSuccessResponse
	err := uc.doRequestAndDecode("DELETE", "", uuids, &res)
	return &res, err
}
