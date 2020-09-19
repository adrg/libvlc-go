/*
Package vlc provides golang bindings for libVLC version 2.X/3.X/4.X.

Usage

Initialization
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

Player example
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
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

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

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	// Start playing the media.
	if err = player.Play(); err != nil {
		log.Fatal(err)
	}

	<-quit

List player example
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

Handling multiple events example
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
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("test.mp3")
	// Set player media from URL:
	media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
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
*/
package vlc
