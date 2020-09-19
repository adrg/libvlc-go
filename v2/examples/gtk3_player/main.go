package main

import (
	"log"
	"os"

	vlc "github.com/adrg/libvlc-go/v2"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.libvlc-go.gtk3-media-player-example"

func builderGetObject(builder *gtk.Builder, name string) glib.IObject {
	obj, _ := builder.GetObject(name)
	return obj
}

func assertErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func assertConv(ok bool) {
	if !ok {
		log.Panic("invalid widget conversion")
	}
}

func playerReleaseMedia(player *vlc.Player) {
	player.Stop()
	if media, _ := player.Media(); media != nil {
		media.Release()
	}
}

func main() {
	// Initialize libVLC module.
	err := vlc.Init("--quiet", "--no-xlib")
	assertErr(err)

	// Create a new player.
	player, err := vlc.NewPlayer()
	assertErr(err)

	// Create new GTK application.
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	assertErr(err)

	app.Connect("activate", func() {
		// Load application layout.
		builder, err := gtk.BuilderNewFromFile("layout.glade")
		assertErr(err)

		// Get application window.
		appWin, ok := builderGetObject(builder, "appWindow").(*gtk.ApplicationWindow)
		assertConv(ok)

		// Get play button.
		playButton, ok := builderGetObject(builder, "playButton").(*gtk.Button)
		assertConv(ok)

		// Add builder signal handlers.
		signals := map[string]interface{}{
			"onRealizePlayerArea": func(playerArea *gtk.DrawingArea) {
				// Set window for the player.
				playerWindow, err := playerArea.GetWindow()
				assertErr(err)
				err = setPlayerWindow(player, playerWindow)
				assertErr(err)
			},
			"onDrawPlayerArea": func(playerArea *gtk.DrawingArea, cr *cairo.Context) {
				cr.SetSourceRGB(0, 0, 0)
				cr.Paint()
			},
			"onActivateOpenFile": func() {
				fileDialog, err := gtk.FileChooserDialogNewWith2Buttons(
					"Choose file...",
					appWin, gtk.FILE_CHOOSER_ACTION_OPEN,
					"Cancel", gtk.RESPONSE_DELETE_EVENT,
					"Open", gtk.RESPONSE_ACCEPT)
				assertErr(err)
				defer fileDialog.Destroy()

				fileFilter, err := gtk.FileFilterNew()
				assertErr(err)
				fileFilter.SetName("Media files")
				fileFilter.AddPattern("*.mp4")
				fileFilter.AddPattern("*.mp3")
				fileDialog.AddFilter(fileFilter)

				if result := fileDialog.Run(); result == gtk.RESPONSE_ACCEPT {
					// Release current media, if any.
					playerReleaseMedia(player)

					// Get selected filename.
					filename := fileDialog.GetFilename()

					// Load media and start playback.
					if _, err := player.LoadMediaFromPath(filename); err != nil {
						log.Printf("Cannot load selected media: %s\n", err)
						return
					}

					player.Play()
					playButton.SetLabel("gtk-media-pause")
				}
			},
			"onActivateQuit": func() {
				app.Quit()
			},
			"onClickPlayButton": func(playButton *gtk.Button) {
				if media, _ := player.Media(); media == nil {
					return
				}

				if player.IsPlaying() {
					player.SetPause(true)
					playButton.SetLabel("gtk-media-play")
				} else {
					player.Play()
					playButton.SetLabel("gtk-media-pause")
				}
			},
			"onClickStopButton": func(stopButton *gtk.Button) {
				player.Stop()
				playButton.SetLabel("gtk-media-play")
			},
		}
		builder.ConnectSignals(signals)

		appWin.ShowAll()
		app.AddWindow(appWin)
	})

	// Cleanup on exit.
	app.Connect("shutdown", func() {
		playerReleaseMedia(player)
		player.Release()
		vlc.Release()
	})

	// Launch the application.
	os.Exit(app.Run(os.Args))
}
