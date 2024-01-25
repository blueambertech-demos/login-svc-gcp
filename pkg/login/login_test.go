package login

import (
	"context"
	"testing"
	"time"

	"github.com/blueambertech-demos/login-svc-gcp/mock"
)

var fakeDbClient *mock.NoSQLClient
var fakeEventQueue *mock.PubSubHandler

func TestMain(m *testing.M) {
	fakeDbClient = &mock.NoSQLClient{}
	fakeDbClient.SetData(make(map[string]map[string]interface{}))
	fakeEventQueue = &mock.PubSubHandler{}
	m.Run()
}

func TestAddLogin(t *testing.T) {
	defer fakeDbClient.ClearData()
	ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	defer canc()
	err := AddLogin(ctx, fakeDbClient, fakeEventQueue, "hello@test.com", "password", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestAddLoginDuplicate(t *testing.T) {
	defer fakeDbClient.ClearData()
	ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	defer canc()
	err := AddLogin(ctx, fakeDbClient, fakeEventQueue, "hello@test.com", "password", nil)
	if err != nil {
		t.Error(err)
	}
	err = AddLogin(ctx, fakeDbClient, fakeEventQueue, "hello@test.com", "password", nil)
	if err == nil {
		t.Error("Duplicate user should be rejected")
	}
}

func TestVerifyCredentials(t *testing.T) {
	defer fakeDbClient.ClearData()
	ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	defer canc()
	err := AddLogin(ctx, fakeDbClient, fakeEventQueue, "hello@test.com", "password", nil)
	if err != nil {
		t.Error(err)
	}

	result, _, err := VerifyCredentials(ctx, fakeDbClient, "hello@test.com", "password")
	if err != nil {
		t.Error(err)
		return
	}
	if !result {
		t.Errorf("Result was false")
	}
}

func TestVerifyCredentialsWrongPass(t *testing.T) {
	defer fakeDbClient.ClearData()
	ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	defer canc()
	err := AddLogin(ctx, fakeDbClient, fakeEventQueue, "hello@test.com", "password", nil)
	if err != nil {
		t.Error(err)
	}

	result, _, err := VerifyCredentials(ctx, fakeDbClient, "hello@test.com", "passwfgdford")
	if err != nil {
		t.Error(err)
		return
	}
	if result {
		t.Errorf("Result should be false")
	}
}

func TestVerifyCredentialsMissingUser(t *testing.T) {
	defer fakeDbClient.ClearData()
	ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	defer canc()

	result, _, err := VerifyCredentials(ctx, fakeDbClient, "hello@test.com", "password")
	if err == nil {
		t.Error("Error was nil")
		return
	}
	if result == true {
		t.Errorf("Result should be false")
	}
}

func TestValidateCredentials(t *testing.T) {
	defer fakeDbClient.ClearData()
	result := ValidateCredentials("valid@valid.com", "validpassword", nil)
	if !result {
		t.Error("incorrect validation of valid details")
		return
	}
	result = ValidateCredentials("invaliduser", "validpassword", nil)
	if result {
		t.Error("incorrect validation of invalid username")
		return
	}
	result = ValidateCredentials("valid@valid.com", "", nil)
	if result {
		t.Error("incorrect validation of invalid password")
		return
	}
}
