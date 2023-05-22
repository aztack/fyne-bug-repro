package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

var a fyne.App
var w fyne.Window

func main() {
	a = app.NewWithID("appId")
	w = a.NewWindow("Render Bug Reproduction")
	a.Settings().SetTheme(theme.DarkTheme())

	showDownloadUI := func() {
		downloadGui, _ := createDownloadGui()

		w.SetContent(downloadGui)
		size := w.Content().Size()
		w.Resize(fyne.NewSize(size.Width, size.Height/3))
		w.SetFixedSize(true)
		go func() {
			appGui, _ := createAppGui(w)
			w.SetFixedSize(true)
			w.SetContent(appGui)
		}()
	}

	if true {
		ctn := createListGui(func(act int, label *widget.Label) {
			showDownloadUI()
		})
		w.SetFixedSize(true)
		w.SetContent(ctn)
	} else {
		showDownloadUI()
	}
	w.ShowAndRun()
}

func createListGui(onNext func(int, *widget.Label)) *fyne.Container {

	label := widget.NewLabel("Available Versions")

	versions := []string{"1.0.0", "1.0.1"}
	list := widget.NewList(
		func() int {
			return 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("<version>")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(string(versions[i]))
		})

	downloadButton := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
		onNext(1, label)
	})
	downloadButton.Disable()
	nextButton := widget.NewButtonWithIcon("Next", theme.MediaSkipNextIcon(), func() {
		onNext(0, label)
	})

	list.OnSelected = func(id widget.ListItemID) {

	}
	expander := NewSpacer(0, 80)
	a := container.NewBorder(nil, nil, expander, nil, list)
	items := []fyne.CanvasObject{
		label, a,
	}

	items = append(items, container.NewHBox(layout.NewSpacer(), downloadButton, nextButton))
	if true {
		docLabel, err := createDocumentLink("https://fyne.io")
		if err == nil {
			items = append([]fyne.CanvasObject{docLabel}, items...)
		}
	}
	return container.NewVBox(items...)
}

func createDownloadGui() (*fyne.Container, *widget.Label) {
	downloadLabel := widget.NewLabel("Downloading Node.js and Dependencies")
	progressBar := widget.NewProgressBarInfinite()
	root := container.NewVBox(downloadLabel, progressBar)
	return root, downloadLabel
}

func createDocumentLink(u string) (*widget.Hyperlink, error) {
	_url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return widget.NewHyperlink("Document", _url), nil
}

func createAppGui(win fyne.Window) (*container.AppTabs, func()) {
	text := widget.NewMultiLineEntry()
	text.SetText("Press button 'Run' to Start")
	text.SetMinRowsVisible(16)

	// Open Asset Directory Button
	openButton := widget.NewButton("Open Asset", func() {

	})
	openButton.Icon = theme.FolderOpenIcon()

	// Buttons
	var runButton *widget.Button = nil
	var stopButton *widget.Button = nil

	onClickStopButton := func() {

	}
	stopButton = widget.NewButton("Stop", onClickStopButton)
	stopButton.Icon = theme.MediaStopIcon()
	stopButton.Disable()

	projectDirEntry := widget.NewEntry()
	// Run button
	onClickRunButton := func() {

	}
	runButton = widget.NewButton("Run", onClickRunButton)
	runButton.Icon = theme.MediaPlayIcon()

	// Exit button
	exitButton := widget.NewButton("Exit", func() {
		a.Quit()
	})
	exitButton.Icon = theme.LogoutIcon()

	environmentSelect := widget.NewSelect([]string{"prod", "test"}, func(s string) {

	})

	openLogButton := widget.NewButton("Open Log", func() {

	})

	openBrowserButton := widget.NewButton("Open Website", func() {

	})
	openBrowserButton.Icon = theme.HomeIcon()

	items1 := []fyne.CanvasObject{
		openLogButton, environmentSelect, openBrowserButton, openButton,
	}
	items2 := []fyne.CanvasObject{
		layout.NewSpacer(), runButton, stopButton, exitButton,
	}
	docLabel, err := createDocumentLink("https://fyne.io")
	if err == nil {
		items1 = append([]fyne.CanvasObject{docLabel}, items1...)
	}
	items1 = append([]fyne.CanvasObject{layout.NewSpacer()}, items1...)
	// first row of buttons
	buttons := container.NewHBox(items1...)
	// second row of buttons
	buttons2 := container.NewHBox(items2...)

	tabConsoleContent := container.NewVBox(
		container.NewMax(text),
		buttons,
		buttons2,
	)

	// Preferences tab
	root := container.NewAppTabs()
	var tab1 *container.TabItem = nil
	var tab2 *container.TabItem = nil
	var applyButton *widget.Button = nil

	projectDirEntry.SetText("/Users")
	projectDirEntry.OnChanged = func(selectedPath string) {

	}
	row1 := container.NewBorder(nil, nil, nil, container.NewHBox(widget.NewButton("...", func() {

	}), widget.NewButton("Reset", func() {

	})), projectDirEntry)

	f := container.New(layout.NewFormLayout(), widget.NewLabel("Project Dir"), row1)

	applyButton = widget.NewButton("Apply", func() {
	})
	applyButton.Disable()
	applyButton.SetIcon(theme.DocumentSaveIcon())
	prefButtons := container.NewHBox(layout.NewSpacer(), applyButton)
	tabPreferencesContent := container.NewMax(
		container.NewPadded(container.NewBorder(nil, prefButtons, nil, nil, f)),
	)

	tab1 = container.NewTabItem("Console", tabConsoleContent)
	tab2 = container.NewTabItem("Preferences", tabPreferencesContent)
	root.Append(tab1)
	root.Append(tab2)

	return root, onClickRunButton
}
