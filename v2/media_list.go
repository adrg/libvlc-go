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
		return nil, errOrDefault(getError(), ErrMediaListCreate)
	}

	return &MediaList{list: list}, nil
}

// Release destroys the media list instance.
func (ml *MediaList) Release() error {
	if err := ml.assertInit(); err != nil {
		return nil
	}

	C.libvlc_media_list_release(ml.list)
	ml.list = nil

	return getError()
}

// AddMedia adds the provided Media instance at the end of the media list.
func (ml *MediaList) AddMedia(m *Media) error {
	if err := m.assertInit(); err != nil {
		return err
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
	if err := m.assertInit(); err != nil {
		return err
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

// IndexOfMedia returns the index of the specified media item in the list.
// NOTE: The same instance of a media item can be present multiple times
// in the list. The method returns the first matched index.
func (ml *MediaList) IndexOfMedia(m *Media) (int, error) {
	if err := m.assertInit(); err != nil {
		return 0, err
	}
	if err := ml.Lock(); err != nil {
		return 0, err
	}
	defer ml.Unlock()

	idx := int(C.libvlc_media_list_index_of_item(ml.list, m.media))
	if idx < 0 {
		return 0, errOrDefault(getError(), ErrMediaNotFound)
	}

	return idx, nil
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
	if err := ml.assertInit(); err != nil {
		return false, err
	}

	return (C.libvlc_media_list_is_readonly(ml.list) != C.int(0)), getError()
}

// AssociatedMedia returns the media instance associated with the list,
// if one exists. A media instance is automatically associated with the
// list of its sub-items.
// NOTE: Do not call Release on the returned media instance.
func (ml *MediaList) AssociatedMedia() (*Media, error) {
	if err := ml.assertInit(); err != nil {
		return nil, err
	}

	media := C.libvlc_media_list_media(ml.list)
	if media == nil {
		return nil, errOrDefault(getError(), ErrMediaNotFound)
	}

	// This call will not release the media. Instead, it will decrement
	// the reference count increased by libvlc_media_list_media.
	C.libvlc_media_release(media)

	return &Media{media: media}, nil
}

// AssociateMedia associates the specified media with the media list instance.
// NOTE: If another media instance is already associated with the list, it
// will be released.
func (ml *MediaList) AssociateMedia(m *Media) error {
	if err := ml.assertInit(); err != nil {
		return err
	}
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_set_media(ml.list, m.media)
	return nil
}

// Lock makes the caller the current owner of the media list.
func (ml *MediaList) Lock() error {
	if err := ml.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_lock(ml.list)
	return getError()
}

// Unlock releases ownership of the media list.
func (ml *MediaList) Unlock() error {
	if err := ml.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_unlock(ml.list)
	return getError()
}

// EventManager returns the event manager responsible for the media list.
func (ml *MediaList) EventManager() (*EventManager, error) {
	if err := ml.assertInit(); err != nil {
		return nil, err
	}

	manager := C.libvlc_media_list_event_manager(ml.list)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}

func (ml *MediaList) assertInit() error {
	if ml == nil || ml.list == nil {
		return ErrMediaListNotInitialized
	}

	return nil
}
