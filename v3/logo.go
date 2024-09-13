package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"time"
	"unsafe"
)

// LogoFile represents a logo file which can be used as a media player's logo.
// The logo of a player can also be composed of a series of alternating files.
type LogoFile struct {
	path            string
	displayDuration time.Duration
	opacity         int
}

// NewLogoFileFromPath returns a new logo file with the specified path.
// The file is displyed for the provided duration. If the specified display
// duration is negative, the global display duration set on the logo the file
// is applied to is used.
// The provided opacity must be a value between 0 (transparent) and 255 (opaque).
// If the specified opacity is negative, the global opacity set on the logo
// the file is applied to is used.
func NewLogoFileFromPath(path string, displayDuration time.Duration, opacity int) (*LogoFile, error) {
	if path == "" {
		return nil, ErrInvalid
	}

	return &LogoFile{
		path:            path,
		displayDuration: displayDuration,
		opacity:         opacity,
	}, nil
}

// NewLogoFileFromImage returns a new logo file with the specified image.
// The file is displyed for the provided duration. If the specified display
// duration is negative, the global display duration set on the logo the file
// is applied to is used.
// The provided opacity must be a value between 0 (transparent) and 255 (opaque).
// If the specified opacity is negative, the global opacity set on the logo
// the file is applied to is used.
func NewLogoFileFromImage(img image.Image, displayDuration time.Duration, opacity int) (*LogoFile, error) {
	if img == nil {
		return nil, ErrInvalid
	}

	// Create temporary file.
	f, err := os.CreateTemp("", "logo_*.png")
	if err != nil {
		return nil, err
	}

	// Encode logo image to temporary file.
	if err := png.Encode(f, img); err != nil {
		return nil, err
	}

	// Close temporary file.
	path := f.Name()
	if err := f.Close(); err != nil {
		return nil, err
	}

	return NewLogoFileFromPath(path, displayDuration, opacity)
}

// Logo represents a logo that can be displayed over a media instance.
//
//	For more information see https://wiki.videolan.org/Documentation:Modules/logo.
type Logo struct {
	player *Player
}

func newLogo(player *Player) *Logo {
	return &Logo{
		player: player,
	}
}

// Enable enables or disables the logo. By default, the logo is disabled.
func (l *Logo) Enable(enable bool) error {
	return l.setInt(C.libvlc_logo_enable, boolToInt(enable))
}

// SetFiles sets the sequence of files to be displayed for the logo.
func (l *Logo) SetFiles(files ...*LogoFile) error {
	fileFmts := make([]string, 0, len(files))
	for _, file := range files {
		if file == nil {
			continue
		}
		fileFmt := file.path
		if file.displayDuration >= 0 {
			fileFmt = fmt.Sprintf("%s,%d", fileFmt, file.displayDuration.Milliseconds())
		}
		if file.opacity >= 0 {
			fileFmt = fmt.Sprintf("%s,%d", fileFmt, file.opacity)
		}

		fileFmts = append(fileFmts, fileFmt)
	}

	if len(fileFmts) == 0 {
		return nil
	}

	return l.setString(C.libvlc_logo_file, strings.Join(fileFmts, ";"))
}

// Position returns the position of the logo, relative to its container.
// Default: vlc.PositionTopLeft.
func (l *Logo) Position() (Position, error) {
	iVal, err := l.getInt(C.libvlc_logo_position)
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

// SetPosition sets the position of the logo, relative to its container.
func (l *Logo) SetPosition(position Position) error {
	switch {
	case position >= PositionBottom:
		position += 2
	case position >= PositionTop:
		position++
	}

	return l.setInt(C.libvlc_logo_position, int(position))
}

// X returns the X coordinate of the logo. The returned value is
// relative to the position of the logo inside its container, i.e. the
// position set using Logo.SetPosition method.
// Default: 0.
func (l *Logo) X() (int, error) {
	return l.getInt(C.libvlc_logo_x)
}

// SetX sets the X coordinate of the logo. The value is specified
// relative to the position of the logo inside its container, i.e. the
// position set using the `Logo.SetPosition` method.
//
//	NOTE: the method has no effect if the position of the logo is set to
//	`vlc.PositionCenter`, `vlc.PositionTop` or `vlc.PositionBottom`.
func (l *Logo) SetX(x int) error {
	return l.setInt(C.libvlc_logo_x, x)
}

// Y returns the Y coordinate of the logo. The returned value is
// relative to the position of the logo inside its container, i.e. the
// position set using the `Logo.SetPosition` method.
// Default: 0.
func (l *Logo) Y() (int, error) {
	return l.getInt(C.libvlc_logo_y)
}

// SetY sets the Y coordinate of the logo. The value is specified
// relative to the position of the logo inside its container.
//
//	NOTE: the method has no effect if the position of the logo is set to
//	`vlc.PositionCenter`, `vlc.PositionLeft` or `vlc.PositionRight`.
func (l *Logo) SetY(y int) error {
	return l.setInt(C.libvlc_logo_y, y)
}

// Opacity returns the global opacity of the logo.
// The returned opacity is a value between 0 (transparent) and 255 (opaque).
// The global opacity can be overridden by each provided logo file.
// Default: 255.
func (l *Logo) Opacity() (int, error) {
	return l.getInt(C.libvlc_logo_opacity)
}

// SetOpacity sets the global opacity of the logo. If an opacity override is
// not specified when setting the logo files, the global opacity is used. The
// opacity is specified as an integer between 0 (transparent) and 255 (opaque).
// The global opacity can be overridden by each provided logo file.
func (l *Logo) SetOpacity(opacity int) error {
	return l.setInt(C.libvlc_logo_opacity, opacity)
}

// DisplayDuration returns the global duration for which a logo file
// is set to be displayed before displaying the next one (if one is available).
// The global display duration can be overridden by each provided logo file.
// Default: 1s.
func (l *Logo) DisplayDuration() (time.Duration, error) {
	iVal, err := l.getInt(C.libvlc_logo_delay)
	return time.Duration(iVal) * time.Millisecond, err
}

// SetDisplayDuration sets the duration for which to display a logo file
// before displaying the next one (if one is available).
// The global display duration can be overridden by each provided logo file.
func (l *Logo) SetDisplayDuration(displayDuration time.Duration) error {
	return l.setInt(C.libvlc_logo_delay, int(displayDuration.Milliseconds()))
}

// RepeatCount returns the number of times the logo sequence is set
// to be repeated.
func (l *Logo) RepeatCount() (int, error) {
	return l.getInt(C.libvlc_logo_repeat)
}

// SetRepeatCount sets the number of times the logo sequence should repeat.
// Pass in `-1` to repeat the logo sequence indefinitely, `0` to disable logo
// sequence looping or a positive number to repeat the logo sequence a specific
// number of times.
// Default: -1 (the logo sequence is repeated indefinitely).
func (l *Logo) SetRepeatCount(count int) error {
	if count > 0 {
		count++
	}

	return l.setInt(C.libvlc_logo_repeat, count)
}

func (l *Logo) getInt(option C.uint) (int, error) {
	if err := l.player.assertInit(); err != nil {
		return 0, err
	}

	return int(C.libvlc_video_get_logo_int(l.player.player, option)), nil
}

func (l *Logo) setInt(option C.uint, val int) error {
	if err := l.player.assertInit(); err != nil {
		return err
	}

	C.libvlc_video_set_logo_int(l.player.player, option, C.int(val))
	return nil
}

func (l *Logo) setString(option C.uint, val string) error {
	if err := l.player.assertInit(); err != nil {
		return err
	}

	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))

	C.libvlc_video_set_logo_string(l.player.player, option, cVal)
	return nil
}
