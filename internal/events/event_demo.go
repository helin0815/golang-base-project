package events

type EventDemo struct {
}

func (e *EventDemo) Name() string {
	return "this is a demo event"
}
