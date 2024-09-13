package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>
#include <stdlib.h>
*/
import "C"
import (
	"image/color"
	"time"
	"unsafe"
)

// Marquee represents a marquee text than can be displayed over a media
// instance, along with its visual properties.
//
//	For more information see https://wiki.videolan.org/Documentation:Modules/marq.
type Marquee struct {
	player *Player
}

func newMarquee(player *Player) *Marquee {
	return &Marquee{
		player: player,
	}
}

// Enable enables or disables the marquee. By default, the marquee is disabled.
func (m *Marquee) Enable(enable bool) error {
	return m.setInt(C.libvlc_marquee_Enable, boolToInt(enable))
}

// Text returns the marquee text.
// Default: "".
func (m *Marquee) Text() (string, error) {
	return m.getString(C.libvlc_marquee_Text)
}

// SetText sets the marquee text.
// The specified text can contain time format string sequences which are
// converted to the requested time values at runtime. Most of the time
// conversion specifiers supported by the `strftime` C function can be used.
//
//	Common time format string sequences:
//	%Y = year, %m = month, %d = day, %H = hour, %M = minute, %S = second.
//	For more information see https://en.cppreference.com/w/c/chrono/strftime.
func (m *Marquee) SetText(text string) error {
	return m.setString(C.libvlc_marquee_Text, text)
}

// Color returns the marquee text color.
// Opacity information is included in the returned color.
// Default: white.
func (m *Marquee) Color() (color.Color, error) {
	// Get color.
	rgb, err := m.getInt(C.libvlc_marquee_Color)
	if err != nil {
		return nil, err
	}

	// Get alpha.
	alpha, err := m.getInt(C.libvlc_marquee_Opacity)
	if err != nil {
		return nil, err
	}

	return color.RGBA{
		R: uint8((rgb >> 16) & 0x0ff),
		G: uint8((rgb >> 8) & 0x0ff),
		B: uint8((rgb) & 0x0ff),
		A: uint8(alpha & 0x0ff),
	}, err
}

// SetColor sets the color of the marquee text. The opacity of the text
// is also set, based on the alpha value of the color.
func (m *Marquee) SetColor(color color.Color) error {
	r, g, b, a := color.RGBA()

	// Set color.
	rgb := int((((r >> 8) & 0x0ff) << 16) |
		(((g >> 8) & 0x0ff) << 8) |
		((b >> 8) & 0x0ff))
	if err := m.setInt(C.libvlc_marquee_Color, rgb); err != nil {
		return err
	}

	// Set alpha.
	alpha := int((a >> 8) & 0x0ff)
	if err := m.setInt(C.libvlc_marquee_Opacity, alpha); err != nil {
		return err
	}

	return nil
}

// Opacity returns the opacity of the marquee text.
// The returned opacity is a value between 0 (transparent) and 255 (opaque).
// Default: 255.
func (m *Marquee) Opacity() (int, error) {
	return m.getInt(C.libvlc_marquee_Opacity)
}

// SetOpacity sets the opacity of the marquee text. The opacity is specified
// as an integer between 0 (transparent) and 255 (opaque).
func (m *Marquee) SetOpacity(opacity int) error {
	return m.setInt(C.libvlc_marquee_Opacity, opacity)
}

// Position returns the position of the marquee, relative to its container.
// Default: vlc.PositionTopLeft.
func (m *Marquee) Position() (Position, error) {
	iVal, err := m.getInt(C.libvlc_marquee_Position)
	if err != nil {
		return PositionDisable, err
	}

	switch {
	case iVal >= 8:
		iVal -= 2
	case iVal >= 4:
		iVal--
	}

	if iVal < 0 {
		return PositionTopLeft, nil
	}
	return Position(iVal), nil
}

// SetPosition sets the position of the marquee, relative to its container.
func (m *Marquee) SetPosition(position Position) error {
	switch {
	case position >= PositionBottom:
		position += 2
	case position >= PositionTop:
		position++
	}

	return m.setInt(C.libvlc_marquee_Position, int(position))
}

// X returns the X coordinate of the marquee text. The returned value is
// relative to the position of the marquee inside its container, i.e. the
// position set using the `Marquee.SetPosition` method.
// Default: 0.
func (m *Marquee) X() (int, error) {
	return m.getInt(C.libvlc_marquee_X)
}

// SetX sets the X coordinate of the marquee text. The value is specified
// relative to the position of the marquee inside its container.
//
//	NOTE: the method has no effect if the position of the marquee is set to
//	`vlc.PositionCenter`, `vlc.PositionTop` or `vlc.PositionBottom`.
func (m *Marquee) SetX(x int) error {
	return m.setInt(C.libvlc_marquee_X, x)
}

// Y returns the Y coordinate of the marquee text. The returned value is
// relative to the position of the marquee inside its container, i.e. the
// position set using the `Marquee.SetPosition` method.
// Default: 0.
func (m *Marquee) Y() (int, error) {
	return m.getInt(C.libvlc_marquee_Y)
}

// SetY sets the Y coordinate of the marquee text. The value is specified
// relative to the position of the marquee inside its container.
//
//	NOTE: the method has no effect if the position of the marquee is set to
//	`vlc.PositionCenter`, `vlc.PositionLeft` or `vlc.PositionRight`.
func (m *Marquee) SetY(y int) error {
	return m.setInt(C.libvlc_marquee_Y, y)
}

// Size returns the font size used to render the marquee text.
// Default: 0 (default font size is used).
func (m *Marquee) Size() (int, error) {
	return m.getInt(C.libvlc_marquee_Size)
}

// SetSize sets the font size used to render the marquee text.
func (m *Marquee) SetSize(size int) error {
	return m.setInt(C.libvlc_marquee_Size, size)
}

// RefreshInterval returns the interval between marquee text updates.
// The marquee text refreshes mainly when using time format string sequences.
// Default: 1s.
func (m *Marquee) RefreshInterval() (time.Duration, error) {
	iVal, err := m.getInt(C.libvlc_marquee_Refresh)
	return time.Duration(iVal) * time.Millisecond, err
}

// SetRefreshInterval sets the interval between marquee text updates.
// The marquee text refreshes mainly when using time format string sequences.
func (m *Marquee) SetRefreshInterval(refreshInterval time.Duration) error {
	return m.setInt(C.libvlc_marquee_Refresh, int(refreshInterval.Milliseconds()))
}

// DisplayDuration returns the duration for which the marquee text
// is set to be displayed.
// Default: 0 (the marquee is displayed indefinitely).
func (m *Marquee) DisplayDuration() (time.Duration, error) {
	iVal, err := m.getInt(C.libvlc_marquee_Timeout)
	return time.Duration(iVal) * time.Millisecond, err
}

// SetDisplayDuration sets the duration for which to display the marquee text.
func (m *Marquee) SetDisplayDuration(displayDuration time.Duration) error {
	return m.setInt(C.libvlc_marquee_Timeout, int(displayDuration.Milliseconds()))
}

func (m *Marquee) getInt(option C.uint) (int, error) {
	if err := m.player.assertInit(); err != nil {
		return 0, err
	}

	return int(C.libvlc_video_get_marquee_int(m.player.player, option)), nil
}

func (m *Marquee) setInt(option C.uint, val int) error {
	if err := m.player.assertInit(); err != nil {
		return err
	}

	C.libvlc_video_set_marquee_int(m.player.player, option, C.int(val))
	return nil
}

func (m *Marquee) getString(option C.uint) (string, error) {
	if err := m.player.assertInit(); err != nil {
		return "", err
	}

	cVal := C.libvlc_video_get_marquee_string(m.player.player, option)
	if cVal == nil {
		return "", errOrDefault(getError(), ErrInvalid)
	}
	defer C.free(unsafe.Pointer(cVal))

	return C.GoString(cVal), nil
}

func (m *Marquee) setString(option C.uint, val string) error {
	if err := m.player.assertInit(); err != nil {
		return err
	}

	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))

	C.libvlc_video_set_marquee_string(m.player.player, option, cVal)
	return nil
}
