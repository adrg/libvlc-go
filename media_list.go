package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"

// MediaList represents a collection of media files.
type MediaList struct {
	list *C.libvlc_media_list_t
}

// NewMediaList creates an empty media list.
func NewMediaList() (*MediaList, error) {
	if inst == nil {
		return nil, ErrModuleNotInitialized
	}

	var list *C.libvlc_media_list_t
	if list = C.libvlc_media_list_new(inst.handle); list == nil {
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
		return ErrMediaListNotInitialized
	}
	if m == nil || m.media == nil {
		return ErrMediaNotInitialized
	}

	C.libvlc_media_list_add_media(ml.list, m.media)
	return getError()
}

// AddMediaFromPath loads a media file from path and adds it
// to the the media list.
func (ml *MediaList) AddMediaFromPath(path string) error {
	media, err := NewMediaFromPath(path)
	if err != nil {
		return err
	}

	return ml.AddMedia(media)
}

// AddMediaFromURL loads a media file from url and adds it
// to the the media list.
func (ml *MediaList) AddMediaFromURL(url string) error {
	media, err := NewMediaFromURL(url)
	if err != nil {
		return err
	}

	return ml.AddMedia(media)
}

// Lock makes the caller the current owner of the media list.
func (ml *MediaList) Lock() error {
	if ml.list == nil {
		return ErrMediaListNotInitialized
	}

	C.libvlc_media_list_lock(ml.list)
	return getError()
}

// Unlock releases ownership of the media list.
func (ml *MediaList) Unlock() error {
	if ml.list == nil {
		return ErrMediaListNotInitialized
	}

	C.libvlc_media_list_unlock(ml.list)
	return getError()
}

// EventManager returns the event manager responsible for the media list.
func (ml *MediaList) EventManager() (*EventManager, error) {
	if ml.list == nil {
		return nil, ErrMediaListNotInitialized
	}

	manager := C.libvlc_media_list_event_manager(ml.list)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}
