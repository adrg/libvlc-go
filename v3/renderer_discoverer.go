package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

// RendererDiscoveryCallback is used by renderer discovery services to
// report discovery events.
type RendererDiscoveryCallback func(Event, *Renderer)

// RendererDiscovererDescriptor contains information about a renderer
// discovery service. Pass the `Name` field to the NewRendererDiscoverer
// method in order to create a new discovery service instance.
type RendererDiscovererDescriptor struct {
	Name     string
	LongName string
}

// ListRendererDiscoverers returns a list of descriptors identifying the
// available renderer discovery services.
func ListRendererDiscoverers() ([]*RendererDiscovererDescriptor, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	// Get renderer discoverer descriptors.
	var cDescriptors **C.libvlc_rd_description_t

	count := int(C.libvlc_renderer_discoverer_list_get(inst.handle, &cDescriptors))
	if count <= 0 || cDescriptors == nil {
		return nil, nil
	}
	defer C.libvlc_renderer_discoverer_list_release(cDescriptors, C.size_t(count))

	// Parse renderer discoverer descriptors.
	descriptors := make([]*RendererDiscovererDescriptor, 0, count)
	for i := 0; i < count; i++ {
		// Get current renderer discoverer descriptor.
		cDescriptorPtr := unsafe.Pointer(uintptr(unsafe.Pointer(cDescriptors)) +
			uintptr(i)*unsafe.Sizeof(*cDescriptors))
		if cDescriptorPtr == nil {
			return nil, ErrRendererDiscovererParse
		}

		cDescriptor := *(**C.libvlc_rd_description_t)(cDescriptorPtr)
		if cDescriptor == nil {
			return nil, ErrRendererDiscovererParse
		}

		// Parse renderer discoverer descriptor.
		descriptors = append(descriptors, &RendererDiscovererDescriptor{
			Name:     C.GoString(cDescriptor.psz_name),
			LongName: C.GoString(cDescriptor.psz_longname),
		})
	}

	return descriptors, nil
}

// RendererDiscoverer represents a renderer discovery service.
// Discovery services use different discovery protocols (e.g. mDNS)
// in order to find available media renderers (e.g. Chromecast).
type RendererDiscoverer struct {
	sync.Mutex

	discoverer *C.libvlc_renderer_discoverer_t
	renderers  map[*C.libvlc_renderer_item_t]*Renderer
	eventIDs   []EventID
	isRunning  bool
}

// NewRendererDiscoverer instantiates the renderer discovery service
// identified by the specified name. Use the ListRendererDiscoverers
// method to obtain the list of available discovery service descriptors.
// NOTE: Call the Release method on the discovery service instance in
// order to free the allocated resources.
func NewRendererDiscoverer(name string) (*RendererDiscoverer, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	discoverer := C.libvlc_renderer_discoverer_new(inst.handle, cName)
	if discoverer == nil {
		return nil, errOrDefault(getError(), ErrRendererDiscovererCreate)
	}

	return &RendererDiscoverer{
		discoverer: discoverer,
		renderers:  map[*C.libvlc_renderer_item_t]*Renderer{},
	}, nil
}

// Release stops and destroys the renderer discovery service along
// with all the renderers found by the instance.
func (rd *RendererDiscoverer) Release() error {
	if err := rd.assertInit(); err != nil {
		return nil
	}

	rd.Lock()
	defer rd.Unlock()

	// Stop discovery service.
	if err := rd.stop(); err != nil {
		return err
	}

	// Release renderers.
	for _, renderer := range rd.renderers {
		renderer.release()
	}
	rd.renderers = nil

	// Release discovery service.
	C.libvlc_renderer_discoverer_release(rd.discoverer)
	rd.discoverer = nil

	return getError()
}

// Start starts the renderer discovery service and reports discovery
// events through the specified callback function.
func (rd *RendererDiscoverer) Start(cb RendererDiscoveryCallback) error {
	if err := rd.assertInit(); err != nil {
		return nil
	}

	rd.Lock()
	defer rd.Unlock()

	// Check if the discovery service is running.
	if rd.isRunning {
		return nil
	}

	// Retrieve event manager.
	manager, err := rd.eventManager()
	if err != nil {
		return err
	}

	// Create event callback.
	eventCallback := func(event *C.libvlc_event_t, userData interface{}) {
		if err := rd.assertInit(); err != nil {
			return
		}
		if event == nil {
			return
		}

		rd.Lock()
		defer rd.Unlock()

		cRenderer := *(**C.libvlc_renderer_item_t)(unsafe.Pointer(&event.u[0]))
		if cRenderer == nil {
			return
		}

		renderer, ok := rd.renderers[cRenderer]
		if !ok {
			renderer = &Renderer{renderer: cRenderer}
			rd.renderers[cRenderer] = renderer
			renderer.hold()
		}

		switch event := Event(event._type); event {
		case RendererDiscovererItemAdded:
			cb(event, renderer)
		case RendererDiscovererItemDeleted:
			cb(event, renderer)
			delete(rd.renderers, cRenderer)
			renderer.release()
		}
	}

	// Attach discovery service events.
	events := []Event{
		RendererDiscovererItemAdded,
		RendererDiscovererItemDeleted,
	}

	var eventIDs []EventID
	for _, event := range events {
		eventID, err := manager.attach(event, nil, eventCallback, nil)
		if err != nil {
			return err
		}

		eventIDs = append(eventIDs, eventID)
	}
	defer func() {
		if !rd.isRunning {
			manager.Detach(eventIDs...)
		}
	}()

	// Start discovery service.
	if C.libvlc_renderer_discoverer_start(rd.discoverer) < 0 {
		return errOrDefault(getError(), ErrRendererDiscovererStart)
	}

	rd.isRunning = true
	rd.eventIDs = eventIDs
	return nil
}

// Stop stops the discovery service.
func (rd *RendererDiscoverer) Stop() error {
	if err := rd.assertInit(); err != nil {
		return err
	}

	rd.Lock()
	defer rd.Unlock()
	return rd.stop()
}

func (rd *RendererDiscoverer) stop() error {
	// Check if the discovery service is stopped.
	if !rd.isRunning {
		return nil
	}

	// Detach events.
	manager, err := rd.eventManager()
	if err != nil {
		return err
	}

	manager.Detach(rd.eventIDs...)
	rd.eventIDs = nil

	// Stop discovery service.
	C.libvlc_renderer_discoverer_stop(rd.discoverer)

	rd.isRunning = false
	return nil
}

// eventManager returns the event manager responsible for the renderer
// discovery service.
func (rd *RendererDiscoverer) eventManager() (*EventManager, error) {
	if err := rd.assertInit(); err != nil {
		return nil, err
	}

	manager := C.libvlc_renderer_discoverer_event_manager(rd.discoverer)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}

func (rd *RendererDiscoverer) assertInit() error {
	if rd == nil || rd.discoverer == nil {
		return ErrRendererDiscovererNotInitialized
	}

	return nil
}
