package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher creates a new instance of EventDispatcher.
//
// It initializes the handlers map to store event handlers.
//
// Returns a pointer to the newly created EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Register adds an event handler for a specific event name.
// If the handler is already registered for the event, it returns an error.
//
// Parameters:
//
// - eventName: The name of the event to register the handler for.
//
// - handler: The event handler to be registered.
//
// Returns:
//
// - An error if the handler is already registered for the event, otherwise nil.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
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

// Clear removes all registered event handlers.
//
// It reinitializes the handlers map, effectively clearing all registered handlers.
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

// Has checks if a specific event handler is registered for a given event name.
//
// Parameters:
//
// - eventName: The name of the event to check the handler for.
//
// - handler: The event handler to check for.
//
// Returns:
//
// - true if the handler is registered for the event, otherwise false.
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

// Dispatch triggers all handlers registered for a specific event.
//
// It calls the Handle method on each registered handler for the event.
//
// Parameters:
//
// - event: The event to be dispatched to the handlers.
//
// Returns:
//
// - An error if the dispatch process fails, otherwise nil.
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

// Remove removes a specific event handler for a given event name.
//
// If the handler is found, it is removed from the list of handlers for the event.
//
// Parameters:
//
// - eventName: The name of the event to remove the handler from.
//
// - handler: The event handler to be removed.
//
// Returns:
//
// - An error if the handler is not found for the event, otherwise nil.
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
