package vlc

//#cgo LDFLAGS: -lvlc
//#include <stdlib.h>
//#include <stdio.h>
//#include <vlc/vlc.h>
//typedef const struct libvlc_event_t clibvlc_event_t;
import "C"
import "unsafe"

//export goCallback
func goCallback(e *C.clibvlc_event_t, userData unsafe.Pointer) {
	token := *(*int)(userData)

	handleLock.Lock()
	ectx := handleMap[token]
	handleLock.Unlock()

	event := &Event{
		Type:   EventType(e._type),
		target: nil, // TODO
	}
	// event.desc.Write(e.u[:])

	ectx.call(event, ectx.userData)
}
