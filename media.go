package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// MediaState represents the state of a media file.
type MediaState uint

// Media states.
const (
	MediaNothingSpecial MediaState = iota
	MediaOpening
	MediaBuffering
	MediaPlaying
	MediaPaused
	MediaStopped
	MediaEnded
	MediaError
)

// Media is an abstract representation of a playable media file.
type Media struct {
	media *C.libvlc_media_t
}

// NewMediaFromPath creates a Media instance from the provided path.
func NewMediaFromPath(path string) (*Media, error) {
	return newMedia(path, true)
}

// NewMediaFromURL creates a Media instance from the provided URL.
func NewMediaFromURL(url string) (*Media, error) {
	return newMedia(url, false)
}

// Release destroys the media instance.
func (m *Media) Release() error {
	if m.media == nil {
		return nil
	}

	C.libvlc_media_release(m.media)
	m.media = nil

	return getError()
}

func newMedia(path string, local bool) (*Media, error) {
	if inst == nil {
		return nil, ErrModuleNotInitialized
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t
	if local {
		media = C.libvlc_media_new_path(inst.handle, cPath)
	} else {
		media = C.libvlc_media_new_location(inst.handle, cPath)
	}

	if media == nil {
		return nil, getError()
	}

	return &Media{media: media}, nil
}
