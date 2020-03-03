package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
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

// MediaMetaKey uniquely identifies a type of media metadata.
type MediaMetaKey uint

// Media metadata types.
const (
	MediaTitle MediaMetaKey = iota
	MediaArtist
	MediaGenre
	MediaCopyright
	MediaAlbum
	MediaTrackNumber
	MediaDescription
	MediaRating
	MediaDate
	MediaSetting
	MediaURL
	MediaLanguage
	MediaNowPlaying
	MediaPublisher
	MediaEncodedBy
	MediaArtworkURL
	MediaTrackID
	MediaTrackTotal
	MediaDirector
	MediaSeason
	MediaEpisode
	MediaShowName
	MediaActors
	MediaAlbumArtist
	MediaDiscNumber
	MediaDiscTotal
)

// Validate checks if the media metadata key is valid.
func (mt MediaMetaKey) Validate() error {
	if mt > MediaDiscTotal {
		return fmt.Errorf("invalid media meta key: %d", mt)
	}

	return nil
}

// MediaParseOption defines different options for parsing media files.
type MediaParseOption uint

// Media parse options.
var (
	// Parse media if it is a local file.
	MediaParseLocal MediaParseOption = 0x00

	// Parse media if it is a local or network file.
	MediaParseNetwork MediaParseOption = 0x01

	// Fetch metadata, track information and cover art using local resources.
	MediaFetchLocal MediaParseOption = 0x02

	// Fetch metadata, track information and cover art using network resources.
	MediaFetchNetwork MediaParseOption = 0x04

	// Interact with the user when preparsing media. Set this flag in order to
	// receive a callback if the input media is asking for credentials.
	MediaParseInteract MediaParseOption = 0x08
)

// MediaParseStatus represents the parsing status of a media file.
type MediaParseStatus uint

// Media parse statuses.
const (
	MediaParseUnstarted MediaParseStatus = iota
	MediaParseSkipped
	MediaParseFailed
	MediaParseTimeout
	MediaParseDone
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
	if err := m.assertInit(); err != nil {
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
	if err := m.assertInit(); err != nil {
		return err
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
	if err := m.assertInit(); err != nil {
		return nil, err
	}

	var stats C.libvlc_media_stats_t
	if int(C.libvlc_media_get_stats(m.media, &stats)) != 1 {
		return nil, errOrDefault(getError(), ErrMissingMediaStats)
	}

	return newMediaStats(&stats)
}

// Location returns the media location, which can be either a local path or
// a URL, depending on how the media was loaded.
func (m *Media) Location() (string, error) {
	if err := m.assertInit(); err != nil {
		return "", err
	}

	mrl := C.libvlc_media_get_mrl(m.media)
	if mrl == nil {
		return "", ErrMissingMediaLocation
	}
	defer C.free(unsafe.Pointer(mrl))

	return urlToPath(C.GoString(mrl))
}

// Meta reads the value of the specified media metadata key.
func (m *Media) Meta(key MediaMetaKey) (string, error) {
	if err := m.assertInit(); err != nil {
		return "", err
	}
	if err := key.Validate(); err != nil {
		return "", err
	}

	val := C.libvlc_media_get_meta(m.media, C.libvlc_meta_t(key))
	if val == nil {
		return "", nil
	}
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val), nil
}

// SetMeta sets the specified media metadata key to the provided value.
// In order to save the metadata on the media file, call SaveMeta.
func (m *Media) SetMeta(key MediaMetaKey, val string) error {
	if err := m.assertInit(); err != nil {
		return err
	}
	if err := key.Validate(); err != nil {
		return err
	}

	cVal := C.CString(val)
	C.libvlc_media_set_meta(m.media, C.libvlc_meta_t(key), cVal)
	C.free(unsafe.Pointer(cVal))
	return nil
}

// SaveMeta saves the previously set media metadata.
func (m *Media) SaveMeta() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	if int(C.libvlc_media_save_meta(m.media)) != 1 {
		return errOrDefault(getError(), ErrMediaMetaSave)
	}

	return nil
}

// ParseWithOptions fetches art, metadata and track information asynchronously,
// using the specified options. Listen to the MediaParsedChanged event on the
// media event manager the track when the parsing has finished. However, if the
// media was already parsed, the event is not sent.
// If no option is provided, the media is parsed only if it is a local file.
// The timeout parameter specifies the maximum amount of time allowed to
// preparse the media, in milliseconds.
//   // Timeout values:
//   timeout <  0: use default preparse time.
//   timeout == 0: wait indefinitely.
//   timeout >  0: wait the specified number of milliseconds.
func (m *Media) ParseWithOptions(timeout int, opts ...MediaParseOption) error {
	var flags MediaParseOption
	for _, opt := range opts {
		flags |= opt
	}

	// If no options are specified, parse only local media files.
	if flags == 0 {
		flags = MediaParseLocal
	}

	// Use default preparse timeout for invalid timeout values.
	if timeout < -1 {
		timeout = -1
	}

	if int(C.libvlc_media_parse_with_options(m.media,
		C.libvlc_media_parse_flag_t(flags), C.int(timeout))) != 0 {
		return errOrDefault(getError(), ErrMediaParse)
	}

	return nil
}

// Parse fetches local art, metadata and track information synchronously.
// NOTE: deprecated in libVLC v3.0.0+. Use ParseWithOptions instead.
func (m *Media) Parse() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_parse(m.media)
	return getError()
}

// ParseAsync fetches local art, metadata and track information asynchronously.
// Listen to the MediaParsedChanged event on the media event manager the track
// when the parsing has finished. However, if the media was already parsed,
// the event is not sent.
// NOTE: deprecated in libVLC v3.0.0+. Use ParseWithOptions instead.
func (m *Media) ParseAsync() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_parse_async(m.media)
	return getError()
}

// StopParse stops the parsing of the media. When the media parsing is
// stopped, the MediaParsedChanged event is sent and the parsing status
// of the media is set to MediaParseTimeout.
func (m *Media) StopParse() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_parse_stop(m.media)
	return getError()
}

// ParseStatus returns the parsing status of the media.
func (m *Media) ParseStatus() (MediaParseStatus, error) {
	if err := m.assertInit(); err != nil {
		return MediaParseUnstarted, err
	}

	return MediaParseStatus(C.libvlc_media_get_parsed_status(m.media)), getError()
}

// IsParsed returns true if the media was parsed.
// NOTE: deprecated in libVLC v3.0.0+. Use ParseStatus instead.
func (m *Media) IsParsed() (bool, error) {
	if err := m.assertInit(); err != nil {
		return false, err
	}

	return C.libvlc_media_is_parsed(m.media) != 0, getError()
}

// EventManager returns the event manager responsible for the media.
func (m *Media) EventManager() (*EventManager, error) {
	if err := m.assertInit(); err != nil {
		return nil, err
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

func (m *Media) assertInit() error {
	if m == nil || m.media == nil {
		return ErrMediaNotInitialized
	}

	return nil
}

func newMedia(path string, local bool) (*Media, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	var media *C.libvlc_media_t
	if local {
		if _, err := os.Stat(path); err != nil {
			return nil, err
		}

		media = C.libvlc_media_new_path(inst.handle, cPath)
	} else {
		media = C.libvlc_media_new_location(inst.handle, cPath)
	}

	if media == nil {
		return nil, errOrDefault(getError(), ErrMediaCreate)
	}

	return &Media{media: media}, nil
}
