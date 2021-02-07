package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>

extern void eventDispatch(libvlc_event_t*, void*);

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
	return em.attach(event, callback, nil, userData)
}

// attach registers callbacks for an event notification.
func (em *EventManager) attach(event Event, externalCallback EventCallback,
	internalCallback internalEventCallback, userData interface{}) (EventID, error) {
	if err := inst.assertInit(); err != nil {
		return 0, err
	}
	if externalCallback == nil && internalCallback == nil {
		return 0, ErrInvalidEventCallback
	}

	id := inst.events.add(event, externalCallback, internalCallback, userData)
	if C.eventAttach(em.manager, C.libvlc_event_type_t(event), C.ulong(id)) != 0 {
		return 0, getError()
	}

	return id, nil
}

// Detach unregisters the specified event notification.
func (em *EventManager) Detach(eventIDs ...EventID) {
	if err := inst.assertInit(); err != nil {
		return
	}

	for _, eventID := range eventIDs {
		ctx, ok := inst.events.get(eventID)
		if !ok {
			continue
		}

		inst.events.remove(eventID)
		C.eventDetach(em.manager, C.libvlc_event_type_t(ctx.event), C.ulong(eventID))
	}
}

//export eventDispatch
func eventDispatch(event *C.libvlc_event_t, userData unsafe.Pointer) {
	if err := inst.assertInit(); err != nil {
		return
	}

	ctx, ok := inst.events.get(EventID(uintptr(userData)))
	if !ok {
		return
	}

	// Execute external callback.
	if ctx.externalCallback != nil {
		ctx.externalCallback(ctx.event, ctx.userData)
	}

	// Execute internal callback.
	if ctx.internalCallback != nil {
		ctx.internalCallback(event, ctx.userData)
	}
}
