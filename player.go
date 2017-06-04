package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Player struct {
	player *C.libvlc_media_player_t
	media  *Media
}

// NewPlayer creates an instance of a single-media player.
func NewPlayer() (*Player, error) {
	if instance == nil {
		return nil, errors.New("Module must be initialized first")
	}

	if player := C.libvlc_media_player_new(instance); player != nil {
		return &Player{player: player}, nil
	}

	return nil, getError()
}

// Release destroys the media player instance.
func (p *Player) Release() error {
	if p.player == nil {
		return nil
	}

	if p.media != nil {
		if err := p.media.Release(); err != nil {
			return err
		}

		p.media = nil
	}

	C.libvlc_media_player_release(p.player)
	p.player = nil

	return getError()
}

// Play plays the current media.
func (p *Player) Play() error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}
	if p.IsPlaying() {
		return nil
	}

	if C.libvlc_media_player_play(p.player) < 0 {
		return getError()
	}

	return nil
}

// IsPlaying returns a boolean value specifying if the player is currently
// playing.
func (p *Player) IsPlaying() bool {
	if p.player == nil {
		return false
	}

	return C.libvlc_media_player_is_playing(p.player) != 0
}

// Stop cancels the currently playing media, if there is one.
func (p *Player) Stop() error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}
	if !p.IsPlaying() {
		return nil
	}

	C.libvlc_media_player_stop(p.player)
	return getError()
}

// SetPause sets the pause state of the media player.
// Pass in true to pause the current media, or false to resume it.
func (p *Player) SetPause(pause bool) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	toggle := 0
	if pause {
		toggle = 1
	}

	C.libvlc_media_player_set_pause(p.player, C.int(toggle))
	return getError()
}

// TogglePause pauses/resumes the player.
// Calling this method has no effect if there is no media.
func (p *Player) TogglePause() error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	C.libvlc_media_player_pause(p.player)
	return getError()
}

// Volume returns the volume of the player.
func (p *Player) Volume() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must be initialized first")
	}

	return int(C.libvlc_audio_get_volume(p.player)), getError()
}

// SetVolume sets the volume of the player.
func (p *Player) SetVolume(volume int) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	C.libvlc_audio_set_volume(p.player, C.int(volume))
	return getError()
}

// SetMedia sets the provided media as the current media of the player
// The previous player media, if it exists, is automatically released when
// calling this method.
func (p *Player) SetMedia(m *Media) error {
	return p.setMedia(m)
}

// SetMediaFromPath loads the media from the specified path and adds it as the
// current media of the player.
// The previous player media, if it exists, is automatically released when
// calling this method.
func (p *Player) SetMediaFromPath(path string) error {
	m, err := newMedia(path, true)
	if err != nil {
		return err
	}

	return p.setMedia(m)
}

// SetMediaFromURL loads the media from the specified URL and adds it as the
// current media of the player.
// The previous player media, if it exists, is automatically released when
// calling this method.
func (p *Player) SetMediaFromURL(url string) error {
	m, err := newMedia(url, false)
	if err != nil {
		return err
	}

	return p.setMedia(m)
}

// SetAudioOutput selects an audio output module.
// Any change will take be effect only after playback is stopped and restarted.
// Audio output cannot be changed while playing.
func (p *Player) SetAudioOutput(output string) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	cOutput := C.CString(output)
	defer C.free(unsafe.Pointer(cOutput))

	if C.libvlc_audio_output_set(p.player, cOutput) != 0 {
		return getError()
	}

	return nil
}

// MediaLength returns media length in milliseconds.
func (p *Player) MediaLength() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must be initialized first")
	}

	return int(C.libvlc_media_player_get_length(p.player)), getError()
}

// MediaState returns the state of the current media.
func (p *Player) MediaState() (MediaState, error) {
	if p.player == nil {
		return 0, errors.New("A player must be initialized first")
	}

	state := int(C.libvlc_media_player_get_state(p.player))
	return MediaState(state), getError()
}

// MediaPosition returns media position as a
// float percentage between 0.0 and 1.0.
func (p *Player) MediaPosition() (float32, error) {
	if p.player == nil {
		return 0, errors.New("A player must be initialized first")
	}

	return float32(C.libvlc_media_player_get_position(p.player)), getError()
}

// SetMediaPosition sets media position as percentage between 0.0 and 1.0.
// Some formats and protocols do not support this.
func (p *Player) SetMediaPosition(pos float32) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	C.libvlc_media_player_set_position(p.player, C.float(pos))
	return getError()
}

// MediaTime returns media time in milliseconds.
func (p *Player) MediaTime() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must be initialized first")
	}

	return int(C.libvlc_media_player_get_time(p.player)), getError()
}

// SetMediaTime sets the media time in milliseconds.
// Some formats and protocals do not support this.
func (p *Player) SetMediaTime(t int) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}

	C.libvlc_media_player_set_time(p.player, C.libvlc_time_t(int64(t)))
	return getError()
}

// WillPlay returns true if the current media is not in a
// finished or error state.
func (p *Player) WillPlay() bool {
	if p.player == nil {
		return false
	}

	return C.libvlc_media_player_will_play(p.player) != 0
}

func (p *Player) setMedia(m *Media) error {
	if p.player == nil {
		return errors.New("A player must be initialized first")
	}
	if m.media == nil {
		return errors.New("The media must be initialized first")
	}

	if p.media != nil {
		if err := p.media.Release(); err != nil {
			return err
		}

		p.media = nil
	}

	p.media = m
	C.libvlc_media_player_set_media(p.player, m.media)

	return getError()
}
