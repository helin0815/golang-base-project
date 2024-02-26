package events

import (
	"context"
	"fmt"
	"testing"
)

func TestEvents(t *testing.T) {
	EBEventDemo.Publish(context.Background(), &EventDemo{})
	EBEventDemo.Subscribe(context.Background(), func(ctx context.Context, event *EventDemo) error {
		fmt.Println(event)
		return nil
	})
}
