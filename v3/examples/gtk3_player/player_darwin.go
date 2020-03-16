package main

/*
#cgo pkg-config: gdk-3.0
#include <gdk/gdk.h>

GDK_AVAILABLE_IN_ALL NSWindow* gdk_quartz_window_get_nswindow(GdkWindow *window);

void* getNsWindow(GdkWindow *w) {
	return (void*)gdk_quartz_window_get_nswindow(w);
}
*/
import "C"
import (
	"unsafe"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/gdk"
)

type NSWindow struct {
	ID unsafe.Pointer
}

func setPlayerWindow(player *vlc.Player, window *gdk.Window) error {
	handle := &NSWindow{
		unsafe.Pointer(C.getNsWindow((*C.GdkWindow)(unsafe.Pointer(window.Native())))),
	}
	return player.SetNSObject(uintptr(unsafe.Pointer(handle)))
}
