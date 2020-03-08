package main

/*
 * Sample GTK player.
 * See https://github.com/mattn/go-gtk for go-gtk installation instructions.
 */
import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	// Initialize libVLC module.
	if err := vlc.Init("--quiet", "--no-xlib"); err != nil {
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
		if media, _ := player.Media(); media != nil {
			media.Release()
		}

		player.Release()
	}()

	// Create GTK window.
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTypeHint(gdk.WINDOW_TYPE_HINT_DIALOG)
	window.SetDefaultSize(800, 600)
	window.SetTitle("libvlc-go player")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	// Create window container.
	vBox := gtk.NewVBox(false, 0)
	window.Add(vBox)

	// Add player file menu layout.
	openItem := gtk.NewMenuItemWithLabel("Open")
	exitItem := gtk.NewMenuItemWithLabel("Exit")
	fileMenu := gtk.NewMenu()
	fileMenu.Append(openItem)
	fileMenu.Append(exitItem)

	fileItem := gtk.NewMenuItemWithLabel("File")
	fileItem.SetSubmenu(fileMenu)
	menuBar := gtk.NewMenuBar()
	menuBar.Append(fileItem)
	vBox.PackStart(menuBar, false, false, 0)

	// Player video area layout.
	playerWidget := gtk.NewDrawingArea()
	vBox.PackStart(playerWidget, true, true, 0)

	// Wait for the realize event and attach the player to the widget window ID.
	playerWidget.Connect("realize", func(ctx *glib.CallbackContext) {
		// Set window for the player.
		player.SetXWindow(uint32(playerWidget.GetWindow().GetNativeWindowID()))
	})

	// Player controls area layout.
	playButton := gtk.NewButtonFromStock("gtk-media-play")
	stopButton := gtk.NewButtonFromStock("gtk-media-stop")

	hBox := gtk.NewHBox(false, 0)
	hBox.PackStart(playButton, false, false, 0)
	hBox.PackStart(stopButton, false, false, 0)
	vBox.PackStart(hBox, false, false, 0)

	// File open menu item event callback.
	openItem.Connect("activate", func(ctx *glib.CallbackContext) {
		// Create file chooser dialog.
		fileDialog := gtk.NewFileChooserDialog(
			"Choose file...",
			window,
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)

		// Add file filter.
		fileFilter := gtk.NewFileFilter()
		fileFilter.AddPattern("*.mp4")
		fileFilter.AddPattern("*.mp3")
		fileDialog.AddFilter(fileFilter)

		fileDialog.Response(func() {
			// Release current media.
			if media, _ := player.Media(); media != nil {
				media.Release()
			}
			player.Stop()

			// Get selected filename.
			filename := fileDialog.GetFilename()
			fileDialog.Destroy()

			// Load media and start playback.
			if _, err := player.LoadMediaFromPath(filename); err != nil {
				log.Printf("Cannot load selected media: %s\n", err)
				return
			}

			player.Play()
			playButton.SetLabel("gtk-media-pause")
		})

		// Open file dialog.
		fileDialog.Run()
	})

	// Exit menu item event callback.
	exitItem.Connect("activate", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	// Play/Pause button event callback.
	playButton.Connect("clicked", func(ctx *glib.CallbackContext) {
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
	})

	// Stop button event callback.
	stopButton.Connect("clicked", func(ctx *glib.CallbackContext) {
		player.Stop()
		playButton.SetLabel("gtk-media-play")
	})

	window.ShowAll()
	gtk.Main()
}
