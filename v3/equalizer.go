package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
import "C"

// EqualizerPresetCount returns the number of available equalizer presets.
func EqualizerPresetCount() uint {
	return uint(C.libvlc_audio_equalizer_get_preset_count())
}

// EqualizerPresetName returns the name of the equalizer preset with the
// specified index. The index must be a number greater than or equal to 0
// and less than EqualizerPresetCount(). The function returns an empty string
// for invalid indices.
func EqualizerPresetName(index uint) string {
	return C.GoString(C.libvlc_audio_equalizer_get_preset_name(C.uint(index)))
}

// EqualizerPresetNames returns the names of all available equalizer presets,
// sorted by their indices in ascending order.
func EqualizerPresetNames() []string {
	// Get preset count.
	count := EqualizerPresetCount()

	// Get preset names.
	names := make([]string, 0, count)
	for i := uint(0); i < count; i++ {
		names = append(names, EqualizerPresetName(i))
	}

	return names
}

// EqualizerBandCount returns the number of distinct equalizer frequency bands.
func EqualizerBandCount() uint {
	return uint(C.libvlc_audio_equalizer_get_band_count())
}

// EqualizerBandFrequency returns the frequency of the equalizer band with the
// specified index. The index must be a number greater than or equal to 0 and
// less than EqualizerBandCount(). The function returns -1 for invalid indices.
func EqualizerBandFrequency(index uint) float64 {
	return float64(C.libvlc_audio_equalizer_get_band_frequency(C.uint(index)))
}

// EqualizerBandFrequencies returns the frequencies of all available equalizer
// bands, sorted by their indices in ascending order.
func EqualizerBandFrequencies() []float64 {
	// Get band count.
	count := EqualizerBandCount()

	// Get band frequencies.
	frequencies := make([]float64, 0, count)
	for i := uint(0); i < count; i++ {
		frequencies = append(frequencies, EqualizerBandFrequency(i))
	}

	return frequencies
}

// Equalizer represents an audio equalizer. Use Player.SetEqualizer to assign
// the equalizer to a player instance.
type Equalizer struct {
	equalizer *C.libvlc_equalizer_t
}

// NewEqualizer returns a new equalizer with all frequency values set to zero.
func NewEqualizer() (*Equalizer, error) {
	equalizer := C.libvlc_audio_equalizer_new()
	if equalizer == nil {
		return nil, errOrDefault(getError(), ErrEqualizerCreate)
	}

	return &Equalizer{equalizer: equalizer}, nil
}

// NewEqualizerFromPreset returns a new equalizer with the frequency values
// copied from the preset with the specified index. The index must be a number
// greater than or equal to 0 and less than EqualizerPresetCount().
func NewEqualizerFromPreset(index uint) (*Equalizer, error) {
	equalizer := C.libvlc_audio_equalizer_new_from_preset(C.uint(index))
	if equalizer == nil {
		return nil, errOrDefault(getError(), ErrEqualizerCreate)
	}

	return &Equalizer{equalizer: equalizer}, nil
}

// Release destroys the equalizer instance.
func (e *Equalizer) Release() error {
	if err := e.assertInit(); err != nil {
		return nil
	}

	C.libvlc_audio_equalizer_release(e.equalizer)
	e.equalizer = nil

	return getError()
}

// PreampValue returns the pre-amplification value of the equalizer in Hz.
func (e *Equalizer) PreampValue() (float64, error) {
	if err := e.assertInit(); err != nil {
		return 0, err
	}

	value := C.libvlc_audio_equalizer_get_preamp(e.equalizer)
	return float64(value), getError()
}

// SetPreampValue sets the pre-amplification value of the equalizer.
// The specified amplification value is clamped to the [-20.0, 20.0] Hz range.
func (e *Equalizer) SetPreampValue(value float64) error {
	if err := e.assertInit(); err != nil {
		return err
	}

	if C.libvlc_audio_equalizer_set_preamp(e.equalizer, C.float(value)) != 0 {
		return errOrDefault(getError(), ErrEqualizerAmpValueSet)
	}

	return nil
}

// AmpValueAtIndex returns the amplification value for the equalizer frequency
// band with the specified index, in Hz. The index must be a number greater
// than or equal to 0 and less than EqualizerBandCount().
func (e *Equalizer) AmpValueAtIndex(index uint) (float64, error) {
	if err := e.assertInit(); err != nil {
		return 0, err
	}

	value := C.libvlc_audio_equalizer_get_amp_at_index(e.equalizer, C.uint(index))
	return float64(value), getError()
}

// SetAmpValueAtIndex sets the amplification value for the equalizer frequency
// band with the specified index, in Hz. The index must be a number greater
// than or equal to 0 and less than EqualizerBandCount().
func (e *Equalizer) SetAmpValueAtIndex(value float64, index uint) error {
	if err := e.assertInit(); err != nil {
		return err
	}

	if C.libvlc_audio_equalizer_set_amp_at_index(e.equalizer, C.float(value), C.uint(index)) != 0 {
		return errOrDefault(getError(), ErrEqualizerAmpValueSet)
	}

	return nil
}

func (e *Equalizer) assertInit() error {
	if e == nil || e.equalizer == nil {
		return ErrEqualizerNotInitialized
	}

	return nil
}
