package main

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
)

func main() {
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new player.
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Add a media file from path or from URL.
	// Set player media from URL:
	// media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
	// Set player media from path:
	media, err := player.LoadMediaFromPath("test.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	// Create event handler.
	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}) {
		switch event {
		case vlc.MediaPlayerEndReached:
			log.Println("Player end reached")
			close(quit)
		case vlc.MediaPlayerTimeChanged:
			media, err := player.Media()
			if err != nil {
				log.Println(err)
				break
			}

			stats, err := media.Stats()
			if err != nil {
				log.Println(err)
				break
			}

			log.Printf("%+v\n", stats)
		}
	}

	// Register events with the event manager.
	events := []vlc.Event{
		vlc.MediaPlayerTimeChanged,
		vlc.MediaPlayerEndReached,
	}

	var eventIDs []vlc.EventID
	for _, event := range events {
		eventID, err := manager.Attach(event, eventCallback, nil)
		if err != nil {
			log.Fatal(err)
		}

		eventIDs = append(eventIDs, eventID)
	}

	// De-register attached events.
	defer func() {
		for _, eventID := range eventIDs {
			manager.Detach(eventID)
		}
	}()

	// Start playing the media.
	if err = player.Play(); err != nil {
		log.Fatal(err)
	}

	<-quit
}
