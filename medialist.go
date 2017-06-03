package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"
import "errors"

type MediaList struct {
	list *C.libvlc_media_list_t
}

func NewMediaList() (*MediaList, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	var list *C.libvlc_media_list_t = nil
	list = C.libvlc_media_list_new(instance)
	if list == nil {
		return nil, getError()
	}

	return &MediaList{list: list}, nil
}

func (l *MediaList) Release() error {
	if l.list == nil {
		return nil
	}

	C.libvlc_media_list_release(l.list)
	return getError()
}

func (l *MediaList) Retain() error {
	if l.list == nil {
		return nil
	}

	C.libvlc_media_list_retain(l.list)
	return getError()
}

// AddMedia adds a Media instance to the list.
// Lock() the media while performing this operation
func (l *MediaList) AddMedia(m *Media) error {
	if l.list == nil || m.media == nil {
		return errors.New("Media must first be initialized")
	}

	C.libvlc_media_list_add_media(l.list, m.media)
	return getError()
}
