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
	return getError()
}

func (p *Player) Play() error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	if p.IsPlaying() {
		return nil
	}

	if C.libvlc_media_player_play(p.player) < 0 {
		return getError()
	}

	return nil
}

func (p *Player) IsPlaying() bool {
	if p.player == nil {
		return false
	}

	return C.libvlc_media_player_is_playing(p.player) != 0
}

func (p *Player) Stop() error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	if !p.IsPlaying() {
		return nil
	}

	C.libvlc_media_player_stop(p.player)
	return getError()
}

func (p *Player) Pause(pause bool) error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	toggle := 0
	if pause {
		toggle = 1
	}

	C.libvlc_media_player_set_pause(p.player, C.int(toggle))
	return getError()
}

func (p *Player) GetVolume() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must first be initialized")
	}

	return int(C.libvlc_audio_get_volume(p.player)), getError()
}

func (p *Player) SetVolume(volume int) error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	C.libvlc_audio_set_volume(p.player, C.int(volume))
	return getError()
}

func (p *Player) SetMedia(path string, local bool) error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	if p.media != nil {
		if err := p.media.Release(); err != nil {
			return err
		}

		p.media = nil
	}

	media, err := newMedia(path, local)
	if err != nil {
		return err
	}

	p.media = media
	C.libvlc_media_player_set_media(p.player, media.media)

	return getError()
}

func (p *Player) SetAudioOutput(output string) error {
	if p.player == nil {
		return errors.New("A player must first be initialized")
	}

	cOutput := C.CString(output)
	defer C.free(unsafe.Pointer(cOutput))

	if C.libvlc_audio_output_set(p.player, cOutput) != 0 {
		return getError()
	}

	return nil
}

// GetTime returns media time in milliseconds
func (p *Player) GetTime() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must first be initialized")
	}
	return int(C.libvlc_media_player_get_time(p.player)), getError()
}

// GetLength returns media length in milliseconds.
func (p *Player) GetLength() (int, error) {
	if p.player == nil {
		return 0, errors.New("A player must first be initialized")
	}
	return int(C.libvlc_media_player_get_length(p.player)), getError()
}
