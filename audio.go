package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "errors"

type AudioOutput struct {
	Name        string
	Description string
}

// AudioOutputList returns a list of audio output devices that can be used
// with an instance of a player.
func AudioOutputList() ([]*AudioOutput, error) {
	if instance == nil {
		return nil, errors.New("Module must be initialized first")
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
