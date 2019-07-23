package vlc

import "sync"

// EventID uniquely identifies a registered event.
type EventID uint64

// EventCallback represents an event notification callback function.
type EventCallback func(Event, interface{})

type eventContext struct {
	event    Event
	callback EventCallback
	userData interface{}
}

type eventRegistry struct {
	sync.RWMutex

	contexts map[EventID]*eventContext
	sequence EventID
}

func newEventRegistry() *eventRegistry {
	return &eventRegistry{
		contexts: map[EventID]*eventContext{},
	}
}

func (er *eventRegistry) get(id EventID) (*eventContext, bool) {
	ctx, ok := er.contexts[id]
	return ctx, ok
}

func (er *eventRegistry) add(event Event, callback EventCallback, userData interface{}) EventID {
	er.Lock()
	defer er.Unlock()

	er.sequence++
	er.contexts[er.sequence] = &eventContext{
		event:    event,
		callback: callback,
		userData: userData,
	}

	return er.sequence
}

func (er *eventRegistry) remove(id EventID) {
	er.Lock()
	delete(er.contexts, id)
	er.Unlock()
}
