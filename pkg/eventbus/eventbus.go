package eventbus

func NewEventBus[T Event]() EventBus[T] {
	return NewLocalEventBus[T]()
}
