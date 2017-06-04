package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"
import "errors"

type MediaList struct {
	list *C.libvlc_media_list_t
}

// NewMediaList creates an empty media list.
func NewMediaList() (*MediaList, error) {
	if instance == nil {
		return nil, errors.New("Module must be initialized first")
	}

	var list *C.libvlc_media_list_t
	if list = C.libvlc_media_list_new(instance); list == nil {
		return nil, getError()
	}

	return &MediaList{list: list}, nil
}

// Release destroys the media list instance.
func (ml *MediaList) Release() error {
	if ml.list == nil {
		return nil
	}

	C.libvlc_media_list_release(ml.list)
	ml.list = nil

	return getError()
}

// AddMedia adds a Media instance to the media list.
func (ml *MediaList) AddMedia(m *Media) error {
	if ml.list == nil {
		return errors.New("Media list must be initialized first")
	}
	if m.media == nil {
		return errors.New("Media must be initialized first")
	}

	C.libvlc_media_list_add_media(ml.list, m.media)
	return getError()
}
