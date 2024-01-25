package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blueambertech-demos/login-svc-gcp/data"
	"github.com/blueambertech-demos/login-svc-gcp/mock"
	"github.com/blueambertech/logging"
)

var testContext context.Context = context.Background()

func TestMain(m *testing.M) {
	logging.Setup(testContext, data.ServiceName)
	defer logging.DeferredCleanup(testContext)
	DbClient = mock.NewNoSQLClient()
	Secrets = &mock.SecretManager{}
	Events = &mock.PubSubHandler{}
	m.Run()
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil).WithContext(testContext)
	w := httptest.NewRecorder()
	healthHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
}

func TestAddLoginHandler(t *testing.T) {
	body, err := getTestPostBody("test@test.com", "somepass")
	if err != nil {
		t.Error(err)
		return
	}
	req := httptest.NewRequest("POST", "/login/add", bytes.NewReader(body)).WithContext(testContext)
	w := httptest.NewRecorder()
	addLoginHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
}

func TestAddLoginHandlerWrongMethod(t *testing.T) {
	body, err := getTestPostBody("test@test.com", "somepass")
	if err != nil {
		t.Error(err)
		return
	}
	req := httptest.NewRequest("GET", "/login/add", bytes.NewReader(body)).WithContext(testContext)
	w := httptest.NewRecorder()
	addLoginHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
}

func TestAddLoginHandlerBadData(t *testing.T) {
	body, err := getTestPostBody("", "")
	if err != nil {
		t.Error(err)
		return
	}
	req := httptest.NewRequest("POST", "/login/add", bytes.NewReader(body)).WithContext(testContext)
	w := httptest.NewRecorder()
	addLoginHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
}

func getTestPostBody(un, pw string) ([]byte, error) {
	details := LoginFormDetails{
		Username: un,
		Password: pw,
	}
	return json.Marshal(details)
}
