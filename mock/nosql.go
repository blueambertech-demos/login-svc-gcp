package mock

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/blueambertech-demos/login-svc-gcp/data"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type NoSQLClient struct {
	data map[string]map[string]interface{}
}

func NewNoSQLClient() *NoSQLClient {
	return &NoSQLClient{
		data: map[string]map[string]interface{}{},
	}
}

func (f *NoSQLClient) Read(_ context.Context, _, id string) (map[string]interface{}, error) {
	d := f.data[id]
	return structs.Map(d), nil
}

func (f *NoSQLClient) Insert(_ context.Context, _ string, data interface{}) (string, error) {
	id := fmt.Sprintf("%d", rand.New(rand.NewSource(535345)).Int())
	f.data[id] = structs.Map(data)
	return id, nil
}

func (f *NoSQLClient) InsertWithID(_ context.Context, _, id string, data interface{}) error {
	f.data[id] = structs.Map(data)
	return nil
}

func (f *NoSQLClient) Where(_ context.Context, _, _, _, val string) (map[string]map[string]interface{}, error) {
	var details = map[string]map[string]interface{}{}
	for i, v := range f.data {
		var d data.LoginDetails
		if err := mapstructure.Decode(v, &d); err != nil {
			return nil, err
		}
		// Assuming key is UserName and op is == for simplicity
		if d.UserName == val {
			details[i] = v
		}
	}
	return details, nil
}

func (f *NoSQLClient) Exists(_ context.Context, _, key string) (bool, error) {
	_, ok := f.data[key]
	if !ok {
		return false, nil
	}
	return true, nil
}

func (f *NoSQLClient) SetData(d map[string]map[string]interface{}) {
	f.data = d
}

func (f *NoSQLClient) ClearData() {
	for k := range f.data {
		delete(f.data, k)
	}
}
