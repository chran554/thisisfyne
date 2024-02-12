package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"thisisfyne/internal/app/selfie"
	w "thisisfyne/internal/pkg/widget"
)

var currentSelfieSet *selfie.SelfieSet
var selfies []*selfie.SelfieSet
var mainArea *fyne.Container
var selfieSelectionListWidget *widget.List

func init() {
	mainArea = container.New(layout.NewStackLayout())
}

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("This is fyne")

	toolbarWidget := toolbar(window)

	imageFiles, err := selfie.LoadImageFiles("images/selfies")
	if err != nil {
		log.Println("Could not preload images from 'images/selfies', use 'open directory' to load images manually.")
	}
	selfies = selfie.ConvertToSelfieSets(imageFiles)

	selfieSelectionListWidget = selfieSetWidgetList()

	mainLayout := container.NewBorder(toolbarWidget, nil, selfieSelectionListWidget, nil, mainArea)
	updateMainArea()

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(mainLayout)
	window.ShowAndRun()
}

func selfieSetWidgetList() *widget.List {
	imageList := widget.NewList(
		func() int {
			return len(selfies)
		},
		func() fyne.CanvasObject {
			fmt.Println("Create")
			labelledImage := w.NewSelfieTreeItem(nil, 150)
			return labelledImage
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			fmt.Printf("Update %d: %s\n", i, selfies[i].Name)
			o.(*w.SelfieTreeItem).SetSelfies(selfies[i])
		})

	imageList.OnSelected = func(i widget.ListItemID) {
		fmt.Printf("Selected %d: %s\n", i, selfies[i].Name)
		currentSelfieSet = selfies[i]
		updateMainArea()
	}

	imageList.OnUnselected = func(i widget.ListItemID) {
		fmt.Printf("Unselected %d: %s\n", i, selfies[i].Name)
	}

	return imageList
}

func updateMainArea() {
	if currentSelfieSet != nil {
		setMainAreaSelfies(currentSelfieSet)
	} else {
		setPrimaryAreaEmpty()
	}
}

func setMainAreaSelfies(selfies *selfie.SelfieSet) {
	if selfies != nil {
		primaryArea := container.NewVBox()
		secondaryArea := container.NewHBox()

		primarySelfie := selfies.PrimaryImage
		primaryImage := canvas.NewImageFromImage(*primarySelfie.Image)
		primaryImage.FillMode = canvas.ImageFillOriginal
		primaryImage.ScaleMode = canvas.ImageScaleSmooth

		primaryInfoImage := w.NewInfoImage(primarySelfie.Image, 300, primarySelfie.FileName, primarySelfie.Info1, primarySelfie.Info2, primarySelfie.Info3, primarySelfie.Info4, primarySelfie.Info5)

		statusButtons := w.NewStatusSelectButtons(selfies.Status, func(status selfie.SelfieSetStatus) {
			selfies.Status = status
			selfieSelectionListWidget.Refresh()
		})

		primaryArea.Add(primaryInfoImage)
		primaryArea.Add(container.NewPadded(statusButtons))

		for _, secondarySelfie := range selfies.SecondaryImages {
			secondaryImage := canvas.NewImageFromImage(*secondarySelfie.Image)
			secondaryImage.FillMode = canvas.ImageFillOriginal
			secondaryImage.ScaleMode = canvas.ImageScaleSmooth

			secondaryInfoImage := w.NewInfoImage(secondarySelfie.Image, 300, secondarySelfie.FileName, secondarySelfie.Info1, secondarySelfie.Info2, secondarySelfie.Info3, secondarySelfie.Info4, secondarySelfie.Info5)
			secondaryArea.Add(secondaryInfoImage)
		}

		secondaryAreaScroll := container.NewHScroll(secondaryArea)

		area := container.NewBorder(nil, nil, primaryArea, nil, secondaryAreaScroll)

		mainArea.RemoveAll()
		mainArea.Add(area)
		mainArea.Refresh()
	}
}

func setPrimaryAreaEmpty() {
	mainArea.RemoveAll()

	mainImage := canvas.NewImageFromFile("documentation/this_is_fyne.jpg")
	mainImage.ScaleMode = canvas.ImageScaleSmooth
	mainImage.FillMode = canvas.ImageFillOriginal

	mainArea.Add(mainImage)
	mainArea.Refresh()
}

func toolbar(parent fyne.Window) *widget.Toolbar {
	toolbarWidget := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			loadSelfies(parent)
		}),
		widget.NewToolbarAction(theme.DownloadIcon(), func() {
			exportSelfiesStatus(parent)
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MoveUpIcon(), func() {}),
		widget.NewToolbarAction(theme.MoveDownIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	return toolbarWidget
}

func loadSelfies(parent fyne.Window) {
	folderOpenDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri != nil {
			path := uri.Path()
			imageFiles, err := selfie.LoadImageFiles(path)
			if err != nil {
				panic(err)
			}
			selfies = selfie.ConvertToSelfieSets(imageFiles)
			currentSelfieSet = nil

			setMainAreaSelfies(nil)
			selfieSelectionListWidget.UnselectAll()
			selfieSelectionListWidget.Refresh()
		}
	}, parent)

	directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
	if err != nil {
		log.Fatal(err) //print the error if obtained
	}

	mfileURI := storage.NewFileURI(directory)
	mfileLister, _ := storage.ListerForURI(mfileURI)
	folderOpenDialog.SetLocation(mfileLister)

	folderOpenDialog.Show()
}

func exportSelfiesStatus(parent fyne.Window) {
	jsonData, err := json.MarshalIndent(selfies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fileSaveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer != nil {
			_, err = writer.Write(jsonData)
		}
	}, parent)

	fileSaveDialog.SetFileName("selfies-report.json")
	fileSaveDialog.SetFilter(storage.NewExtensionFileFilter([]string{"json"}))

	fileSaveDialog.Show()
	if err != nil {
		log.Fatal(err)
	}
}
