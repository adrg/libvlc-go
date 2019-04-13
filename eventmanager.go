package vlc

//#cgo LDFLAGS: -lvlc
//#include <stdlib.h>
//#include <stdio.h>
//#include <vlc/vlc.h>
//typedef const struct libvlc_event_t *clibvlc_event_t;
//extern void goCallback(clibvlc_event_t, void*);
//static inline int goAttach(libvlc_event_manager_t* em, libvlc_event_type_t et, unsigned long userData) {
//    return libvlc_event_attach(em, et, goCallback, (void *) userData);
//}
//static inline int goDetach(libvlc_event_manager_t* em, libvlc_event_type_t et, unsigned long userData) {
//    libvlc_event_detach(em, et, goCallback, (void *) userData);
//}
import "C"

import (
	"reflect"
	"sync"
	"unsafe"
)

// EventCallback is the type of an event callback function.
type EventCallback func(Event, interface{})

type eventContext struct {
	token    uint64
	et       C.libvlc_event_type_t
	call     EventCallback
	userData interface{}
}

var (
	eventRegistry = struct {
		sync.RWMutex
		nextToken uint64
		contexts  map[uint64]*eventContext
	}{
		contexts: make(map[uint64]*eventContext),
	}
)

// EventManager wraps a libvlc event manager.
type EventManager struct {
	Manager *C.libvlc_event_manager_t
}

// NewEventManager returns a new EventManager instance wrapping a libvlc manager.
func NewEventManager(manager *C.libvlc_event_manager_t) *EventManager {
	return &EventManager{
		Manager: manager,
	}
}

// Attach registers a callback for an event notification.
func (em *EventManager) Attach(eventType Event, callback EventCallback, userData interface{}) (uint64, error) {
	ectx := &eventContext{
		et:       C.libvlc_event_type_t(eventType),
		call:     callback,
		userData: userData,
	}
	eventRegistry.Lock()
	eventRegistry.nextToken++
	ectx.token = eventRegistry.nextToken
	eventRegistry.contexts[eventRegistry.nextToken] = ectx
	eventRegistry.Unlock()

	if C.goAttach(em.Manager, ectx.et, (C.ulong)(ectx.token)) != 0 {
		return 0, getError()
	}

	return ectx.token, nil
}

// Detach unregisters an event notification.
func (em *EventManager) Detach(eventType Event, callback EventCallback, userData interface{}) {
	var ectx *eventContext

	same := func(ctx *eventContext) bool {
		return (ctx.et == C.libvlc_event_type_t(eventType) &&
			reflect.DeepEqual(ctx.call, callback) && reflect.DeepEqual(ctx.userData, userData))
	}

	eventRegistry.Lock()
	for token := range eventRegistry.contexts {
		c := eventRegistry.contexts[token]
		if same(c) {
			// we have the listener
			ectx = c
			break
		}
	}
	delete(eventRegistry.contexts, ectx.token)
	eventRegistry.Unlock()

	C.goDetach(em.Manager, ectx.et, (C.ulong)(ectx.token))

}

//export goCallback
func goCallback(event C.clibvlc_event_t, userDataC unsafe.Pointer) {
	userData := uint64(uintptr(userDataC))
	ectx := eventRegistry.contexts[userData]
	ectx.call(Event(event._type), ectx.userData)
}
