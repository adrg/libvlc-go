package main

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v2"
)

func main() {
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new list player.
	player, err := vlc.NewListPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Create a new media list.
	list, err := vlc.NewMediaList()
	if err != nil {
		log.Fatal(err)
	}
	defer list.Release()

	err = list.AddMediaFromPath("localpath/test1.mp3")
	if err != nil {
		log.Fatal(err)
	}

	err = list.AddMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
	if err != nil {
		log.Fatal(err)
	}

	// Set player media list.
	if err = player.SetMediaList(list); err != nil {
		log.Fatal(err)
	}

	// Media files can be added to the list after the list has been added
	// to the player. The player will play these files as well.
	err = list.AddMediaFromPath("localpath/test2.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	// Register the media end reached event with the event manager.
	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}) {
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaListPlayerPlayed, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	// Start playing the media list.
	if err = player.Play(); err != nil {
		log.Fatal(err)
	}

	<-quit
}
