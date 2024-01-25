package mock

import "context"

type SecretManager struct {
	data map[string]interface{}
}

func NewSecretManager() *SecretManager {
	sm := &SecretManager{
		data: map[string]interface{}{},
	}
	sm.data["jwt-auth-token-key"] = "somekey"
	return sm
}

func (sm *SecretManager) Get(_ context.Context, key string) (interface{}, error) {
	return sm.data[key], nil
}
