package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>

typedef const struct libvlc_event_t* clibvlc_event_t;
extern void eventDispatch(clibvlc_event_t, void*);

static inline int eventAttach(libvlc_event_manager_t* em, libvlc_event_type_t et, unsigned long userData) {
    return libvlc_event_attach(em, et, eventDispatch, (void*)userData);
}
static inline int eventDetach(libvlc_event_manager_t* em, libvlc_event_type_t et, unsigned long userData) {
    libvlc_event_detach(em, et, eventDispatch, (void*)userData);
}
*/
import "C"

import (
	"unsafe"
)

// EventManager wraps a libvlc event manager.
type EventManager struct {
	manager *C.libvlc_event_manager_t
}

// newEventManager returns a new event manager instance.
func newEventManager(manager *C.libvlc_event_manager_t) *EventManager {
	return &EventManager{
		manager: manager,
	}
}

// Attach registers a callback for an event notification.
func (em *EventManager) Attach(event Event, callback EventCallback, userData interface{}) (EventID, error) {
	if callback == nil {
		return 0, ErrInvalidEventCallback
	}

	id := inst.events.add(event, callback, userData)
	if C.eventAttach(em.manager, C.libvlc_event_type_t(event), C.ulong(id)) != 0 {
		return 0, getError()
	}

	return id, nil
}

// Detach unregisters the specified event notification.
func (em *EventManager) Detach(eventID EventID) {
	ctx, ok := inst.events.get(eventID)
	if !ok {
		return
	}

	inst.events.remove(eventID)
	C.eventDetach(em.manager, C.libvlc_event_type_t(ctx.event), C.ulong(eventID))
}

//export eventDispatch
func eventDispatch(event C.clibvlc_event_t, userData unsafe.Pointer) {
	ctx, ok := inst.events.get(EventID(uintptr(userData)))
	if !ok {
		return
	}
	if ctx.callback == nil {
		return
	}

	ctx.callback(ctx.event, ctx.userData)
}
