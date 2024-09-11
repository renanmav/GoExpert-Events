package events

import (
	"sync"
	"time"
)

type EventName string

type EventInterface interface {
	GetName() EventName
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName EventName, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName EventName, handler EventHandlerInterface) error
	Has(eventName EventName, handler EventHandlerInterface) bool
	Clear() error
}
