package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
)

type ListPlayer struct {
	player *C.libvlc_media_list_player_t
	list   *MediaList
}

// NewPlayer creates an instance of a multi-media player.
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

// SetPause sets the pause state of the media player.
// Pass in true to pause the current media, or false to resume it.
func (lp *ListPlayer) Pause() error {
	if lp.player == nil {
		return errors.New("A list player must be initialized first")
	}

	C.libvlc_media_list_player_pause(lp.player)
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
