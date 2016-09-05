package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

var instance *C.libvlc_instance_t

func Init(args ...string) error {
	if instance != nil {
		return nil
	}

	argc := len(args)
	argv := make([]*C.char, argc)

	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	defer func() {
		for i := range argv {
			C.free(unsafe.Pointer(argv[i]))
		}
	}()

	instance = C.libvlc_new(C.int(argc), *(***C.char)(unsafe.Pointer(&argv)))
	if instance == nil {
		return getError()
	}

	return nil
}

func Release() error {
	if instance == nil {
		return nil
	}

	C.libvlc_release(instance)
	return getError()
}

func NewPlayer() (*Player, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	if player := C.libvlc_media_player_new(instance); player != nil {
		return &Player{player: player}, nil
	}

	return nil, getError()
}

func AudioOutputList() ([]*AudioOutput, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	outputs := C.libvlc_audio_output_list_get(instance)
	if outputs == nil {
		return nil, getError()
	}
	defer C.libvlc_audio_output_list_release(outputs)

	audioOutputs := []*AudioOutput{}
	for p := outputs; p != nil; p = (*C.libvlc_audio_output_t)(p.p_next) {
		audioOutput := &AudioOutput{
			Name:        C.GoString(p.psz_name),
			Description: C.GoString(p.psz_description),
		}

		audioOutputs = append(audioOutputs, audioOutput)
	}

	return audioOutputs, getError()
}

func newMedia(path string, local bool) (*Media, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t = nil
	if local {
		media = C.libvlc_media_new_path(instance, cPath)
	} else {
		media = C.libvlc_media_new_location(instance, cPath)
	}

	if media == nil {
		return nil, getError()
	}

	return &Media{media: media}, nil
}

func getError() error {
	msg := C.libvlc_errmsg()
	if msg != nil {
		return errors.New(C.GoString(msg))
	}

	return nil
}
