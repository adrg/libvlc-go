package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
	"time"
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
	if mt < MediaTitle || mt > MediaDiscTotal {
		return fmt.Errorf("invalid media meta key: %d", mt)
	}

	return nil
}

// MediaStats contains playback statistics for a media file.
type MediaStats struct {
	// Input statistics.
	ReadBytes    int     // Input bytes read.
	InputBitRate float64 // Input bit rate.

	// Demux statistics.
	DemuxReadBytes     int     // Demux bytes read (demuxed data size).
	DemuxBitRate       float64 // Demux bit rate (content bit rate).
	DemuxCorrupted     int     // Demux corruptions (discarded).
	DemuxDiscontinuity int     // Demux discontinuities (dropped).

	// Video output statistics.
	DecodedVideo      int // Number of decoded video blocks.
	DisplayedPictures int // Number of displayed frames.
	LostPictures      int // Number of lost frames.

	// Audio output statistics.
	DecodedAudio       int // Number of decoded audio blocks.
	PlayedAudioBuffers int // Number of played audio buffers.
	LostAudioBuffers   int // Number of lost audio buffers.
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

// MediaScreenOptions provides configuration options for creating media
// instances from the current computer screen.
type MediaScreenOptions struct {
	// Screen capture area.
	X      int // Left edge coordinate of the subscreen. Default: 0.
	Y      int // Top edge coordinate of the subscreen. Default: 0.
	Width  int // Width of the subscreen. Default: 0 (full screen width).
	Height int // Height of the subscreen. Default: 0 (full screen height).

	// Screen capture frame rate. Default: 0.
	FPS float64

	// Follow the mouse when capturing a subscreen. Default value: false.
	FollowMouse bool

	// Mouse cursor image to use. If specified, the cursor will be overlayed
	// on the captured video. Default: "".
	// NOTE: Windows only.
	CursorImage string

	// Optimize the capture by fragmenting the screen in chunks of predefined
	// height (16 might be a good value). Default: 0 (disabled).
	// NOTE: Windows only.
	FragmentSize int
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

// NewMediaFromScreen creates a Media instance from the current computer
// screen, using the specified options.
// NOTE: This functionality requires the VLC screen module to be installed.
// See installation instructions at https://github.com/adrg/libvlc-go/wiki.
// See https://wiki.videolan.org/Documentation:Modules/screen.
func NewMediaFromScreen(opts *MediaScreenOptions) (*Media, error) {
	media, err := newMedia("screen://", false)
	if err != nil {
		return nil, err
	}
	if opts == nil {
		return media, nil
	}

	var mediaOpts []string
	if opts.X > 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-left=%d", opts.X))
	}
	if opts.Y > 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-top=%d", opts.Y))
	}
	if opts.Width > 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-width=%d", opts.Width))
	}
	if opts.Height > 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-height=%d", opts.Height))
	}
	if opts.FPS != 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-fps=%f", opts.FPS))
	}
	if opts.FollowMouse {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-follow-mouse"))
	}
	if opts.FragmentSize > 0 {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-fragment-size=%d", opts.FragmentSize))
	}
	if opts.CursorImage != "" {
		mediaOpts = append(mediaOpts, fmt.Sprintf(":screen-mouse-image=%s", opts.CursorImage))
	}

	if len(mediaOpts) > 0 {
		return media, media.AddOptions(mediaOpts...)
	}

	return media, nil
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

// Duplicate duplicates the current media instance.
// NOTE: Call the Release method on the returned media in order to free
// the allocated resources.
func (m *Media) Duplicate() (*Media, error) {
	if err := m.assertInit(); err != nil {
		return nil, err
	}

	media := C.libvlc_media_duplicate(m.media)
	if media == nil {
		return nil, errOrDefault(getError(), ErrMediaCreate)
	}

	return &Media{media: media}, nil
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

// State returns the current state of the media instance.
func (m *Media) State() (MediaState, error) {
	if err := m.assertInit(); err != nil {
		return 0, err
	}

	state := int(C.libvlc_media_get_state(m.media))
	return MediaState(state), getError()
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

// Duration returns the media duration in milliseconds.
// NOTE: The duration can only be obtained for parsed media instances. Either
// play the media once or call one of the parsing methods first.
func (m *Media) Duration() (time.Duration, error) {
	if err := m.assertInit(); err != nil {
		return 0, err
	}

	duration := C.libvlc_media_get_duration(m.media)
	return time.Duration(duration) * time.Millisecond, getError()
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

// Parse fetches local art, metadata and track information synchronously.
func (m *Media) Parse() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_parse(m.media)
	return getError()
}

// ParseAsync fetches local art, metadata and track information asynchronously.
// Listen to MediaParsedChanged event on the media event manager the track when
// the parsing has finished. However, if the media was already parsed, the
// event will not be sent.
func (m *Media) ParseAsync() error {
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_parse_async(m.media)
	return getError()
}

// IsParsed returns true if the media was parsed.
func (m *Media) IsParsed() (bool, error) {
	if err := m.assertInit(); err != nil {
		return false, err
	}

	return C.libvlc_media_is_parsed(m.media) != 0, getError()
}

// SubItems returns a media list containing the sub-items of the current
// media instance. If the media does not have any sub-items, an empty media
// list is returned.
// NOTE: Call the Release method on the returned media list in order to free
// the allocated resources.
func (m *Media) SubItems() (*MediaList, error) {
	if err := m.assertInit(); err != nil {
		return nil, err
	}

	var subitems *C.libvlc_media_list_t
	if subitems = C.libvlc_media_subitems(m.media); subitems == nil {
		return nil, errOrDefault(getError(), ErrMediaListNotFound)
	}

	return &MediaList{list: subitems}, nil
}

// Tracks returns the tracks (audio, video, subtitle) of the current media.
// NOTE: The tracks can only be obtained for parsed media instances. Either
// play the media once or call one of the parsing methods first.
func (m *Media) Tracks() ([]*MediaTrack, error) {
	if err := m.assertInit(); err != nil {
		return nil, err
	}

	// Get media tracks.
	var cTracks **C.libvlc_media_track_t

	count := int(C.libvlc_media_tracks_get(m.media, &cTracks))
	if count == 0 || cTracks == nil {
		return nil, nil
	}
	defer C.libvlc_media_tracks_release(cTracks, C.uint(count))

	// Parse media tracks.
	tracks := make([]*MediaTrack, 0, count)
	for i := 0; i < count; i++ {
		// Get current track pointer.
		cTrack := unsafe.Pointer(uintptr(unsafe.Pointer(cTracks)) +
			uintptr(i)*unsafe.Sizeof(*cTracks))
		if cTrack == nil {
			return nil, ErrMediaTrackNotInitialized
		}

		// Parse media track.
		track, err := parseMediaTrack(*(**C.libvlc_media_track_t)(cTrack))
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
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
