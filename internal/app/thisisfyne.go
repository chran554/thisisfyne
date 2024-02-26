package app

import (
	"encoding/json"
	"fyne.io/fyne/v2"
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
	"thisisfyne/internal/app/resources"
	"thisisfyne/internal/app/selfie"
	w "thisisfyne/internal/pkg/widget"
)

var selectedSelfieSetIndex int
var selfieSets []*selfie.SelfieSet

var mainArea *fyne.Container
var selfieSelectionListWidget *w.SelectionList
var statusButtons *w.SelfieSetStatusButtons
var secondaryAreaScroll *container.Scroll

func init() {
	selectedSelfieSetIndex = -1
	mainArea = container.New(layout.NewStackLayout())
}

func ApplicationContent(window fyne.Window) *fyne.Container {
	window.Canvas().SetOnTypedKey(keyTypeHandler())

	toolbarWidget := toolbar(window)

	// Make an initial attempt to load images from "default" subdirectory.
	imageFiles, err := selfie.LoadImageFiles("images/selfies")
	if err != nil {
		log.Println("Could not preload images from 'images/selfies', use 'open directory' to load images manually.")
	}
	setSelfieSets(selfie.ConvertToSelfieSets(imageFiles))

	selfieSelectionListWidget = selfieSetSelectionListWidget()

	mainLayout := container.NewBorder(toolbarWidget, nil, selfieSelectionListWidget, nil, mainArea)
	updateMainArea()
	return mainLayout
}

func selfieSetSelectionListWidget() *w.SelectionList {
	imageList := w.NewSelectionList(
		func() int {
			return len(selfieSets)
		},
		func() fyne.CanvasObject {
			labelledImage := w.NewSelfieSetTreeItem(150)
			return labelledImage
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*w.SelfieSetTreeItem).SetSelfieSet(selfieSets[i])
		})

	imageList.OnSelected = func(i widget.ListItemID) {
		selectedSelfieSetIndex = i
		updateMainArea()
	}

	return imageList
}

func updateMainArea() {
	if getSelectedSelfieSet() != nil {
		setMainAreaSelfies(getSelectedSelfieSet())
	} else {
		setPrimaryAreaEmpty()
	}
}

func setMainAreaSelfies(selfieSet *selfie.SelfieSet) {
	if selfieSet != nil {
		primaryArea := container.NewVBox()
		secondaryArea := container.NewHBox()

		primarySelfie := selfieSet.PrimaryImage
		primaryInfoImage := w.NewInfoImage(primarySelfie.Image, 300, primarySelfie.FileName, primarySelfie.Info1, primarySelfie.Info2, primarySelfie.Info3, primarySelfie.Info4, primarySelfie.Info5)

		statusButtons = w.NewSelfieSetStatusButtons(func(status selfie.SelfieSetStatus) {
			setSelfieSetStatus(status)
		})
		statusButtons.SetSelfieSetStatus(selfieSet.Status)

		primaryArea.Add(primaryInfoImage)
		primaryArea.Add(container.NewPadded(statusButtons))

		for _, secondarySelfie := range selfieSet.SecondaryImages {
			secondaryInfoImage := w.NewInfoImage(secondarySelfie.Image, 300, secondarySelfie.FileName, secondarySelfie.Info1, secondarySelfie.Info2, secondarySelfie.Info3, secondarySelfie.Info4, secondarySelfie.Info5)
			secondaryArea.Add(secondaryInfoImage)
		}

		secondaryAreaScroll = container.NewHScroll(secondaryArea)

		area := container.NewBorder(nil, nil, primaryArea, nil, secondaryAreaScroll)

		mainArea.RemoveAll()
		mainArea.Add(area)
		mainArea.Refresh()
	} else {
		mainArea.RemoveAll()
		mainArea.Refresh()
	}
}

func setSelfieSetStatus(status selfie.SelfieSetStatus) {
	selfieSet := getSelectedSelfieSet()

	if selfieSet != nil {
		selfieSet.Status = status
		selfieSelectionListWidget.RefreshItem(selectedSelfieSetIndex)
		statusButtons.SetSelfieSetStatus(status)
	}
}

func setPrimaryAreaEmpty() {
	mainArea.RemoveAll()

	mainImage := canvas.NewImageFromResource(resources.ThisIsFyneImageResource)
	mainImage.ScaleMode = canvas.ImageScaleSmooth
	mainImage.FillMode = canvas.ImageFillOriginal

	mainArea.Add(mainImage)
	mainArea.Refresh()
}

func toolbar(parent fyne.Window) *widget.Toolbar {
	toolbarWidget := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() { loadSelfies(parent) }),
		widget.NewToolbarAction(theme.DownloadIcon(), func() { exportSelfiesStatus(parent) }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MoveUpIcon(), func() { selectPreviousListSelfieSet() }),
		widget.NewToolbarAction(theme.MoveDownIcon(), func() { selectNextListSelfieSet() }),
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

			setSelfieSets(selfie.ConvertToSelfieSets(imageFiles))
		}
	}, parent)

	directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
	if err != nil {
		log.Fatal(err)
	}

	fileDialogURI := storage.NewFileURI(directory)
	fileDialogLister, _ := storage.ListerForURI(fileDialogURI)
	folderOpenDialog.SetLocation(fileDialogLister)

	folderOpenDialog.Show()
}

func setSelfieSets(selfies []*selfie.SelfieSet) {
	selfieSets = selfies
	selectedSelfieSetIndex = -1
	setMainAreaSelfies(nil)

	if selfieSelectionListWidget != nil {
		selfieSelectionListWidget.UnselectAll()
		selfieSelectionListWidget.ScrollToTop()
		selfieSelectionListWidget.Refresh()
	}
}

func exportSelfiesStatus(parent fyne.Window) {
	jsonData, err := json.MarshalIndent(selfieSets, "", "  ")
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

func keyTypeHandler() func(event *fyne.KeyEvent) {
	return func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyUp:
			selectPreviousListSelfieSet()

		case fyne.KeyDown:
			selectNextListSelfieSet()

		case fyne.KeyLeft:
			scrollSelfies(1)

		case fyne.KeyRight:
			scrollSelfies(-1)

		case fyne.Key1:
			setSelfieSetStatus(selfie.SelfieSetStatusNotHandled)

		case fyne.Key2:
			setSelfieSetStatus(selfie.SelfieSetStatusOk)

		case fyne.Key3:
			setSelfieSetStatus(selfie.SelfieSetStatusSuspicious)

		case fyne.Key4:
			setSelfieSetStatus(selfie.SelfieSetStatusFake)
		}
	}
}

func getSelectedSelfieSet() *selfie.SelfieSet {
	if selectedSelfieSetIndex == -1 {
		return nil
	} else {
		return selfieSets[selectedSelfieSetIndex]
	}
}

func scrollSelfies(direction int) {
	if secondaryAreaScroll != nil {
		dx := float32(300) / 1 // Default value, scroll 1/1 (100%) of 300 pixels (300 pixels are supposed selfie image width)

		scrollableSelfiesContainer := secondaryAreaScroll.Content.(*fyne.Container)
		amountScrollableSelfies := len(scrollableSelfiesContainer.Objects)
		if amountScrollableSelfies > 0 {
			containerMinWidth := scrollableSelfiesContainer.MinSize().Width        // Min width of container with al scrollable selfies
			selfieMinWidth := containerMinWidth / float32(amountScrollableSelfies) // Actual selfie image (mean) min width in container
			dx = selfieMinWidth / 1                                                // Scroll 1/1 (100%) image min width
		}

		// Make a smooth scroll
		animationLastValue := float32(0)
		animation := fyne.NewAnimation(canvas.DurationStandard, func(f float32) {
			animationDeltaProgress := f - animationLastValue
			secondaryAreaScroll.Scrolled(&fyne.ScrollEvent{Scrolled: fyne.Delta{DX: float32(direction) * dx * animationDeltaProgress, DY: 0}})
			animationLastValue = f
		})
		animation.Curve = fyne.AnimationEaseOut
		animation.Start()

		// Make quick scroll (rather jump) in scroll pane
		// secondaryAreaScroll.Scrolled(&fyne.ScrollEvent{Scrolled: fyne.Delta{DX: dx, DY: 0}})
	}
}

func selectNextListSelfieSet() {
	selfieSetCount := len(selfieSets)
	if selectedSelfieSetIndex == -1 && selfieSetCount > 0 {
		selectedSelfieSetIndex = 0
		selfieSelectionListWidget.Select(selectedSelfieSetIndex)
	} else if selectedSelfieSetIndex >= 0 && selectedSelfieSetIndex < selfieSetCount-1 {
		selfieSelectionListWidget.Unselect(selectedSelfieSetIndex)
		selectedSelfieSetIndex++
		selectedSelfieSetIndex = min(selfieSetCount-1, selectedSelfieSetIndex)
		selfieSelectionListWidget.Select(selectedSelfieSetIndex)
	}
}

func selectPreviousListSelfieSet() {
	selfieSetCount := len(selfieSets)
	if selectedSelfieSetIndex == -1 && selfieSetCount > 0 {
		selectedSelfieSetIndex = selfieSetCount - 1
		selfieSelectionListWidget.Select(selectedSelfieSetIndex)
	} else if selectedSelfieSetIndex > 0 && selectedSelfieSetIndex < selfieSetCount {
		selfieSelectionListWidget.Unselect(selectedSelfieSetIndex)
		selectedSelfieSetIndex--
		selectedSelfieSetIndex = max(0, selectedSelfieSetIndex)
		selfieSelectionListWidget.Select(selectedSelfieSetIndex)
	}
}
