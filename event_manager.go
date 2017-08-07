package vlc

//#cgo LDFLAGS: -lvlc
//#include <stdlib.h>
//#include <stdio.h>
//#include <vlc/vlc.h>
//typedef const struct libvlc_event_t clibvlc_event_t;
//extern void goCallback(clibvlc_event_t*, void*);
//static inline int goAttach(libvlc_event_manager_t* em, libvlc_event_type_t et, void* userData) {
//    return libvlc_event_attach(em, et, goCallback, userData);
//}
import "C"

import (
	"log"
	"sync"
	"unsafe"
)

type Callback func(*Event, interface{})

type eventContext struct {
	token    int
	et       C.libvlc_event_type_t
	call     Callback
	userData interface{}
}

var (
	handleLock  sync.Mutex
	handleToken int
	handleMap   = make(map[int]*eventContext)
)

type EventManager struct {
	event_manager *C.libvlc_event_manager_t
}

func NewEventManager(event_manager *C.libvlc_event_manager_t) *EventManager {
	return &EventManager{
		event_manager: event_manager,
	}
}

func (em *EventManager) Attach(eventType EventType, callback Callback, userData interface{}) (int, error) {
	ectx := &eventContext{
		et:       C.libvlc_event_type_t(eventType),
		call:     callback,
		userData: userData,
	}
	handleLock.Lock()
	handleToken++
	ectx.token = handleToken
	handleMap[ectx.token] = ectx
	handleLock.Unlock()

	if C.goAttach(em.event_manager, ectx.et, unsafe.Pointer(&handleToken)) != 0 {
		log.Fatal("TODO")
	}

	return ectx.token, nil
}

func (em *EventManager) Dettach(eventType EventType, callback Callback, userData interface{}) (uintptr, error) {
	log.Fatal("TODO")
	return 0, nil
}
