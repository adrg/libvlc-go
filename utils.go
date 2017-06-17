package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "errors"

func getError() error {
	msg := C.libvlc_errmsg()
	if msg != nil {
		return errors.New(C.GoString(msg))
	}

	return nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
