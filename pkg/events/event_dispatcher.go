package events

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("handler already registered")
	ErrHandlerNotFound          = errors.New("handler not found")
)

type EventDispatcher struct {
	handlers map[EventName][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[EventName][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName EventName, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[EventName][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName EventName, handler EventHandlerInterface) bool {
	handlers, ok := ed.handlers[eventName]
	if !ok {
		return false
	}
	for _, h := range handlers {
		if h == handler {
			return true
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	handlers, ok := ed.handlers[event.GetName()]
	if !ok {
		return ErrHandlerNotFound
	}
	for _, handler := range handlers {
		handler.Handle(event)
	}
	return nil
}
