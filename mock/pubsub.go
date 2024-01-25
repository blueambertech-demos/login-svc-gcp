package mock

import (
	"context"
	"time"
)

type PubSubHandler struct {
}

func (pb *PubSubHandler) Subscribe(_ context.Context, _ string, _ time.Duration, _ func(c context.Context, msgData []byte)) error {
	return nil
}

func (pb *PubSubHandler) Push(_ context.Context, _, _ string) error {
	return nil
}
