package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

var (
	obsubsApp fyne.App
	window    fyne.Window
)

func init() {

}

func main() {
	obsubsApp = app.New()
	window = obsubsApp.NewWindow("OBS Subtitles Controller")

	newItem := fyne.NewMenuItem("Open subs", OpenFileHandler)
	newMenu := fyne.NewMenu("File", newItem)
	mainMenu := fyne.NewMainMenu(newMenu)
	window.SetMainMenu(mainMenu)
	window.SetMaster()
	window.Resize(fyne.NewSize(1000, 1000))

	window.ShowAndRun()
}

func readFile(f fyne.URIReadCloser, err error) {
	if err != nil {
		dialog.ShowError(err, window)
	}

	if f == nil {
		dialog.ShowError(fmt.Errorf("Unable to open file"), window)
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	if data == nil {
		dialog.ShowError(fmt.Errorf("No data found"), window)
	}

	log.Println(string(data))

	err = f.Close()
	if err != nil {
		dialog.ShowError(err, window)
	}
}

func OpenFileHandler() {
	fd := dialog.NewFileOpen(readFile, window)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
	fd.Show()
}
