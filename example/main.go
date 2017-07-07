package main

import (
	"log"
	"time"

	vlc "github.com/tarrsalah/libvlc-go"
)

func main() {
	// Initialize libvlc. Additional command line arguments can be passed in
	// to libvlc by specifying them in the Init function.
	if err := vlc.Init("--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new player
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Add a media file from path or from URL.
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	media, err := player.LoadMediaFromPath("./sample.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

	// Play
	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}

	// Wait some amount of time for the media to start playing.
	// Depends on the version of libvlc. From my tests, libvlc 3.X does not
	// need this delay.
	// TODO: Implement proper callbacks for getting the state of the media.
	time.Sleep(1 * time.Second)

	eventManager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	callback := func(event *vlc.Event, userData interface{}) {
		log.Println(userData)
	}
	eventManager.Attach(vlc.MediaPlayerPaused, callback, "userData passed to the callback")
	time.Sleep(2 * time.Second)
	player.TogglePause()

	// If the media played is a live stream the length will be 0
	length, err := player.MediaLength()
	if err != nil || length == 0 {
		length = 1000 * 60
	}

	time.Sleep(time.Duration(length) * time.Millisecond)
}
