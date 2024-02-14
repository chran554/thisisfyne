package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"thisisfyne/internal/app/selfie"
)

type SelfieTreeItem struct {
	widget.BaseWidget

	Selfies *selfie.SelfieSet

	ImageSize int

	// private UI components
	textUI   *canvas.Text
	imageUI  *canvas.Image
	statusUI *StatusIcon
}

func NewSelfieTreeItem(selfies *selfie.SelfieSet, size int) *SelfieTreeItem {
	selfieTreeItem := &SelfieTreeItem{
		Selfies:   selfies,
		ImageSize: size,
	}

	selfieTreeItem.ExtendBaseWidget(selfieTreeItem)

	return selfieTreeItem
}

//func (li *LabelledImage) Tapped(*fyne.PointEvent) {
//	log.Printf("Clicked labelled image '%s'", li.Text)
//}

func (li *SelfieTreeItem) CreateRenderer() fyne.WidgetRenderer {
	var img *image.Image
	var txt string
	var sts selfie.SelfieSetStatus

	if li.Selfies == nil {
		tmpImg := image.Image(image.NewRGBA(image.Rectangle{Min: image.Point{X: 50, Y: 50}, Max: image.Point{X: 50, Y: 50}}))
		img = &tmpImg
		txt = ""
		sts = selfie.SelfieSetStatusNotHandled
	} else {
		img = li.Selfies.PrimaryImage.Image
		txt = li.Selfies.Name
		sts = li.Selfies.Status
	}

	canvasText := canvas.NewText(txt, colornames.Gray)
	canvasText.Alignment = fyne.TextAlignCenter
	canvasText.TextSize = 10

	canvasImg := canvas.NewImageFromImage(*img)
	canvasImg.SetMinSize(fyne.NewSquareSize(float32(li.ImageSize)))
	canvasImg.ScaleMode = canvas.ImageScaleSmooth
	canvasImg.FillMode = canvas.ImageFillContain

	// widgetButton := widget.NewButton("", nil)
	// widgetButton.Icon, widgetButton.Importance = iconAttributesFromStatus(sts, true)
	// widgetButton.Disable()
	// widgetButton.IconPlacement = widget.ButtonIconLeadingText

	statusIcon := NewStatusIcon(sts)

	li.textUI = canvasText
	li.imageUI = canvasImg
	li.statusUI = statusIcon

	bg := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 24})
	bg.CornerRadius = theme.SelectionRadiusSize()
	bg.StrokeColor = color.RGBA{R: 128, G: 128, B: 128, A: 32}
	bg.StrokeWidth = 1.5

	c := container.NewBorder(nil, canvasText, statusIcon, canvasImg)
	c = container.NewPadded(container.NewPadded(c))
	c = container.NewStack(bg, c)

	return widget.NewSimpleRenderer(c)
}

func (li *SelfieTreeItem) SetSelfies(selfies *selfie.SelfieSet) {
	li.Selfies = selfies

	li.imageUI.Image = *li.Selfies.PrimaryImage.Image
	li.textUI.Text = li.Selfies.Name
	li.statusUI.SetStatus(li.Selfies.Status)
	//li.statusUI.Icon, li.statusUI.Importance = iconAttributesFromStatus(li.Selfies.Status, true)

	li.imageUI.SetMinSize(fyne.NewSquareSize(float32(li.ImageSize)))

	li.Refresh()
}
