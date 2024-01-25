package mock

import "context"

type SecretManager struct {
	data map[string]interface{}
}

func (sm *SecretManager) Get(_ context.Context, key string) (interface{}, error) {
	return sm.data[key], nil
}
