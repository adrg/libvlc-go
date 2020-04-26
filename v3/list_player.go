// +build !legacy

package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"

// PlaybackMode defines playback modes for a media list.
type PlaybackMode uint

// Playback modes.
const (
	Default PlaybackMode = iota
	Loop
	Repeat
)

// ListPlayer is an enhanced media player used to play media lists.
type ListPlayer struct {
	player *C.libvlc_media_list_player_t
	list   *MediaList
}

// NewListPlayer creates a new list player instance.
func NewListPlayer() (*ListPlayer, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	player := C.libvlc_media_list_player_new(inst.handle)
	if player == nil {
		return nil, errOrDefault(getError(), ErrListPlayerCreate)
	}

	return &ListPlayer{player: player}, nil
}

// Release destroys the list player instance.
func (lp *ListPlayer) Release() error {
	if err := lp.assertInit(); err != nil {
		return nil
	}

	C.libvlc_media_list_player_release(lp.player)
	lp.player = nil

	return getError()
}

// Player returns the underlying Player instance of the list player.
func (lp *ListPlayer) Player() (*Player, error) {
	if err := lp.assertInit(); err != nil {
		return nil, err
	}

	player := C.libvlc_media_list_player_get_media_player(lp.player)
	if player == nil {
		return nil, getError()
	}

	// This call will not release the player. Instead, it will decrement the
	// reference count increased by libvlc_media_list_player_get_media_player.
	C.libvlc_media_player_release(player)

	return &Player{player: player}, nil
}

// SetPlayer sets the underlying Player instance of the list player.
func (lp *ListPlayer) SetPlayer(player *Player) error {
	if err := lp.assertInit(); err != nil {
		return err
	}
	if err := player.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_player_set_media_player(lp.player, player.player)
	return getError()
}

// Play plays the current media list.
func (lp *ListPlayer) Play() error {
	if err := lp.assertInit(); err != nil {
		return err
	}
	if lp.IsPlaying() {
		return nil
	}

	C.libvlc_media_list_player_play(lp.player)
	return getError()
}

// PlayNext plays the next media in the current media list.
func (lp *ListPlayer) PlayNext() error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	if C.libvlc_media_list_player_next(lp.player) < 0 {
		return getError()
	}

	return nil
}

// PlayPrevious plays the previous media in the current media list.
func (lp *ListPlayer) PlayPrevious() error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	if C.libvlc_media_list_player_previous(lp.player) < 0 {
		return getError()
	}

	return nil
}

// PlayAtIndex plays the media at the specified index from the
// current media list.
func (lp *ListPlayer) PlayAtIndex(index uint) error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	idx := C.int(index)
	if C.libvlc_media_list_player_play_item_at_index(lp.player, idx) < 0 {
		return getError()
	}

	return nil
}

// PlayItem plays the specified media item. The item must be part of the
// current media list of the player.
func (lp *ListPlayer) PlayItem(m *Media) error {
	if err := lp.assertInit(); err != nil {
		return err
	}
	if err := m.assertInit(); err != nil {
		return err
	}

	if C.libvlc_media_list_player_play_item(lp.player, m.media) < 0 {
		return errOrDefault(getError(), ErrMediaNotFound)
	}

	return nil
}

// IsPlaying returns a boolean value specifying if the player is currently
// playing.
func (lp *ListPlayer) IsPlaying() bool {
	if err := lp.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_list_player_is_playing(lp.player) != 0
}

// Stop cancels the currently playing media list, if there is one.
func (lp *ListPlayer) Stop() error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_player_stop(lp.player)
	return getError()
}

// SetPause sets the pause state of the list player.
// Pass in true to pause the current media, or false to resume it.
func (lp *ListPlayer) SetPause(pause bool) error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_player_set_pause(lp.player, C.int(boolToInt(pause)))
	return getError()
}

// TogglePause pauses/resumes the player.
// Calling this method has no effect if there is no media.
func (lp *ListPlayer) TogglePause() error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_list_player_pause(lp.player)
	return getError()
}

// SetPlaybackMode sets the player playback mode for the media list.
// By default, it plays the media list once and then stops.
func (lp *ListPlayer) SetPlaybackMode(mode PlaybackMode) error {
	if err := lp.assertInit(); err != nil {
		return err
	}

	m := C.libvlc_playback_mode_t(mode)
	C.libvlc_media_list_player_set_playback_mode(lp.player, m)
	return getError()
}

// MediaState returns the state of the current media.
func (lp *ListPlayer) MediaState() (MediaState, error) {
	if err := lp.assertInit(); err != nil {
		return 0, err
	}

	state := int(C.libvlc_media_list_player_get_state(lp.player))
	return MediaState(state), getError()
}

// MediaList returns the current media list of the player, if one exists.
func (lp *ListPlayer) MediaList() *MediaList {
	return lp.list
}

// SetMediaList sets the media list to be played.
func (lp *ListPlayer) SetMediaList(ml *MediaList) error {
	if err := lp.assertInit(); err != nil {
		return err
	}
	if err := ml.assertInit(); err != nil {
		return err
	}

	lp.list = ml
	C.libvlc_media_list_player_set_media_list(lp.player, ml.list)

	return getError()
}

// EventManager returns the event manager responsible for the list player.
func (lp *ListPlayer) EventManager() (*EventManager, error) {
	if err := lp.assertInit(); err != nil {
		return nil, err
	}

	manager := C.libvlc_media_list_player_event_manager(lp.player)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}

func (lp *ListPlayer) assertInit() error {
	if lp == nil || lp.player == nil {
		return ErrListPlayerNotInitialized
	}

	return nil
}
