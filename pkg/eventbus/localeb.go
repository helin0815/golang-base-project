package eventbus

import (
	"context"
	"sync"
	"time"

	"gitlabee.chehejia.com/gopkg/lsego/pkg/log"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/otelutils"
)

type LocalEventBus[T Event] struct {
	sync.RWMutex

	busChan     chan ctxEvent[T]
	subHandlers []EventHandler[T]
}

func NewLocalEventBus[T Event]() EventBus[T] {
	eb := &LocalEventBus[T]{
		busChan: make(chan ctxEvent[T]),
	}

	go eb.startListening(context.Background())
	return eb
}

func (l *LocalEventBus[T]) startListening(ctx context.Context) {
	log.Infof("LocalEventBus started")
	for {
		select {
		case ie := <-l.busChan:
			l.RLock()
			traceId := otelutils.TraceIdFromContext(ie.ctx)
			otelutils.SpanTrackEvent(ie.ctx, "IEB Handle", map[string]string{"eventName": ie.event.Name()})
			for idx, handler := range l.subHandlers {
				log.Infof("[%s] event handler: %s-%d", traceId, ie.event.Name(), idx)

				h := handler
				go func() {
					if err := h(ie.ctx, ie.event); err != nil {
						log.Errorf("[%s] event %s handle failed: %s", traceId, ie.event.Name(), err)
					}
				}()
			}
			l.RUnlock()
		case <-ctx.Done():
			return
		}
	}
}
func (l *LocalEventBus[T]) Publish(ctx context.Context, event T) {
	l.busChan <- ctxEvent[T]{ctx: WithoutCancel(ctx), event: event}
}

func (l *LocalEventBus[T]) Subscribe(ctx context.Context, handler EventHandler[T]) {
	l.RLock()
	defer l.RUnlock()

	l.subHandlers = append(l.subHandlers, handler)
}

// TODO: 因为go1.20还不支持 context.WithoutCancel，所以在这里自行实现一个，等升级1.21之后可以把它换掉
type noCancel struct {
	ctx context.Context
}

func (c noCancel) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c noCancel) Done() <-chan struct{}             { return nil }
func (c noCancel) Err() error                        { return nil }
func (c noCancel) Value(key interface{}) interface{} { return c.ctx.Value(key) }

// WithoutCancel returns a context that is never canceled.
func WithoutCancel(ctx context.Context) context.Context {
	return noCancel{ctx: ctx}
}
