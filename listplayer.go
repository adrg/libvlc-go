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
	media  *MediaList
}

func NewListPlayer() (*ListPlayer, error) {
	if instance == nil {
		return nil, errors.New("Module must be first initialized")
	}

	if player := C.libvlc_media_list_player_new(instance); player != nil {
		return &ListPlayer{player: player}, nil
	}

	return nil, getError()
}

func (p *ListPlayer) Release() error {
	if p.player == nil {
		return nil
	}

	if p.media != nil {
		if err := p.media.Release(); err != nil {
			return err
		}

		p.media = nil
	}

	C.libvlc_media_list_player_release(p.player)
	return getError()
}

func (p *ListPlayer) Play() error {
	if p.player == nil {
		return errors.New("A list player must first be initialized")
	}

	if p.IsPlaying() {
		return nil
	}

	C.libvlc_media_list_player_play(p.player)
	return nil
}

func (p *ListPlayer) IsPlaying() bool {
	if p.player == nil {
		return false
	}

	return C.libvlc_media_list_player_is_playing(p.player) != 0
}

func (p *ListPlayer) Stop() error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	if !p.IsPlaying() {
		return nil
	}

	C.libvlc_media_list_player_stop(p.player)
	return getError()
}

// SetPause sets the pause state of the media player.
// Pass in true to pause the current media, or false to resume it.
func (p *ListPlayer) Pause() error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	C.libvlc_media_list_player_pause(p.player)
	return getError()
}

// TogglePause pauses/resumes the player.
// Calling this method has no effect if there is no media.
func (p *ListPlayer) TogglePause() error {
	if p.player == nil {
		return errors.New("A list player must first be initialized")
	}

	C.libvlc_media_list_player_pause(p.player)
	return getError()
}

func (p *ListPlayer) SetMediaFromPath(path string) error {

	return nil
}

func (p *ListPlayer) SetMediaList(l *MediaList) error {
	if p.player == nil || l.list == nil {
		return errors.New("A MediaList and ListPlayer must first be initialized")
	}

	C.libvlc_media_list_player_set_media_list(p.player, l.list)
	return getError()
}
