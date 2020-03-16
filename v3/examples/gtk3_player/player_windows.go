package main

/*
#cgo pkg-config: gdk-3.0
#include <gdk/gdk.h>
#include <gdk/gdkwin32.h>
*/
import "C"
import (
	"unsafe"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/gdk"
)

func setPlayerWindow(player *vlc.Player, window *gdk.Window) error {
	handle := C.gdk_win32_window_get_handle((*C.GdkWindow)(unsafe.Pointer(window.Native())))
	return player.SetHWND(uintptr(unsafe.Pointer(handle)))
}
