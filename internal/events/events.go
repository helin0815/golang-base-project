package events

import "gitlabee.chehejia.com/k8s/liks-gitops/pkg/eventbus"

func newBus[T eventbus.Event]() eventbus.EventBus[T] {
	return eventbus.NewEventBus[T]()
}

var (
	EBEventDemo = newBus[*EventDemo]()
)
