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
	if err := inst.assertInit(); err != nil {
		return nil, err
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

// AddMedia adds the provided Media instance at the end of the media list.
func (ml *MediaList) AddMedia(m *Media) error {
	if m == nil || m.media == nil {
		return ErrMediaNotInitialized
	}
	if err := ml.Lock(); err != nil {
		return err
	}
	defer ml.Unlock()

	C.libvlc_media_list_add_media(ml.list, m.media)
	return getError()
}

// AddMediaFromPath loads the media file at the specified path and adds it at
// the end of the media list.
func (ml *MediaList) AddMediaFromPath(path string) error {
	media, err := NewMediaFromPath(path)
	if err != nil {
		return err
	}

	return ml.AddMedia(media)
}

// AddMediaFromURL loads the media file at the specified URL and adds it at the
// end of the the media list.
func (ml *MediaList) AddMediaFromURL(url string) error {
	media, err := NewMediaFromURL(url)
	if err != nil {
		return err
	}

	return ml.AddMedia(media)
}

// InsertMedia inserts the provided Media instance in the list,
// at the specified index.
func (ml *MediaList) InsertMedia(m *Media, index uint) error {
	if m == nil || m.media == nil {
		return ErrMediaNotInitialized
	}
	if err := ml.Lock(); err != nil {
		return err
	}
	defer ml.Unlock()

	C.libvlc_media_list_insert_media(ml.list, m.media, C.int(index))
	return getError()
}

// InsertMediaFromPath loads the media file at the provided path and inserts
// it in the list, at the specified index.
func (ml *MediaList) InsertMediaFromPath(path string, index uint) error {
	media, err := NewMediaFromPath(path)
	if err != nil {
		return err
	}

	return ml.InsertMedia(media, index)
}

// InsertMediaFromURL loads the media file at the provided URL and inserts
// it in the list, at the specified index.
func (ml *MediaList) InsertMediaFromURL(url string, index uint) error {
	media, err := NewMediaFromURL(url)
	if err != nil {
		return err
	}

	return ml.InsertMedia(media, index)
}

// RemoveMediaAtIndex removes the media item at the specified index
// from the list.
func (ml *MediaList) RemoveMediaAtIndex(index uint) error {
	if err := ml.Lock(); err != nil {
		return err
	}
	defer ml.Unlock()

	C.libvlc_media_list_remove_index(ml.list, C.int(index))
	return getError()
}

// MediaAtIndex returns the media item at the specified index from the list.
func (ml *MediaList) MediaAtIndex(index uint) (*Media, error) {
	if err := ml.Lock(); err != nil {
		return nil, err
	}
	defer ml.Unlock()

	media := C.libvlc_media_list_item_at_index(ml.list, C.int(index))
	if media == nil {
		return nil, getError()
	}

	// This call will not release the media. Instead, it will decrement
	// the reference count increased by libvlc_media_list_item_at_index.
	C.libvlc_media_release(media)

	return &Media{media}, nil
}

// Count returns the number of media items in the list.
func (ml *MediaList) Count() (int, error) {
	if err := ml.Lock(); err != nil {
		return 0, err
	}
	defer ml.Unlock()

	return int(C.libvlc_media_list_count(ml.list)), getError()
}

// IsReadOnly specifies if the media list can be modified.
func (ml *MediaList) IsReadOnly() (bool, error) {
	if ml.list == nil {
		return false, ErrMediaListNotInitialized
	}

	return (C.libvlc_media_list_is_readonly(ml.list) != C.int(0)), getError()
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
