package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"

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
