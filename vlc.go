package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <vlc/libvlc_version.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type instance struct {
	handle *C.libvlc_instance_t
	events *eventRegistry
}

func (i *instance) assertInit() error {
	if i == nil || i.handle == nil {
		return ErrModuleNotInitialized
	}

	return nil
}

var inst *instance

// VersionInfo contains details regarding the version of the libVLC module.
type VersionInfo struct {
	Major uint
	Minor uint
	Patch uint
}

// String returns a string representation of the version.
func (v VersionInfo) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

var moduleVersion = VersionInfo{
	Major: C.LIBVLC_VERSION_MAJOR,
	Minor: C.LIBVLC_VERSION_MINOR,
	Patch: C.LIBVLC_VERSION_REVISION,
}

// Init creates an instance of the libVLC module.
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
		return errOrDefault(getError(), ErrModuleInitialize)
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

// Version returns details regarding the version of the libVLC module.
func Version() VersionInfo {
	return moduleVersion
}
