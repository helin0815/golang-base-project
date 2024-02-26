package eventbus

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type DemoEvent struct {
	payload string
}

func (d *DemoEvent) Name() string {
	return "demo"
}

func TestLocalEB(t *testing.T) {
	ctx := context.Background()
	eb := NewEventBus[*DemoEvent]()
	eb.Subscribe(ctx, func(ctx context.Context, payload *DemoEvent) error {
		fmt.Println(payload)
		assert.Equal(t, &DemoEvent{payload: "sss"}, payload)
		return nil
	})
	eb.Subscribe(ctx, func(ctx context.Context, payload *DemoEvent) error {
		fmt.Println(payload)
		assert.Equal(t, &DemoEvent{payload: "sss"}, payload)
		return nil
	})

	eb.Publish(ctx, &DemoEvent{payload: "sss"})
	time.Sleep(time.Second)
}
