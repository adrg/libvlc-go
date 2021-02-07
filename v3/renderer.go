package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"

// RendererType represents the type of a renderer.
type RendererType string

// Renderer types.
const (
	RendererChromecast RendererType = "chromecast"
)

// RendererFlags contains flags describing a renderer (e.g. capabilities).
type RendererFlags struct {
	AudioEnabled bool
	VideoEnabled bool
}

// Renderer represents a medium capable of rendering media files.
type Renderer struct {
	renderer *C.libvlc_renderer_item_t
}

// Name returns the name of the renderer.
func (r *Renderer) Name() (string, error) {
	if err := r.assertInit(); err != nil {
		return "", nil
	}

	return C.GoString(C.libvlc_renderer_item_name(r.renderer)), nil
}

// Type returns the type of the renderer.
func (r *Renderer) Type() (RendererType, error) {
	if err := r.assertInit(); err != nil {
		return "", nil
	}

	return RendererType(C.GoString(C.libvlc_renderer_item_type(r.renderer))), nil
}

// Flags returns the flags of the renderer.
func (r *Renderer) Flags() (*RendererFlags, error) {
	if err := r.assertInit(); err != nil {
		return nil, err
	}

	flags := C.libvlc_renderer_item_flags(r.renderer)

	return &RendererFlags{
		AudioEnabled: (flags | C.LIBVLC_RENDERER_CAN_AUDIO) != 0,
		VideoEnabled: (flags | C.LIBVLC_RENDERER_CAN_VIDEO) != 0,
	}, nil
}

// IconURI returns the icon URI of the renderer.
func (r *Renderer) IconURI() (string, error) {
	if err := r.assertInit(); err != nil {
		return "", nil
	}

	return C.GoString(C.libvlc_renderer_item_icon_uri(r.renderer)), nil
}

func (r *Renderer) hold() {
	if err := r.assertInit(); err != nil {
		return
	}

	C.libvlc_renderer_item_hold(r.renderer)
}

func (r *Renderer) release() {
	if err := r.assertInit(); err != nil {
		return
	}

	C.libvlc_renderer_item_release(r.renderer)
}

func (r *Renderer) assertInit() error {
	if r == nil || r.renderer == nil {
		return ErrRendererNotInitialized
	}

	return nil
}
