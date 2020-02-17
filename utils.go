package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "errors"

func getError() error {
	msg := C.libvlc_errmsg()
	if msg == nil {
		return nil
	}

	err := errors.New(C.GoString(msg))
	C.libvlc_clearerr()
	return err
}

func errOrDefault(err, defaultErr error) error {
	if err != nil {
		return err
	}

	return defaultErr
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
