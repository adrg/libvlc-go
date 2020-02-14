package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type instance struct {
	handle *C.libvlc_instance_t
	events *eventRegistry
}

var inst *instance

// Init creates an instance of the VLC module.
// Must be called only once and the module instance must be released using
// the Release function.
func Init(args ...string) error {
	if inst != nil {
		return nil
	}

	argc := len(args)
	argv := make([]*C.char, argc)

	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	defer func() {
		for i := range argv {
			C.free(unsafe.Pointer(argv[i]))
		}
	}()

	handle := C.libvlc_new(C.int(argc), *(***C.char)(unsafe.Pointer(&argv)))
	if handle == nil {
		return errOrDefault(getError(), ErrModuleNotInitialized)
	}

	inst = &instance{
		handle: handle,
		events: newEventRegistry(),
	}

	return nil
}

// Release destroys the instance created by the Init function.
func Release() error {
	if inst == nil {
		return nil
	}

	C.libvlc_release(inst.handle)
	inst = nil

	return getError()
}
