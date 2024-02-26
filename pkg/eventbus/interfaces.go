package eventbus

import "context"

type Event interface {
	Name() string
}

type ctxEvent[T Event] struct {
	ctx context.Context

	event T
}

type EventHandler[T Event] func(ctx context.Context, event T) error

type EventBus[T Event] interface {
	Publish(ctx context.Context, event T)
	Subscribe(ctx context.Context, handler EventHandler[T])
}
