package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// MediaState represents the state of a media file.
type MediaState uint

// Media states.
const (
	MediaNothingSpecial MediaState = iota
	MediaOpening
	MediaBuffering
	MediaPlaying
	MediaPaused
	MediaStopped
	MediaEnded
	MediaError
)

// MediaStats contains playback statistics for a media file.
type MediaStats struct {
	// Input statistics.
	ReadBytes    int     // input bytes read.
	InputBitRate float64 // input bitrate.

	// Demux statistics.
	DemuxReadBytes     int     // demux bytes read (demuxed data size).
	DemuxBitRate       float64 // demux bitrate (content bitrate).
	DemuxCorrupted     int     // demux corruptions (discarded).
	DemuxDiscontinuity int     // demux discontinuities (dropped).

	// Video output statistics.
	DecodedVideo      int // number of decoded video blocks.
	DisplayedPictures int // number of displayed frames.
	LostPictures      int // number of lost frames.

	// Audio output statistics.
	DecodedAudio       int // number of decoded audio blocks.
	PlayedAudioBuffers int // number of played audio buffers.
	LostAudioBuffers   int // number of lost audio buffers.
}

func newMediaStats(st *C.libvlc_media_stats_t) (*MediaStats, error) {
	if st == nil {
		return nil, ErrInvalidMediaStats
	}

	return &MediaStats{
		// Input statistics.
		ReadBytes:    int(st.i_read_bytes),
		InputBitRate: float64(st.f_input_bitrate),

		// Demux statistics.
		DemuxReadBytes:     int(st.i_demux_read_bytes),
		DemuxBitRate:       float64(st.f_demux_bitrate),
		DemuxCorrupted:     int(st.i_demux_corrupted),
		DemuxDiscontinuity: int(st.i_demux_discontinuity),

		// Video output statistics.
		DecodedVideo:      int(st.i_decoded_video),
		DisplayedPictures: int(st.i_displayed_pictures),
		LostPictures:      int(st.i_lost_pictures),

		// Audio output statistics.
		DecodedAudio:       int(st.i_decoded_audio),
		PlayedAudioBuffers: int(st.i_played_abuffers),
		LostAudioBuffers:   int(st.i_lost_abuffers),
	}, nil
}

// Media is an abstract representation of a playable media file.
type Media struct {
	media *C.libvlc_media_t
}

// NewMediaFromPath creates a Media instance from the provided path.
func NewMediaFromPath(path string) (*Media, error) {
	return newMedia(path, true)
}

// NewMediaFromURL creates a Media instance from the provided URL.
func NewMediaFromURL(url string) (*Media, error) {
	return newMedia(url, false)
}

// Release destroys the media instance.
func (m *Media) Release() error {
	if m.media == nil {
		return nil
	}

	C.libvlc_media_release(m.media)
	m.media = nil

	return getError()
}

// AddOptions adds the specified options to the media. The specified options
// determine how a media player reads the media, allowing advanced reading or
// streaming on a per-media basis.
func (m *Media) AddOptions(options ...string) error {
	if m == nil || m.media == nil {
		return ErrMediaNotInitialized
	}

	for _, option := range options {
		if err := m.addOption(option); err != nil {
			return err
		}
	}

	return nil
}

// Stats returns playback statistics for the media.
func (m *Media) Stats() (*MediaStats, error) {
	if m == nil || m.media == nil {
		return nil, ErrMediaNotInitialized
	}

	var stats C.libvlc_media_stats_t
	if int(C.libvlc_media_get_stats(m.media, &stats)) != 1 {
		return nil, errOrDefault(getError(), ErrMissingMediaStats)
	}

	return newMediaStats(&stats)
}

// EventManager returns the event manager responsible for the media.
func (m *Media) EventManager() (*EventManager, error) {
	if m == nil || m.media == nil {
		return nil, ErrMediaNotInitialized
	}

	manager := C.libvlc_media_event_manager(m.media)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}

func (m *Media) addOption(option string) error {
	if option == "" {
		return nil
	}

	cOption := C.CString(option)
	defer C.free(unsafe.Pointer(cOption))

	C.libvlc_media_add_option(m.media, cOption)
	return getError()
}

func newMedia(path string, local bool) (*Media, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t
	if local {
		media = C.libvlc_media_new_path(inst.handle, cPath)
	} else {
		media = C.libvlc_media_new_location(inst.handle, cPath)
	}

	if media == nil {
		return nil, getError()
	}

	return &Media{media: media}, nil
}
