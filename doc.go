/*
Package libvlc-go Provides golang bindings for libvlc version 2.X/3.X.

Usage

Initialization
	// Initialize libvlc. Additional command line arguments can be passed in
	// to libvlc by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()


Player example
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
	media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
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

	// If the media played is a live stream the length will be 0
	length, err := player.MediaLength()
	if err != nil || length == 0 {
		length = 1000 * 60
	}

	time.Sleep(time.Duration(length) * time.Millisecond)

List player example
    // Create a new list player
    player, err := vlc.NewListPlayer()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        player.Stop()
        player.Release()
    }()

    // Create a new media list
    list, err := vlc.NewMediaList()
    if err != nil {
        log.Fatal(err)
    }
    defer list.Release()

    err = list.AddMediaFromPath("localpath/example1.mp3")
    if err != nil {
        log.Fatal(err)
    }

    err = list.AddMediaFromURL("http://example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Set player media list
    err = player.SetMediaList(list)
    if err != nil {
        log.Fatal(err)
    }

    // Media files can be added to the list after the list has been added
    // to the player. The player will play these files as well.
    err = list.AddMediaFromPath("localpath/example2.mp3")
    if err != nil {
        log.Fatal(err)
    }

    // Wait some amount of time for the media to start playing.
    // Depends on the version of libvlc. From my tests, libvlc 3.X does not
    // need this delay.
    // TODO: Implement proper callbacks for getting the state of the media.
    time.Sleep(1 * time.Second)

    // Play
    err = player.Play()
    if err != nil {
        log.Fatal(err)
    }

    time.Sleep(60 * 1000 * time.Millisecond)
*/
package vlc
