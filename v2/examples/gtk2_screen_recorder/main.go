package main

/*
 * GTK screen recorder.
 * libVLC screen module must be installed.
 * See https://github.com/adrg/libvlc-go/wiki for installation instructions.
 * Uses go-gtk.
 * See https://github.com/mattn/go-gtk for installation instructions.
 */
import (
	"fmt"
	"log"

	vlc "github.com/adrg/libvlc-go/v2"
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
	window.SetResizable(false)
	window.SetTitle("libvlc-go screen recorder")
	window.SetIconName("media-record")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	// Create window container.
	container := gtk.NewVBox(false, 0)
	window.Add(container)

	// Create capture area layout.
	areaSelectRadio1 := gtk.NewRadioButtonWithLabel(nil, "Entire screen")
	areaSelectRadio2 := gtk.NewRadioButtonWithLabel(areaSelectRadio1.GetGroup(), "Select rectangle")
	areaSelectBox := gtk.NewHBox(false, 0)
	areaSelectBox.PackStart(areaSelectRadio1, false, false, 10)
	areaSelectBox.PackStart(areaSelectRadio2, false, false, 10)

	xLabel := gtk.NewLabel("X")
	xLabel.SetAlignment(0, 0)
	xInput := gtk.NewSpinButtonWithRange(0, 10000, 1)
	xBox := gtk.NewVBox(false, 0)
	xBox.PackStart(xLabel, false, false, 0)
	xBox.PackStart(xInput, false, false, 0)

	yLabel := gtk.NewLabel("Y")
	yLabel.SetAlignment(0, 0)
	yInput := gtk.NewSpinButtonWithRange(0, 10000, 1)
	yInput.SetText("0")
	yBox := gtk.NewVBox(false, 0)
	yBox.PackStart(yLabel, false, false, 0)
	yBox.PackStart(yInput, false, false, 0)

	wLabel := gtk.NewLabel("Width")
	wLabel.SetAlignment(0, 0)
	wInput := gtk.NewSpinButtonWithRange(0, 10000, 1)
	wInput.SetText("0")
	wBox := gtk.NewVBox(false, 0)
	wBox.PackStart(wLabel, false, false, 0)
	wBox.PackStart(wInput, false, false, 0)

	hLabel := gtk.NewLabel("Height")
	hLabel.SetAlignment(0, 0)
	hInput := gtk.NewSpinButtonWithRange(0, 10000, 1)
	hInput.SetText("0")
	hBox := gtk.NewVBox(false, 0)
	hBox.PackStart(hLabel, false, false, 0)
	hBox.PackStart(hInput, false, false, 0)

	areaBox := gtk.NewHBox(false, 0)
	areaBox.SetSensitive(false)
	areaBox.PackStart(xBox, false, false, 10)
	areaBox.PackStart(yBox, false, false, 10)
	areaBox.PackStart(wBox, false, false, 10)
	areaBox.PackStart(hBox, false, false, 10)
	areaBoxSpacer := gtk.NewVBox(false, 0)
	areaBoxSpacer.PackStart(areaSelectBox, false, false, 10)
	areaBoxSpacer.PackStart(areaBox, true, true, 10)

	areaFrame := gtk.NewFrame("Capture area")
	areaFrame.Add(areaBoxSpacer)
	areaSpacer := gtk.NewHBox(false, 0)
	areaSpacer.PackStart(areaFrame, true, true, 10)
	container.PackStart(areaSpacer, true, true, 10)

	// Enable disable area select controls.
	areaSelectRadio1.Connect("clicked", func(ctx *glib.CallbackContext) {
		areaBox.SetSensitive(areaSelectRadio2.GetActive())
	})

	// Create recording options layout.
	fpsLabel := gtk.NewLabel("FPS")
	fpsInput := gtk.NewSpinButtonWithRange(0, 1000, 1)
	fpsInput.SetValue(30)
	fpsBox := gtk.NewHBox(false, 0)
	fpsBox.PackStart(fpsInput, false, false, 5)
	fpsBox.PackStart(fpsLabel, false, false, 0)

	followMouseCheck := gtk.NewCheckButtonWithLabel("Follow mouse")

	recordingBox := gtk.NewHBox(false, 0)
	recordingBox.PackStart(followMouseCheck, false, false, 10)
	recordingBox.PackEnd(fpsBox, false, false, 10)
	recordingBoxSpacer := gtk.NewVBox(false, 0)
	recordingBoxSpacer.PackStart(recordingBox, true, true, 10)

	recordingFrame := gtk.NewFrame("Recording options")
	recordingFrame.Add(recordingBoxSpacer)
	recordingSpacer := gtk.NewHBox(false, 0)
	recordingSpacer.PackStart(recordingFrame, true, true, 10)
	container.PackStart(recordingSpacer, true, true, 10)

	// Create destination file layout.
	destInput := gtk.NewEntry()
	destInput.SetEditable(false)
	destButton := gtk.NewButtonWithLabel("Choose file...")
	destButton.Connect("clicked", func(ctx *glib.CallbackContext) {
		// Create file chooser dialog.
		fileDialog := gtk.NewFileChooserDialog(
			"Choose file...",
			window,
			gtk.FILE_CHOOSER_ACTION_SAVE,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)

		// Add file filter.
		fileFilter := gtk.NewFileFilter()
		fileFilter.AddPattern("*.mp4")
		fileDialog.AddFilter(fileFilter)

		fileDialog.Response(func() {
			destInput.SetText(fileDialog.GetFilename())
		})

		fileDialog.Run()
		fileDialog.Destroy()
	})

	destBox := gtk.NewHBox(false, 0)
	destBox.PackStart(destInput, true, true, 10)
	destBox.PackStart(destButton, false, false, 10)
	destBoxSpacer := gtk.NewVBox(false, 0)
	destBoxSpacer.PackStart(destBox, true, true, 10)

	destFrame := gtk.NewFrame("Destination file")
	destFrame.Add(destBoxSpacer)
	destSpacer := gtk.NewHBox(false, 0)
	destSpacer.PackStart(destFrame, true, true, 10)
	container.PackStart(destSpacer, true, true, 10)

	// Create controls layout.
	recordButton := gtk.NewButtonFromStock("gtk-media-record")
	recordButton.Connect("clicked", func(ctx *glib.CallbackContext) {
		if player.IsPlaying() {
			recordButton.SetLabel("gtk-media-record")
			areaFrame.SetSensitive(true)
			recordingFrame.SetSensitive(true)
			destFrame.SetSensitive(true)
			player.Stop()
			return
		}

		destPath := destInput.GetText()
		if destPath == "" {
			destInput.GrabFocus()
			return
		}

		// Create screen media options.
		mediaOpts := &vlc.MediaScreenOptions{
			FPS:         fpsInput.GetValue(),
			FollowMouse: followMouseCheck.GetActive(),
		}

		// Configure screen capture area.
		if areaSelectRadio2.GetActive() {
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
		areaFrame.SetSensitive(false)
		destFrame.SetSensitive(false)
		recordingFrame.SetSensitive(false)
	})

	exitButton := gtk.NewButtonFromStock("gtk-close")
	exitButton.Connect("clicked", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	ctrlBox := gtk.NewHBox(false, 0)
	ctrlBox.PackStart(recordButton, false, false, 10)
	ctrlBox.PackEnd(exitButton, false, false, 10)

	container.PackStart(ctrlBox, true, true, 10)

	window.ShowAll()
	gtk.Main()
}
