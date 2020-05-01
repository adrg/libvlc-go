package main

import (
	"fmt"
	"log"
	"os"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.libvlc-go.gtk3-screen-recorder-example"

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

		// Get capture area frame controls.
		captureAreaFrame, ok := builderGetObject(builder, "captureAreaFrame").(*gtk.Frame)
		assertConv(ok)

		entireScreenRadio, ok := builderGetObject(builder, "radioEntireScreen").(*gtk.RadioButton)
		assertConv(ok)

		xInput, ok := builderGetObject(builder, "xInput").(*gtk.SpinButton)
		assertConv(ok)

		yInput, ok := builderGetObject(builder, "yInput").(*gtk.SpinButton)
		assertConv(ok)

		wInput, ok := builderGetObject(builder, "widthInput").(*gtk.SpinButton)
		assertConv(ok)

		hInput, ok := builderGetObject(builder, "heightInput").(*gtk.SpinButton)
		assertConv(ok)

		areaRectBox, ok := builderGetObject(builder, "rectangleAreaBox").(*gtk.Box)
		assertConv(ok)

		// Get recording options frame controls.
		recOptionsFrame, ok := builderGetObject(builder, "recordingOptionsFrame").(*gtk.Frame)
		assertConv(ok)

		followMouseCheck, ok := builderGetObject(builder, "followMouseCheck").(*gtk.CheckButton)
		assertConv(ok)

		fpsInput, ok := builderGetObject(builder, "fpsInput").(*gtk.SpinButton)
		assertConv(ok)

		// Get destination file frame controls.
		destFileFrame, ok := builderGetObject(builder, "destinationFileFrame").(*gtk.Frame)
		assertConv(ok)

		destInput, ok := builderGetObject(builder, "destinationInput").(*gtk.Entry)
		assertConv(ok)

		// Add builder signal handlers.
		signals := map[string]interface{}{
			"onClickAreaSelect": func() {
				areaRectBox.SetSensitive(!entireScreenRadio.GetActive())
			},
			"onClickChooseDestinationFile": func() {
				fileDialog, err := gtk.FileChooserDialogNewWith2Buttons(
					"Choose file...",
					appWin, gtk.FILE_CHOOSER_ACTION_SAVE,
					"Cancel", gtk.RESPONSE_DELETE_EVENT,
					"Save", gtk.RESPONSE_ACCEPT)
				assertErr(err)
				defer fileDialog.Destroy()

				fileFilter, err := gtk.FileFilterNew()
				assertErr(err)
				fileFilter.SetName("Media files")
				fileFilter.AddPattern("*.mp4")
				fileDialog.AddFilter(fileFilter)

				if result := fileDialog.Run(); result == gtk.RESPONSE_ACCEPT {
					destInput.SetText(fileDialog.GetFilename())
				}
			},
			"onClickRecord": func(recordButton *gtk.Button) {
				if player.IsPlaying() {
					recordButton.SetLabel("gtk-media-record")
					captureAreaFrame.SetSensitive(true)
					recOptionsFrame.SetSensitive(true)
					destFileFrame.SetSensitive(true)
					player.Stop()
					return
				}

				destPath, _ := destInput.GetText()
				if destPath == "" {
					destInput.GrabFocus()
					return
				}

				// Create screen media options.
				mediaOpts := &vlc.MediaScreenOptions{
					FPS:         fpsInput.GetValue(),
					FollowMouse: followMouseCheck.GetActive(),
				}
				if mediaOpts.FPS <= 0 {
					mediaOpts.FPS = 30
					fpsInput.SetValue(mediaOpts.FPS)
				}

				// Configure screen capture area.
				if !entireScreenRadio.GetActive() {
					mediaOpts.X = int(xInput.GetValue())
					mediaOpts.Y = int(yInput.GetValue())
					mediaOpts.Width = int(wInput.GetValue())
					mediaOpts.Height = int(hInput.GetValue())
				}

				// Create media from screen.
				media, err := vlc.NewMediaFromScreen(mediaOpts)
				if err != nil {
					log.Fatalf("Cannot load screen media: %s\n", err)
				}

				// Configure media to save the recording to the selected destination path.
				saveOpt := fmt.Sprintf(":sout=#transcode{vcodec=h264,vb=0,scale=1}:duplicate{dst=file{dst=%s}}", destPath)
				if err := media.AddOptions(saveOpt); err != nil {
					log.Fatalf("Cannot add media options: %s\n", err)
				}

				// Set player media.
				if err := player.SetMedia(media); err != nil {
					log.Fatal("Cannot set player media: %s\n", err)
				}

				// Start screen recording.
				if err := player.Play(); err != nil {
					log.Fatal("Cannot play media: %s\n", err)
				}

				recordButton.SetLabel("gtk-media-stop")
				captureAreaFrame.SetSensitive(false)
				recOptionsFrame.SetSensitive(false)
				destFileFrame.SetSensitive(false)
			},
			"onClickClose": func() {
				app.Quit()
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
