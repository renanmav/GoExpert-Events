package events

import "time"

type EventName string

type EventInterface interface {
	GetName() EventName
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName EventName, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName EventName, handler EventHandlerInterface) error
	Has(eventName EventName, handler EventHandlerInterface) bool
	Clear() error
}
