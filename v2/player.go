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

// WillPlay returns true if the current media is not in a finished or
// error state.
func (p *Player) WillPlay() bool {
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_will_play(p.player) != 0
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

// TogglePause pauses or resumes the player, depending on its current status.
// Calling this method has no effect if there is no media.
func (p *Player) TogglePause() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_pause(p.player)
	return getError()
}

// CanPause returns true if the media player can be paused.
func (p *Player) CanPause() bool {
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_can_pause(p.player) != 0
}

// IsSeekable returns true if the current media is seekable.
func (p *Player) IsSeekable() bool {
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_is_seekable(p.player) != 0
}

// VideoOutputCount returns the number of video outputs the media player has.
func (p *Player) VideoOutputCount() int {
	if err := p.assertInit(); err != nil {
		return 0
	}

	return int(C.libvlc_media_player_has_vout(p.player))
}

// IsScrambled returns true if the media player is in a scrambled state.
func (p *Player) IsScrambled() bool {
	if err := p.assertInit(); err != nil {
		return false
	}

	return C.libvlc_media_player_program_scrambled(p.player) != 0
}

// PlaybackRate returns the playback rate of the media player.
// NOTE: Depending on the underlying media, the returned rate may be
// different from the real playback rate.
func (p *Player) PlaybackRate() float32 {
	if err := p.assertInit(); err != nil {
		return 0
	}

	return float32(C.libvlc_media_player_get_rate(p.player))
}

// SetPlaybackRate sets the playback rate of the media player.
// NOTE: Depending on the underlying media, changing the playback rate
// might not be supported.
func (p *Player) SetPlaybackRate(rate float32) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_rate(p.player, C.float(rate))
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

// IsFullScreen returns the fullscreen status of the player.
func (p *Player) IsFullScreen() (bool, error) {
	if err := p.assertInit(); err != nil {
		return false, err
	}

	return C.libvlc_get_fullscreen(p.player) != C.int(0), getError()
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

// IsMuted returns a boolean value that specifies whether the audio
// output of the player is muted.
func (p *Player) IsMuted() (bool, error) {
	if err := p.assertInit(); err != nil {
		return false, err
	}

	return C.libvlc_audio_get_mute(p.player) > C.int(0), getError()
}

// SetMute mutes or unmutes the audio output of the player.
// NOTE: If there is no active audio playback stream, the mute status might not
// be available. If digital pass-through (S/PDIF, HDMI, etc.) is in use, muting
// may not be applicable. Some audio output plugins do not support muting.
func (p *Player) SetMute(mute bool) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_audio_set_mute(p.player, C.int(boolToInt(mute)))
	return getError()
}

// ToggleMute mutes or unmutes the audio output of the player, depending on
// the current status.
// NOTE: If there is no active audio playback stream, the mute status might not
// be available. If digital pass-through (S/PDIF, HDMI, etc.) is in use, muting
// may not be applicable. Some audio output plugins do not support muting.
func (p *Player) ToggleMute() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_audio_toggle_mute(p.player)
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

// SetMediaTime sets the media time in milliseconds. Some formats and
// protocols do not support this.
func (p *Player) SetMediaTime(t int) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_time(p.player, C.libvlc_time_t(int64(t)))
	return getError()
}

// NextFrame displays the next video frame, if supported.
func (p *Player) NextFrame() error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_next_frame(p.player)
	return getError()
}

// Scale returns the scaling factor of the current video. A scaling factor
// of zero means the video is configured to fit in the available space.
func (p *Player) Scale() (float64, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return float64(C.libvlc_video_get_scale(p.player)), getError()
}

// SetScale sets the scaling factor of the current video. The scaling factor
// is the ratio of the number of pixels displayed on the screen to the number
// of pixels in the original decoded video. A scaling factor of zero adjusts
// the video to fit in the available space.
// NOTE: Not all video outputs support scaling.
func (p *Player) SetScale(scale float64) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_video_set_scale(p.player, C.float(scale))
	return getError()
}

// AspectRatio returns the aspect ratio of the current video.
func (p *Player) AspectRatio() (string, error) {
	if err := p.assertInit(); err != nil {
		return "", err
	}

	aspectRatio := C.libvlc_video_get_aspect_ratio(p.player)
	if aspectRatio == nil {
		return "", getError()
	}
	defer C.free(unsafe.Pointer(aspectRatio))

	return C.GoString(aspectRatio), getError()
}

// SetAspectRatio sets the aspect ratio of the current video (e.g. `16:9`).
// NOTE: Invalid aspect ratios are ignored.
func (p *Player) SetAspectRatio(aspectRatio string) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	cAspectRatio := C.CString(aspectRatio)
	C.libvlc_video_set_aspect_ratio(p.player, cAspectRatio)
	C.free(unsafe.Pointer(cAspectRatio))
	return getError()
}

// XWindow returns the identifier of the X window the media player is
// configured to render its video output to, or 0 if no window is set.
// The window can be set using the SetXWindow method.
// NOTE: The window identifier is returned even if the player is not currently
// using it (for instance if it is playing an audio-only input).
func (p *Player) XWindow() (uint32, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return uint32(C.libvlc_media_player_get_xwindow(p.player)), getError()
}

// SetXWindow sets an X Window System drawable where the media player can
// render its video output. The call takes effect when the playback starts.
// If it is already started, it might need to be stopped before changes apply.
// If libVLC was built without X11 output support, calling this method has no
// effect.
// NOTE: By default, libVLC captures input events on the video rendering area.
// Use the SetMouseInput and SetKeyInput methods if you want to handle input
// events in your application. By design, the X11 protocol delivers input
// events to only one recipient.
func (p *Player) SetXWindow(windowID uint32) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_xwindow(p.player, C.uint(windowID))
	return getError()
}

// HWND returns the handle of the Windows API window the media player is
// configured to render its video output to, or 0 if no window is set.
// The window can be set using the SetHWND method.
// NOTE: The window handle is returned even if the player is not currently
// using it (for instance if it is playing an audio-only input).
func (p *Player) HWND() (uintptr, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return uintptr(C.libvlc_media_player_get_hwnd(p.player)), getError()
}

// SetHWND sets a Windows API window handle where the media player can render
// its video output. If libVLC was built without Win32/Win64 API output
// support, calling this method has no effect.
// NOTE: By default, libVLC captures input events on the video rendering area.
// Use the SetMouseInput and SetKeyInput methods if you want to handle input
// events in your application.
func (p *Player) SetHWND(hwnd uintptr) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_hwnd(p.player, unsafe.Pointer(hwnd))
	return getError()
}

// NSObject returns the handler of the NSView the media player is configured
// to render its video output to, or 0 if no view is set. See SetNSObject.
func (p *Player) NSObject() (uintptr, error) {
	if err := p.assertInit(); err != nil {
		return 0, err
	}

	return uintptr(C.libvlc_media_player_get_nsobject(p.player)), getError()
}

// SetNSObject sets a NSObject handler where the media player can render
// its video output. Use the vout called "macosx". The object can be a NSView
// or a NSObject following the VLCVideoViewEmbedding protocol.
//   @protocol VLCVideoViewEmbedding <NSObject>
//   - (void)addVoutSubview:(NSView *)view;
//   - (void)removeVoutSubview:(NSView *)view;
//   @end
func (p *Player) SetNSObject(drawable uintptr) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_media_player_set_nsobject(p.player, unsafe.Pointer(drawable))
	return getError()
}

// SetKeyInput enables or disables key press event handling, according to the
// libVLC hotkeys configuration. By default, keyboard events are handled by
// the libVLC video widget.
// NOTE: This method works only for X11 and Win32 at the moment.
// NOTE: On X11, there can be only one subscriber for key press and mouse click
// events per window. If your application has subscribed to these events for
// the X window ID of the video widget, then libVLC will not be able to handle
// key presses and mouse clicks.
func (p *Player) SetKeyInput(enable bool) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_video_set_key_input(p.player, C.uint(boolToInt(enable)))
	return getError()
}

// SetMouseInput enables or disables mouse click event handling. By default,
// mouse events are handled by the libVLC video widget. This is needed for DVD
// menus to work, as well as for a few video filters, such as "puzzle".
// NOTE: This method works only for X11 and Win32 at the moment.
// NOTE: On X11, there can be only one subscriber for key press and mouse click
// events per window. If your application has subscribed to these events for
// the X window ID of the video widget, then libVLC will not be able to handle
// key presses and mouse clicks.
func (p *Player) SetMouseInput(enable bool) error {
	if err := p.assertInit(); err != nil {
		return err
	}

	C.libvlc_video_set_mouse_input(p.player, C.uint(boolToInt(enable)))
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
