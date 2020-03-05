package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// Player is a media player used to play a single media file.
// For playing media lists (playlists) use ListPlayer instead.
type Player struct {
	player *C.libvlc_media_player_t
}

// NewPlayer creates an instance of a single-media player.
func NewPlayer() (*Player, error) {
	if err := inst.assertInit(); err != nil {
		return nil, err
	}

	player := C.libvlc_media_player_new(inst.handle)
	if player == nil {
		return nil, errOrDefault(getError(), ErrPlayerCreate)
	}

	return &Player{player: player}, nil
}

// Release destroys the media player instance.
func (p *Player) Release() error {
	if err := p.assertInit(); err != nil {
		return nil
	}

	C.libvlc_media_player_release(p.player)
	p.player = nil

	return getError()
}

// Play plays the current media.
func (p *Player) Play() error {
	if err := p.assertInit(); err != nil {
		return err
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
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_is_playing(p.player) != 0
}

// Stop cancels the currently playing media, if there is one.
func (p *Player) Stop() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_stop(p.player)
	return getError()
}

// SetPause sets the pause state of the media player.
// Pass in true to pause the current media, or false to resume it.
func (p *Player) SetPause(pause bool) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_pause(p.player, C.int(boolToInt(pause)))
	return getError()
}

// TogglePause pauses/resumes the player.
// Calling this method has no effect if there is no media.
func (p *Player) TogglePause() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_pause(p.player)
	return getError()
}

// SetFullScreen sets the fullscreen state of the media player.
// Pass in true to enable fullscreen, or false to disable it.
func (p *Player) SetFullScreen(fullscreen bool) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_set_fullscreen(p.player, C.int(boolToInt(fullscreen)))
	return getError()
}

// ToggleFullScreen toggles the fullscreen status of the player,
// on non-embedded video outputs.
func (p *Player) ToggleFullScreen() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_toggle_fullscreen(p.player)
	return getError()
}

// IsFullScreen gets the fullscreen status of the current player.
func (p *Player) IsFullScreen() (bool, error) {
	if err := p.assertInit(); err != nil {
		return false, err
	}

	return (C.libvlc_get_fullscreen(p.player) != C.int(0)), getError()
}

// Volume returns the volume of the player.
func (p *Player) Volume() (int, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return int(C.libvlc_audio_get_volume(p.player)), getError()
}

// SetVolume sets the volume of the player.
func (p *Player) SetVolume(volume int) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_audio_set_volume(p.player, C.int(volume))
	return getError()
}

// Media returns the current media of the player, if one exists.
func (p *Player) Media() (*Media, error) {
	if err := p.assertInit(); err != nil {
		return nil, err
	}

	media := C.libvlc_media_player_get_media(p.player)
	if media == nil {
		return nil, nil
	}

	// This call will not release the media. Instead, it will decrement
	// the reference count increased by libvlc_media_player_get_media.
	C.libvlc_media_release(media)

	return &Media{media}, nil
}

// SetMedia sets the provided media as the current media of the player.
func (p *Player) SetMedia(m *Media) error {
	return p.setMedia(m)
}

// LoadMediaFromPath loads the media from the specified path and adds it as the
// current media of the player.
func (p *Player) LoadMediaFromPath(path string) (*Media, error) {
	return p.loadMedia(path, true)
}

// LoadMediaFromURL loads the media from the specified URL and adds it as the
// current media of the player.
func (p *Player) LoadMediaFromURL(url string) (*Media, error) {
	return p.loadMedia(url, false)
}

// SetAudioOutput selects an audio output module.
// Any change will take be effect only after playback is stopped and restarted.
// Audio output cannot be changed while playing.
func (p *Player) SetAudioOutput(output string) error {
	if err := p.assertInit(); err != nil {
		return err
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
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return int(C.libvlc_media_player_get_length(p.player)), getError()
}

// MediaState returns the state of the current media.
func (p *Player) MediaState() (MediaState, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	state := int(C.libvlc_media_player_get_state(p.player))
	return MediaState(state), getError()
}

// MediaPosition returns media position as a
// float percentage between 0.0 and 1.0.
func (p *Player) MediaPosition() (float32, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return float32(C.libvlc_media_player_get_position(p.player)), getError()
}

// SetMediaPosition sets media position as percentage between 0.0 and 1.0.
// Some formats and protocols do not support this.
func (p *Player) SetMediaPosition(pos float32) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_position(p.player, C.float(pos))
	return getError()
}

// MediaTime returns media time in milliseconds.
func (p *Player) MediaTime() (int, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return int(C.libvlc_media_player_get_time(p.player)), getError()
}

// SetMediaTime sets the media time in milliseconds.
// Some formats and protocals do not support this.
func (p *Player) SetMediaTime(t int) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_time(p.player, C.libvlc_time_t(int64(t)))
	return getError()
}

// WillPlay returns true if the current media is not in a
// finished or error state.
func (p *Player) WillPlay() bool {
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_will_play(p.player) != 0
}

// XWindow returns the X window the player renders its video output to, or 0
// if no window is set. The window can be set using the SetXWindow method.
func (p *Player) XWindow() (uint32, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return uint32(C.libvlc_media_player_get_xwindow(p.player)), getError()
}

// SetXWindow sets an X Window System drawable where the media player can
// render its video output. If libVLC was built without X11 output support,
// calling this method has no effect.
func (p *Player) SetXWindow(windowID uint32) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_xwindow(p.player, C.uint(windowID))
	return getError()
}

// HWND returns the Windows API window handle the player renders its video
// output to, or 0 if no window is set. The window can be set using the
// SetHWND method.
func (p *Player) HWND() (uintptr, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return uintptr(C.libvlc_media_player_get_hwnd(p.player)), getError()
}

// SetHWND sets a Windows API window handle where the media player can render
// its video output. If libVLC was built without Win32/Win64 API output
// support, calling this method has no effect.
func (p *Player) SetHWND(hwnd uintptr) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_hwnd(p.player, unsafe.Pointer(hwnd))
	return getError()
}

// EventManager returns the event manager responsible for the media player.
func (p *Player) EventManager() (*EventManager, error) {
	if err := p.assertInit(); err != nil {
		return nil, err
	}

	manager := C.libvlc_media_player_event_manager(p.player)
	if manager == nil {
		return nil, ErrMissingEventManager
	}

	return newEventManager(manager), nil
}

func (p *Player) loadMedia(path string, local bool) (*Media, error) {
	m, err := newMedia(path, local)
	if err != nil {
		return nil, err
	}

	if err = p.setMedia(m); err != nil {
		m.Release()
		return nil, err
	}

	return m, nil
}

func (p *Player) setMedia(m *Media) error {
	if err := p.assertInit(); err != nil {
		return err
	}
	if err := m.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_media(p.player, m.media)
	return getError()
}

func (p *Player) assertInit() error {
	if p == nil || p.player == nil {
		return ErrPlayerNotInitialized
	}

	return nil
}
