package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
)

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

func urlToPath(mrl string) (string, error) {
	url, err := url.Parse(mrl)
	if err != nil {
		return "", err
	}
	if url.Scheme != "file" {
		return mrl, nil
	}
	path := filepath.Clean(url.Path)

	if runtime.GOOS == "windows" {
		sep := string(filepath.Separator)
		if url.Host != "" {
			path = strings.Repeat(sep, 2) + filepath.Join(url.Host, path)
		} else {
			path = strings.TrimLeft(path, sep)
		}
	}

	return path, nil
}
