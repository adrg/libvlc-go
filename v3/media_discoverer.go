package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// MediaDiscoveryCallback is used by media discovery services to report
// discovery events. The callback provides the event, the media instance,
// and the index at which the action takes place in the media list of the
// discovery service.
//
// The available events are:
//   - MediaListWillAddItem
//   - MediaListItemAdded
//   - MediaListWillDeleteItem
//   - MediaListItemDeleted
type MediaDiscoveryCallback func(Event, *Media, int)

// MediaDiscoveryCategory defines categories of media discovery services.
type MediaDiscoveryCategory uint

// Media discovery categories.
const (
	// Devices (e.g. portable devices supporting MTP, discs).
	MediaDiscoveryDevices MediaDiscoveryCategory = iota

	// LAN/WAN services (e.g. UPnP, SMB, SAP).
	MediaDiscoveryLAN

	// Internet services (e.g. podcasts, radio stations).
	MediaDiscoveryInternet

	// Local directories.
	MediaDiscoveryLocal
)

// MediaDiscovererDescriptor contains information about a media
// discovery service. Pass the `Name` field to the NewMediaDiscoverer
// method in order to create a new discovery service instance.
type MediaDiscovererDescriptor struct {
	Name     string
	LongName string
	Category MediaDiscoveryCategory
}

// ListMediaDiscoverers returns a list of descriptors identifying the
// available media discovery services of the specified category.
func ListMediaDiscoverers(category MediaDiscoveryCategory) ([]*MediaDiscovererDescriptor, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	// Get media discoverer descriptors.
	var cDescriptors **C.libvlc_media_discoverer_description_t

	count := int(C.libvlc_media_discoverer_list_get(inst.handle, C.libvlc_media_discoverer_category_t(category), &cDescriptors))
	if count <= 0 || cDescriptors == nil {
		return nil, nil
	}
	defer C.libvlc_media_discoverer_list_release(cDescriptors, C.size_t(count))

	// Parse media discoverer descriptors.
	descriptors := make([]*MediaDiscovererDescriptor, 0, count)
	for i := 0; i < count; i++ {
		// Get current media discoverer descriptor.
		cDescriptorPtr := unsafe.Pointer(uintptr(unsafe.Pointer(cDescriptors)) +
			uintptr(i)*unsafe.Sizeof(*cDescriptors))
		if cDescriptorPtr == nil {
			return nil, ErrMediaDiscovererParse
		}

		cDescriptor := *(**C.libvlc_media_discoverer_description_t)(cDescriptorPtr)
		if cDescriptor == nil {
			return nil, ErrMediaDiscovererParse
		}

		// Parse media discoverer descriptor.
		descriptors = append(descriptors, &MediaDiscovererDescriptor{
			Name:     C.GoString(cDescriptor.psz_name),
			LongName: C.GoString(cDescriptor.psz_longname),
			Category: MediaDiscoveryCategory(cDescriptor.i_cat),
		})
	}

	return descriptors, nil
}

// MediaDiscoverer represents a media discovery service.
// Discovery services use different discovery protocols (e.g. MTP, UPnP, SMB)
// in order to find available media instances.
type MediaDiscoverer struct {
	discoverer *C.libvlc_media_discoverer_t
	stopFunc   func()
}

// NewMediaDiscoverer instantiates the media discovery service identified
// by the specified name. Use the ListMediaDiscoverers method to obtain the
// list of available discovery service descriptors.
//   NOTE: Call the Release method on the discovery service instance in
//   order to free the allocated resources.
func NewMediaDiscoverer(name string) (*MediaDiscoverer, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	discoverer := C.libvlc_media_discoverer_new(inst.handle, cName)
	if discoverer == nil {
		return nil, errOrDefault(getError(), ErrMediaDiscovererCreate)
	}

	return &MediaDiscoverer{
		discoverer: discoverer,
	}, nil
}

// Release stops and destroys the media discovery service along
// with all the media found by the instance.
func (md *MediaDiscoverer) Release() error {
	if err := md.assertInit(); err != nil {
		return nil
	}

	// Stop discovery service.
	md.stop()

	// Release discovery service.
	C.libvlc_media_discoverer_release(md.discoverer)
	md.discoverer = nil

	return getError()
}

// Start starts the media discovery service and reports discovery
// events through the specified callback function.
//   NOTE: The Stop and Release methods should not be called from the callback
//   function. Doing so will result in undefined behavior.
func (md *MediaDiscoverer) Start(cb MediaDiscoveryCallback) error {
	if cb == nil {
		return ErrInvalidEventCallback
	}

	// Stop discovery service, if started.
	if err := md.Stop(); err != nil {
		return err
	}

	// Retrieve media list.
	ml, err := md.MediaList()
	if err != nil {
		return err
	}

	// Retrieve media list event manager.
	manager, err := ml.EventManager()
	if err != nil {
		return err
	}

	// Create event callback.
	eventCallback := func(event *C.libvlc_event_t, userData interface{}) {
		if err := md.assertInit(); err != nil {
			return
		}
		if event == nil {
			return
		}

		// Parse event media.
		cMedia := *(**C.libvlc_media_t)(unsafe.Pointer(&event.u[0]))
		if cMedia == nil {
			return
		}
		media := &Media{media: cMedia}

		// Parse event media index.
		cIndex := (*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(&event.u[0])) +
			unsafe.Sizeof(cMedia)))
		if cIndex == nil {
			return
		}
		index := int(*cIndex)

		switch event := Event(event._type); event {
		case MediaListWillAddItem, MediaListItemAdded,
			MediaListWillDeleteItem, MediaListItemDeleted:
			cb(event, media, index)
		}
	}

	// Attach discovery service events.
	events := []Event{
		MediaListWillAddItem,
		MediaListItemAdded,
		MediaListWillDeleteItem,
		MediaListItemDeleted,
	}

	eventIDs := make([]EventID, 0, len(events))
	for _, event := range events {
		eventID, err := manager.attach(event, nil, eventCallback, nil)
		if err != nil {
			return err
		}

		eventIDs = append(eventIDs, eventID)
	}
	defer func() {
		if md.stopFunc == nil {
			manager.Detach(eventIDs...)
		}
	}()

	// Start discovery service.
	if C.libvlc_media_discoverer_start(md.discoverer) < 0 {
		return errOrDefault(getError(), ErrMediaDiscovererStart)
	}

	md.stopFunc = func() {
		// Detach events.
		manager.Detach(eventIDs...)

		// Stop discovery service.
		C.libvlc_media_discoverer_stop(md.discoverer)
	}

	return nil
}

// Stop stops the discovery service.
func (md *MediaDiscoverer) Stop() error {
	if err := md.assertInit(); err != nil {
		return err
	}

	md.stop()
	return nil
}

// IsRunning returns true if the media discovery service is running.
func (md *MediaDiscoverer) IsRunning() bool {
	if err := md.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_discoverer_is_running(md.discoverer) != 0
}

// MediaList returns the media list associated with the discovery service,
// which contains the found media instances.
//   NOTE: The returned media list is read-only.
func (md *MediaDiscoverer) MediaList() (*MediaList, error) {
	if err := md.assertInit(); err != nil {
		return nil, err
	}

	ml := C.libvlc_media_discoverer_media_list(md.discoverer)
	if ml == nil {
		return nil, errOrDefault(getError(), ErrMediaListNotFound)
	}

	// This call will not release the media list. Instead, it will decrement
	// the reference count increased by libvlc_media_discoverer_media_list.
	C.libvlc_media_list_release(ml)

	return &MediaList{list: ml}, nil
}

func (md *MediaDiscoverer) stop() {
	if md.stopFunc != nil {
		md.stopFunc()
		md.stopFunc = nil
	}
}

func (md *MediaDiscoverer) assertInit() error {
	if md == nil || md.discoverer == nil {
		return ErrMediaDiscovererNotInitialized
	}

	return nil
}
