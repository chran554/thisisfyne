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

var (
	emptyImage = image.Image(image.NewRGBA(image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 1, Y: 1}}))
)

type SelfieSetTreeItem struct {
	widget.BaseWidget

	ImageSize int

	selfieSet *selfie.SelfieSet

	textUI   *canvas.Text
	imageUI  *canvas.Image
	statusUI *StatusIcon
}

func NewSelfieSetTreeItem(size int) *SelfieSetTreeItem {
	selfieSetTreeItem := &SelfieSetTreeItem{
		ImageSize: size,
	}

	selfieSetTreeItem.ExtendBaseWidget(selfieSetTreeItem)

	return selfieSetTreeItem
}

func (li *SelfieSetTreeItem) CreateRenderer() fyne.WidgetRenderer {
	img := emptyImage
	txt := ""

	canvasText := canvas.NewText(txt, colornames.Gray)
	canvasText.Alignment = fyne.TextAlignCenter
	canvasText.TextSize = 10

	canvasImg := canvas.NewImageFromImage(img)
	canvasImg.SetMinSize(fyne.NewSquareSize(float32(li.ImageSize)))
	canvasImg.ScaleMode = canvas.ImageScaleSmooth
	canvasImg.FillMode = canvas.ImageFillContain

	statusIcon := NewStatusIcon()

	li.textUI = canvasText
	li.imageUI = canvasImg
	li.statusUI = statusIcon

	bg := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 24})
	bg.CornerRadius = theme.SelectionRadiusSize()
	bg.StrokeColor = color.RGBA{R: 128, G: 128, B: 128, A: 32}
	bg.StrokeWidth = 1.5

	li.updateUI()

	c := container.NewBorder(nil, canvasText, statusIcon, canvasImg)
	c = container.NewPadded(container.NewPadded(c))
	c = container.NewStack(bg, c)

	return widget.NewSimpleRenderer(c)
}

func (li *SelfieSetTreeItem) SetSelfieSet(selfieSet *selfie.SelfieSet) {
	li.selfieSet = selfieSet
	li.statusUI.SetSelfieSetStatus(selfieSet.Status)
	li.updateUI()
	li.Refresh()
}

func (li *SelfieSetTreeItem) updateUI() {
	img := emptyImage
	txt := ""

	if li.selfieSet != nil {
		img = *li.selfieSet.PrimaryImage.Image
		txt = li.selfieSet.Name
	}

	li.imageUI.Image = img
	li.textUI.Text = txt
}
