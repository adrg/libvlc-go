package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"
import "sync"

// EventID uniquely identifies a registered event.
type EventID uint64

// EventCallback represents an event notification callback function.
type EventCallback func(Event, interface{})

type internalEventCallback func(*C.libvlc_event_t, interface{})

type eventContext struct {
	event            Event
	externalCallback EventCallback
	internalCallback internalEventCallback
	userData         interface{}
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
	if id == 0 {
		return nil, false
	}

	er.RLock()
	ctx, ok := er.contexts[id]
	er.RUnlock()

	return ctx, ok
}

func (er *eventRegistry) add(event Event, externalCallback EventCallback,
	internalCallback internalEventCallback, userData interface{}) EventID {
	er.Lock()

	er.sequence++
	id := er.sequence

	er.contexts[id] = &eventContext{
		event:            event,
		externalCallback: externalCallback,
		internalCallback: internalCallback,
		userData:         userData,
	}

	er.Unlock()
	return id
}

func (er *eventRegistry) remove(id EventID) {
	if id == 0 {
		return
	}

	er.Lock()
	delete(er.contexts, id)
	er.Unlock()
}
