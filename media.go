package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type MediaState int

const (
	MediaIdle MediaState = iota
	MediaOpening
	MediaBuffering
	MediaPlaying
	MediaPaused
	MediaStopped
	MediaEnded
	MediaError
)

type Media struct {
	media *C.libvlc_media_t
}

func (m *Media) Release() error {
	if m.media == nil {
		return nil
	}

	C.libvlc_media_release(m.media)
	return getError()
}

func NewMediaFromPath(path string) (*Media, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t = nil
	media = C.libvlc_media_new_path(instance, cPath)
	if media == nil {
		return nil, getError()
	}

	return &Media{media: media}, nil
}

func NewMediaFromUrl(path string) (*Media, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t = nil
	media = C.libvlc_media_new_location(instance, cPath)
	if media == nil {
		return nil, getError()
	}

	return &Media{media: media}, nil
}
