package mcsmapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Test BuildQueryString with a struct
func TestBuildQueryString(t *testing.T) {
	type testStruct struct {
		DaemonID string `url:"daemonId"`
		UUID     string `url:"uuid"`
	}
	q := testStruct{DaemonID: "abc", UUID: "123"}
	result := BuildQueryString(q)
	if !strings.Contains(result, "daemonId=abc") || !strings.Contains(result, "uuid=123") {
		t.Errorf("unexpected query string: %s", result)
	}
}

// Test BaseRequest.BuildQueryString method
func TestBaseRequest_BuildQueryString(t *testing.T) {
	req := BaseRequest{DaemonID: "d1", UUID: "u1"}
	result := req.BuildQueryString()
	if !strings.Contains(result, "daemonId=d1") || !strings.Contains(result, "uuid=u1") {
		t.Errorf("unexpected query string: %s", result)
	}
}

// Test BaseSuccessResponse JSON marshaling
func TestBaseSuccessResponse_JSON(t *testing.T) {
	resp := BaseSuccessResponse{Status: 200, Data: true, Time: 123456}
	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"status":200`) || !strings.Contains(s, `"data":true`) || !strings.Contains(s, `"time":123456`) {
		t.Errorf("unexpected json: %s", s)
	}
}

// Test BoolResponse JSON marshaling
func TestBoolResponse_JSON(t *testing.T) {
	resp := BoolResponse{Data: true}
	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"data":true`) {
		t.Errorf("unexpected json: %s", s)
	}
}

// --- Client tests ---

func TestNewClient_DefaultHTTPClient(t *testing.T) {
	client := NewClient("token", "http://localhost", nil)
	if client == nil {
		t.Fatal("expected client, got nil")
	}
	if client.httpClient == nil {
		t.Fatal("expected httpClient to be set")
	}
	if client.Daemon == nil || client.Dashboard == nil || client.File == nil || client.Instance == nil || client.User == nil {
		t.Fatal("expected all subclients to be initialized")
	}
}

func TestNewClient_CustomHTTPClient(t *testing.T) {
	custom := &http.Client{Timeout: 5 * time.Second}
	client := NewClient("token", "http://localhost", custom)
	if client.httpClient != custom {
		t.Error("expected custom http client to be used")
	}
}

func TestClient_createRequest(t *testing.T) {
	client := NewClient("token", "http://localhost", nil)
	req, err := client.createRequest("/test", map[string]string{"foo": "bar"}, http.MethodPost)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.Method != http.MethodPost {
		t.Errorf("expected POST, got %s", req.Method)
	}
	if !strings.Contains(req.URL.String(), "/api/test?apikey=token") {
		t.Errorf("unexpected url: %s", req.URL.String())
	}
	body, _ := io.ReadAll(req.Body)
	if !bytes.Contains(body, []byte("foo")) {
		t.Errorf("expected body to contain foo")
	}
}

func TestClient_createRequest_WithQuery(t *testing.T) {
	client := NewClient("token", "http://localhost", nil)
	req, err := client.createRequest("/test?x=1", map[string]string{"foo": "bar"}, http.MethodPost)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(req.URL.String(), "/api/test?x=1&apikey=token") {
		t.Errorf("unexpected url: %s", req.URL.String())
	}
}

func TestClient_doRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected application/json header")
		}
		if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
			t.Errorf("expected X-Requested-With header")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClient("token", srv.URL, nil)
	req, _ := client.createRequest("/test", nil, http.MethodGet)
	resp, err := client.doRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestClient_sendRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()
	client := NewClient("token", srv.URL, nil)
	resp, err := client.sendRequest(http.MethodPost, "/test", map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
}

func TestClient_doRequestAndDecode(t *testing.T) {
	type outStruct struct {
		Msg string `json:"msg"`
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"msg":"hello"}`))
	}))
	defer srv.Close()
	client := NewClient("token", srv.URL, nil)
	var out outStruct
	err := client.doRequestAndDecode(http.MethodGet, "/test", nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Msg != "hello" {
		t.Errorf("expected msg=hello, got %s", out.Msg)
	}
}

func TestClient_doRequestAndDecode_BadJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`notjson`))
	}))
	defer srv.Close()
	client := NewClient("token", srv.URL, nil)
	var out map[string]any
	err := client.doRequestAndDecode(http.MethodGet, "/test", nil, &out)
	if err == nil || !strings.Contains(err.Error(), "decode response failed") {
		t.Errorf("expected decode error, got %v", err)
	}
}

func TestClient_doRequestAndDecode_RequestError(t *testing.T) {
	client := NewClient("token", "http://127.0.0.1:0", nil) // invalid port
	var out map[string]any
	err := client.doRequestAndDecode(http.MethodGet, "/test", nil, &out)
	if err == nil || !strings.Contains(err.Error(), "request failed") {
		t.Errorf("expected request failed error, got %v", err)
	}
}
