// +build !legacy

package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
)

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

// NewListPlayer creates an instance of a multi-media player.
func NewListPlayer() (*ListPlayer, error) {
	if instance == nil {
		return nil, errors.New("Module must be initialized first")
	}

	if player := C.libvlc_media_list_player_new(instance); player != nil {
		return &ListPlayer{player: player}, nil
	}

	return nil, getError()
}

// Release destroys the media player instance.
func (lp *ListPlayer) Release() error {
	if lp.player == nil {
		return nil
	}

	C.libvlc_media_list_player_release(lp.player)
	lp.player = nil

	return getError()
}

// Player returns the underlying Player instance of the ListPlayer.
func (lp *ListPlayer) Player() (*Player, error) {
	if lp.player == nil {
		return nil, errors.New("A list player must be initialized first")
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

// SetPlayer sets the underlying Player instance of the ListPlayer.
func (lp *ListPlayer) SetPlayer(player *Player) error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}
	if player == nil {
		return errors.New("Provided player cannot be nil")
	}
	if player.player == nil {
		return errors.New("Provided player must be initialized first")
	}

	C.libvlc_media_list_player_set_media_player(lp.player, player.player)
	return getError()
}

// Play plays the current media list.
func (lp *ListPlayer) Play() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}
	if lp.IsPlaying() {
		return nil
	}

	C.libvlc_media_list_player_play(lp.player)
	return getError()
}

// PlayNext plays the next media in the current media list.
func (lp *ListPlayer) PlayNext() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	if C.libvlc_media_list_player_next(lp.player) < 0 {
		return getError()
	}

	return nil
}

// PlayPrevious plays the previous media in the current media list.
func (lp *ListPlayer) PlayPrevious() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	if C.libvlc_media_list_player_previous(lp.player) < 0 {
		return getError()
	}

	return nil
}

// PlayAtIndex plays the media at the specified index from the
// current media list.
func (lp ListPlayer) PlayAtIndex(index uint) error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	idx := C.int(index)
	if C.libvlc_media_list_player_play_item_at_index(lp.player, idx) < 0 {
		return getError()
	}

	return nil
}

// IsPlaying returns a boolean value specifying if the player is currently
// playing.
func (lp *ListPlayer) IsPlaying() bool {
	if lp.player == nil {
		return false
	}

	return C.libvlc_media_list_player_is_playing(lp.player) != 0
}

// Stop cancels the currently playing media list, if there is one.
func (lp *ListPlayer) Stop() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}
	if !lp.IsPlaying() {
		return nil
	}

	C.libvlc_media_list_player_stop(lp.player)
	return getError()
}

// TogglePause pauses/resumes the player.
// Calling this method has no effect if there is no media.
func (lp *ListPlayer) TogglePause() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	C.libvlc_media_list_player_pause(lp.player)
	return getError()
}

// SetPlaybackMode sets the player playback mode for the media list.
// By default, it plays the media list once and then stops.
func (lp *ListPlayer) SetPlaybackMode(mode PlaybackMode) error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	m := C.libvlc_playback_mode_t(mode)
	C.libvlc_media_list_player_set_playback_mode(lp.player, m)
	return getError()
}

// MediaState returns the state of the current media.
func (lp *ListPlayer) MediaState() (MediaState, error) {
	if lp.player == nil {
		return 0, errors.New("A list player must be initialized first")
	}

	state := int(C.libvlc_media_list_player_get_state(lp.player))
	return MediaState(state), getError()
}

// MediaList returns the current media list of the player, if one exists
func (lp *ListPlayer) MediaList() *MediaList {
	return lp.list
}

// SetMediaList sets the media list to be played.
func (lp *ListPlayer) SetMediaList(ml *MediaList) error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}
	if ml.list == nil {
		return errors.New("A media list must be initialized first")
	}

	lp.list = ml
	C.libvlc_media_list_player_set_media_list(lp.player, ml.list)

	return getError()
}
