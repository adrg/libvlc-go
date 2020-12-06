package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"

// AudioOutput is an abstraction for rendering decoded (or pass-through)
// audio samples.
type AudioOutput struct {
	Name        string
	Description string
}

// AudioOutputList returns a list of audio output devices that can be used
// with an instance of a player.
func AudioOutputList() ([]*AudioOutput, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	cOutputs := C.libvlc_audio_output_list_get(inst.handle)
	if cOutputs == nil {
		return nil, errOrDefault(getError(), ErrAudioOutputListMissing)
	}

	var outputs []*AudioOutput
	for n := cOutputs; n != nil; n = (*C.libvlc_audio_output_t)(n.p_next) {
		outputs = append(outputs, &AudioOutput{
			Name:        C.GoString(n.psz_name),
			Description: C.GoString(n.psz_description),
		})
	}

	C.libvlc_audio_output_list_release(cOutputs)
	return outputs, getError()
}

// ModuleDescription contains information about a libVLC module.
type ModuleDescription struct {
	Name      string
	ShortName string
	LongName  string
	Help      string
}

// ListAudioFilters returns the list of available audio filters.
func ListAudioFilters() ([]*ModuleDescription, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	return parseFilterList(C.libvlc_audio_filter_list_get(inst.handle))
}

// ListVideoFilters returns the list of available video filters.
func ListVideoFilters() ([]*ModuleDescription, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	return parseFilterList(C.libvlc_video_filter_list_get(inst.handle))
}

func parseFilterList(cFilters *C.libvlc_module_description_t) ([]*ModuleDescription, error) {
	if cFilters == nil {
		return nil, errOrDefault(getError(), ErrFilterListMissing)
	}

	var filters []*ModuleDescription
	for n := cFilters; n != nil; n = (*C.libvlc_module_description_t)(n.p_next) {
		filters = append(filters, &ModuleDescription{
			Name:      C.GoString(n.psz_name),
			ShortName: C.GoString(n.psz_shortname),
			LongName:  C.GoString(n.psz_longname),
			Help:      C.GoString(n.psz_help),
		})
	}

	C.libvlc_module_description_list_release(cFilters)
	return filters, getError()
}
